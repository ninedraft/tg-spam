package events

import (
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-pkgz/notify"
	tbapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hashicorp/go-multierror"

	"github.com/umputun/tg-spam/app/bot"
)

//go:generate moq --out mocks/tb_api.go --pkg mocks --with-resets --skip-ensure . TbAPI
//go:generate moq --out mocks/spam_logger.go --pkg mocks --with-resets --skip-ensure . SpamLogger
//go:generate moq --out mocks/bot.go --pkg mocks --with-resets --skip-ensure . Bot

// TelegramListener listens to tg update, forward to bots and send back responses
// Not thread safe
type TelegramListener struct {
	TbAPI        TbAPI
	SpamLogger   SpamLogger
	Bot          Bot
	Group        string // can be int64 or public group username (without "@" prefix)
	AdminGroup   string // can be int64 or public group username (without "@" prefix)
	IdleDuration time.Duration
	SuperUsers   SuperUser
	StartupMsg   string
	NoSpamReply  bool
	Dry          bool

	AdminURL        string
	AdminSecret     string
	AdminListenAddr string

	chatID      int64
	adminChatID int64

	msgs struct {
		once sync.Once
		ch   chan bot.Response
	}
}

// TbAPI is an interface for telegram bot API, only subset of methods used
type TbAPI interface {
	GetUpdatesChan(config tbapi.UpdateConfig) tbapi.UpdatesChannel
	Send(c tbapi.Chattable) (tbapi.Message, error)
	Request(c tbapi.Chattable) (*tbapi.APIResponse, error)
	GetChat(config tbapi.ChatInfoConfig) (tbapi.Chat, error)
}

// SpamLogger is an interface for spam logger
type SpamLogger interface {
	Save(msg *bot.Message, response *bot.Response)
}

// SpamLoggerFunc is a function that implements SpamLogger interface
type SpamLoggerFunc func(msg *bot.Message, response *bot.Response)

// Save is a function that implements SpamLogger interface
func (f SpamLoggerFunc) Save(msg *bot.Message, response *bot.Response) {
	f(msg, response)
}

// Bot is an interface for bot events.
type Bot interface {
	OnMessage(msg bot.Message) (response bot.Response)
}

// Do process all events, blocked call
func (l *TelegramListener) Do(ctx context.Context) error {
	log.Printf("[INFO] start telegram listener for %q", l.Group)

	var getChatErr error
	if l.chatID, getChatErr = l.getChatID(l.Group); getChatErr != nil {
		return fmt.Errorf("failed to get chat ID for group %q: %w", l.Group, getChatErr)
	}

	if l.AdminGroup != "" {
		if l.adminChatID, getChatErr = l.getChatID(l.AdminGroup); getChatErr != nil {
			return fmt.Errorf("failed to get chat ID for admin group %q: %w", l.AdminGroup, getChatErr)
		}
		log.Printf("[INFO] admin chat ID: %d", l.adminChatID)
	}

	l.msgs.once.Do(func() {
		l.msgs.ch = make(chan bot.Response, 100)
		if l.IdleDuration == 0 {
			l.IdleDuration = 30 * time.Second
		}
	})

	if l.StartupMsg != "" {
		if err := l.sendBotResponse(bot.Response{Send: true, Text: l.StartupMsg}, l.chatID); err != nil {
			log.Printf("[WARN] failed to send startup message, %v", err)
		}
	}

	// run unban server if all required params are set
	if l.AdminURL != "" && l.AdminSecret != "" && l.AdminListenAddr != "" {
		go l.runUnbanServer(ctx)
	}

	u := tbapi.NewUpdate(0)
	u.Timeout = 60

	updates := l.TbAPI.GetUpdatesChan(u)

	for {
		select {

		case <-ctx.Done():
			return ctx.Err()

		case update, ok := <-updates:
			if !ok {
				return fmt.Errorf("telegram update chan closed")
			}

			if update.Message == nil {
				log.Print("[DEBUG] empty message body")
				continue
			}
			if update.Message.Chat == nil {
				log.Print("[DEBUG] ignoring message not from chat")
				continue
			}

			if err := l.procEvents(update); err != nil {
				log.Printf("[WARN] failed to process update: %v", err)
				continue
			}

		case <-time.After(l.IdleDuration): // hit bots on idle timeout
			resp := l.Bot.OnMessage(bot.Message{Text: "idle"})
			if err := l.sendBotResponse(resp, l.chatID); err != nil {
				log.Printf("[WARN] failed to respond on idle, %v", err)
			}
		}
	}
}

func (l *TelegramListener) procEvents(update tbapi.Update) error {
	msgJSON, errJSON := json.Marshal(update.Message)
	if errJSON != nil {
		return fmt.Errorf("failed to marshal update.Message to json: %w", errJSON)
	}
	log.Printf("[DEBUG] %s", string(msgJSON))

	fromChat := update.Message.Chat.ID
	msg := l.transform(update.Message)
	log.Printf("[DEBUG] incoming msg: %+v", msg)

	resp := l.Bot.OnMessage(*msg)

	if resp.Send && !l.NoSpamReply {
		if err := l.sendBotResponse(resp, fromChat); err != nil {
			log.Printf("[WARN] failed to respond on update, %v", err)
		}
	}

	errs := new(multierror.Error)
	isBanInvoked := resp.Send && resp.BanInterval > 0
	// some bots may request a direct ban for given duration
	if isBanInvoked {
		log.Printf("[DEBUG] ban initiated for %+v", resp)
		l.SpamLogger.Save(msg, &resp)
		banUserStr := l.getBanUsername(resp, update)
		if l.SuperUsers.IsSuper(msg.From.Username) {
			log.Printf("[DEBUG] superuser %s requested ban, ignored", banUserStr)
			l.forwardToAdmin(banUserStr, msg) // forward to admin here is for testing only
			return nil
		}
		banSuccessMessage := fmt.Sprintf("[INFO] %s banned by bot for %v", banUserStr, resp.BanInterval)
		if err := l.banUserOrChannel(resp.BanInterval, fromChat, resp.User.ID, resp.ChannelID); err != nil {
			errs = multierror.Append(errs, fmt.Errorf("failed to ban %s: %w", banUserStr, err))
		} else {
			log.Print(banSuccessMessage)
			if l.adminChatID != 0 && msg.From.ID != 0 {
				l.forwardToAdmin(banUserStr, msg)
			}
		}
	}

	// delete message if requested by bot
	if resp.DeleteReplyTo && resp.ReplyTo != 0 && !l.Dry {
		_, err := l.TbAPI.Request(tbapi.DeleteMessageConfig{ChatID: l.chatID, MessageID: resp.ReplyTo})
		if err != nil {
			errs = multierror.Append(errs, fmt.Errorf("failed to delete message %d: %w", resp.ReplyTo, err))
		}
	}
	return errs.ErrorOrNil()
}

func (l *TelegramListener) forwardToAdmin(banUserStr string, msg *bot.Message) {
	forwardMsg := fmt.Sprintf("**permanently banned [%s](tg://user?id=%d)**\n[unban](%s) if it was a mistake\n\n%s\n----",
		banUserStr, msg.From.ID, l.unbanURL(msg.From.ID), strings.ReplaceAll(msg.Text, "\n", " "))
	e := l.sendBotResponse(bot.Response{Send: true, Text: forwardMsg, ParseMode: tbapi.ModeMarkdown}, l.adminChatID)
	if e != nil {
		log.Printf("[WARN] failed to send admin message, %v", e)
	}
}

func (l *TelegramListener) getBanUsername(resp bot.Response, update tbapi.Update) string {
	if resp.ChannelID == 0 {
		return fmt.Sprintf("%v", resp.User)
	}
	botChat := bot.SenderChat{
		ID: resp.ChannelID,
	}
	if update.Message.SenderChat != nil {
		botChat.UserName = update.Message.SenderChat.UserName
	}
	// if not set, that means the ban comes from superuser and username should be taken from ReplyToMessage
	if botChat.UserName == "" && update.Message.ReplyToMessage.SenderChat != nil {
		botChat.UserName = update.Message.ReplyToMessage.SenderChat.UserName
	}
	return fmt.Sprintf("%v", botChat)
}

// sendBotResponse sends bot's answer to tg channel
func (l *TelegramListener) sendBotResponse(resp bot.Response, chatID int64) error {
	if !resp.Send {
		return nil
	}

	log.Printf("[DEBUG] bot response - %+v, reply-to:%d, parse-mode:%s", resp.Text, resp.ReplyTo, resp.ParseMode)
	tbMsg := tbapi.NewMessage(chatID, resp.Text)
	tbMsg.ParseMode = tbapi.ModeMarkdown
	if resp.ParseMode != "" {
		tbMsg.ParseMode = resp.ParseMode
	}
	tbMsg.DisableWebPagePreview = true
	tbMsg.ReplyToMessageID = resp.ReplyTo
	if _, err := l.TbAPI.Send(tbMsg); err != nil {
		return fmt.Errorf("can't send message to telegram %q: %w", resp.Text, err)
	}

	return nil
}

// Submit message text to telegram's group
func (l *TelegramListener) Submit(ctx context.Context, text string) error {
	l.msgs.once.Do(func() { l.msgs.ch = make(chan bot.Response, 100) })

	select {
	case <-ctx.Done():
		return ctx.Err()
	case l.msgs.ch <- bot.Response{Text: text, Send: true}:
	}
	return nil
}

// SubmitHTML message to telegram's group with HTML mode
func (l *TelegramListener) SubmitHTML(ctx context.Context, text string) error {
	// Remove unsupported HTML tags
	text = notify.TelegramSupportedHTML(text)
	l.msgs.once.Do(func() { l.msgs.ch = make(chan bot.Response, 100) })

	select {
	case <-ctx.Done():
		return ctx.Err()
	case l.msgs.ch <- bot.Response{Text: text, Send: true, ParseMode: tbapi.ModeHTML}:
	}
	return nil
}

func (l *TelegramListener) getChatID(group string) (int64, error) {
	chatID, err := strconv.ParseInt(group, 10, 64)
	if err == nil {
		return chatID, nil
	}

	chat, err := l.TbAPI.GetChat(tbapi.ChatInfoConfig{ChatConfig: tbapi.ChatConfig{SuperGroupUsername: "@" + group}})
	if err != nil {
		return 0, fmt.Errorf("can't get chat for %s: %w", group, err)
	}

	return chat.ID, nil
}

// The bot must be an administrator in the supergroup for this to work
// and must have the appropriate admin rights.
// If channel is provided, it is banned instead of provided user, permanently.
func (l *TelegramListener) banUserOrChannel(duration time.Duration, chatID, userID, channelID int64) error {
	// From Telegram Bot API documentation:
	// > If user is restricted for more than 366 days or less than 30 seconds from the current time,
	// > they are considered to be restricted forever
	// Because the API query uses unix timestamp rather than "ban duration",
	// you do not want to accidentally get into this 30-second window of a lifetime ban.
	// In practice BanDuration is equal to ten minutes,
	// so this `if` statement is unlikely to be evaluated to true.

	if l.Dry {
		log.Printf("[INFO] dry run: ban %d for %v", userID, duration)
		return nil
	}

	if duration < 30*time.Second {
		duration = 1 * time.Minute
	}

	if channelID != 0 {
		resp, err := l.TbAPI.Request(tbapi.BanChatSenderChatConfig{
			ChatID:       chatID,
			SenderChatID: channelID,
			UntilDate:    int(time.Now().Add(duration).Unix()),
		})
		if err != nil {
			return err
		}
		if !resp.Ok {
			return fmt.Errorf("response is not Ok: %v", string(resp.Result))
		}
		return nil
	}

	resp, err := l.TbAPI.Request(tbapi.RestrictChatMemberConfig{
		ChatMemberConfig: tbapi.ChatMemberConfig{
			ChatID: chatID,
			UserID: userID,
		},
		UntilDate: time.Now().Add(duration).Unix(),
		Permissions: &tbapi.ChatPermissions{
			CanSendMessages:       false,
			CanSendMediaMessages:  false,
			CanSendOtherMessages:  false,
			CanAddWebPagePreviews: false,
		},
	})
	if err != nil {
		return err
	}
	if !resp.Ok {
		return fmt.Errorf("response is not Ok: %v", string(resp.Result))
	}

	return nil
}

func (l *TelegramListener) transform(msg *tbapi.Message) *bot.Message {
	message := bot.Message{
		ID:   msg.MessageID,
		Sent: msg.Time(),
		Text: msg.Text,
	}

	if msg.Chat != nil {
		message.ChatID = msg.Chat.ID
	}

	if msg.From != nil {
		message.From = bot.User{
			ID:       msg.From.ID,
			Username: msg.From.UserName,
		}
	}

	if msg.From != nil && strings.TrimSpace(msg.From.FirstName) != "" {
		message.From.DisplayName = msg.From.FirstName
	}
	if msg.From != nil && strings.TrimSpace(msg.From.LastName) != "" {
		message.From.DisplayName += " " + msg.From.LastName
	}

	if msg.SenderChat != nil {
		message.SenderChat = bot.SenderChat{
			ID:       msg.SenderChat.ID,
			UserName: msg.SenderChat.UserName,
		}
	}

	switch {
	case msg.Entities != nil && len(msg.Entities) > 0:
		message.Entities = l.transformEntities(msg.Entities)

	case msg.Photo != nil && len(msg.Photo) > 0:
		sizes := msg.Photo
		lastSize := sizes[len(sizes)-1]
		message.Image = &bot.Image{
			FileID:   lastSize.FileID,
			Width:    lastSize.Width,
			Height:   lastSize.Height,
			Caption:  msg.Caption,
			Entities: l.transformEntities(msg.CaptionEntities),
		}
	}

	// fill in the message's reply-to message
	if msg.ReplyToMessage != nil {
		message.ReplyTo.Text = msg.ReplyToMessage.Text
		message.ReplyTo.Sent = msg.ReplyToMessage.Time()
		if msg.ReplyToMessage.From != nil {
			message.ReplyTo.From = bot.User{
				ID:          msg.ReplyToMessage.From.ID,
				Username:    msg.ReplyToMessage.From.UserName,
				DisplayName: msg.ReplyToMessage.From.FirstName + " " + msg.ReplyToMessage.From.LastName,
			}
		}
		if msg.ReplyToMessage.SenderChat != nil {
			message.ReplyTo.SenderChat = bot.SenderChat{
				ID:       msg.ReplyToMessage.SenderChat.ID,
				UserName: msg.ReplyToMessage.SenderChat.UserName,
			}
		}
	}

	return &message
}

func (l *TelegramListener) transformEntities(entities []tbapi.MessageEntity) *[]bot.Entity {
	if len(entities) == 0 {
		return nil
	}

	result := make([]bot.Entity, 0, len(entities))
	for _, entity := range entities {
		e := bot.Entity{
			Type:   entity.Type,
			Offset: entity.Offset,
			Length: entity.Length,
			URL:    entity.URL,
		}
		if entity.User != nil {
			e.User = &bot.User{
				ID:          entity.User.ID,
				Username:    entity.User.UserName,
				DisplayName: entity.User.FirstName + " " + entity.User.LastName,
			}
		}
		result = append(result, e)
	}

	return &result
}

func (l *TelegramListener) unbanURL(userID int64) string {
	// key is SHA1 of user ID + secret
	key := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%d::%s", userID, l.AdminSecret))))
	return fmt.Sprintf("%s?user=%d&token=%s", l.AdminURL, userID, key)
}

func (l *TelegramListener) runUnbanServer(ctx context.Context) {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Application", "tg-spam")
		if _, err := w.Write([]byte("pong")); err != nil {
			log.Printf("[WARN] failed to write response, %v", err)
		}
	})
	mux.HandleFunc("/unban", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("user")
		token := r.URL.Query().Get("token")
		userID, err := l.getChatID(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "failed to get user ID for %q: %v", id, err)
			return
		}
		expToken := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%d::%s", userID, l.AdminSecret))))
		if len(token) != len(expToken) || subtle.ConstantTimeCompare([]byte(token), []byte(expToken)) != 1 {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		log.Printf("[INFO] unban user %s (%d)", id, userID)
		_, err = l.TbAPI.Request(tbapi.UnbanChatMemberConfig{ChatMemberConfig: tbapi.ChatMemberConfig{UserID: userID, ChatID: l.chatID}})
		if err != nil {
			log.Printf("[WARN] failed to unban %s, %v", id, err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "failed to unban %s: %v", id, err)
			return
		}
		if _, err := w.Write([]byte("ok")); err != nil {
			log.Printf("[WARN] failed to write response, %v", err)
		}
	})

	srv := &http.Server{Addr: l.AdminListenAddr, Handler: mux, ReadTimeout: 5 * time.Second, WriteTimeout: 5 * time.Second}

	go func() {
		<-ctx.Done()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("[WARN] failed to shutdown unban server: %v", err)
		}
	}()

	log.Printf("[INFO] start unban server on %s", l.AdminListenAddr)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("[WARN] failed to run unban server: %v", err)
	}
}
