// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"github.com/umputun/tg-spam/app/bot"
	"sync"
)

// BotMock is a mock implementation of events.Bot.
//
//	func TestSomethingThatUsesBot(t *testing.T) {
//
//		// make and configure a mocked events.Bot
//		mockedBot := &BotMock{
//			AddApprovedUserFunc: func(id int64, name string) error {
//				panic("mock out the AddApprovedUser method")
//			},
//			IsApprovedUserFunc: func(userID int64) bool {
//				panic("mock out the IsApprovedUser method")
//			},
//			OnMessageFunc: func(msg bot.Message) bot.Response {
//				panic("mock out the OnMessage method")
//			},
//			RemoveApprovedUserFunc: func(id int64) error {
//				panic("mock out the RemoveApprovedUser method")
//			},
//			UpdateHamFunc: func(msg string) error {
//				panic("mock out the UpdateHam method")
//			},
//			UpdateSpamFunc: func(msg string) error {
//				panic("mock out the UpdateSpam method")
//			},
//		}
//
//		// use mockedBot in code that requires events.Bot
//		// and then make assertions.
//
//	}
type BotMock struct {
	// AddApprovedUserFunc mocks the AddApprovedUser method.
	AddApprovedUserFunc func(id int64, name string) error

	// IsApprovedUserFunc mocks the IsApprovedUser method.
	IsApprovedUserFunc func(userID int64) bool

	// OnMessageFunc mocks the OnMessage method.
	OnMessageFunc func(msg bot.Message) bot.Response

	// RemoveApprovedUserFunc mocks the RemoveApprovedUser method.
	RemoveApprovedUserFunc func(id int64) error

	// UpdateHamFunc mocks the UpdateHam method.
	UpdateHamFunc func(msg string) error

	// UpdateSpamFunc mocks the UpdateSpam method.
	UpdateSpamFunc func(msg string) error

	// calls tracks calls to the methods.
	calls struct {
		// AddApprovedUser holds details about calls to the AddApprovedUser method.
		AddApprovedUser []struct {
			// ID is the id argument value.
			ID int64
			// Name is the name argument value.
			Name string
		}
		// IsApprovedUser holds details about calls to the IsApprovedUser method.
		IsApprovedUser []struct {
			// UserID is the userID argument value.
			UserID int64
		}
		// OnMessage holds details about calls to the OnMessage method.
		OnMessage []struct {
			// Msg is the msg argument value.
			Msg bot.Message
		}
		// RemoveApprovedUser holds details about calls to the RemoveApprovedUser method.
		RemoveApprovedUser []struct {
			// ID is the id argument value.
			ID int64
		}
		// UpdateHam holds details about calls to the UpdateHam method.
		UpdateHam []struct {
			// Msg is the msg argument value.
			Msg string
		}
		// UpdateSpam holds details about calls to the UpdateSpam method.
		UpdateSpam []struct {
			// Msg is the msg argument value.
			Msg string
		}
	}
	lockAddApprovedUser    sync.RWMutex
	lockIsApprovedUser     sync.RWMutex
	lockOnMessage          sync.RWMutex
	lockRemoveApprovedUser sync.RWMutex
	lockUpdateHam          sync.RWMutex
	lockUpdateSpam         sync.RWMutex
}

// AddApprovedUser calls AddApprovedUserFunc.
func (mock *BotMock) AddApprovedUser(id int64, name string) error {
	if mock.AddApprovedUserFunc == nil {
		panic("BotMock.AddApprovedUserFunc: method is nil but Bot.AddApprovedUser was just called")
	}
	callInfo := struct {
		ID   int64
		Name string
	}{
		ID:   id,
		Name: name,
	}
	mock.lockAddApprovedUser.Lock()
	mock.calls.AddApprovedUser = append(mock.calls.AddApprovedUser, callInfo)
	mock.lockAddApprovedUser.Unlock()
	return mock.AddApprovedUserFunc(id, name)
}

// AddApprovedUserCalls gets all the calls that were made to AddApprovedUser.
// Check the length with:
//
//	len(mockedBot.AddApprovedUserCalls())
func (mock *BotMock) AddApprovedUserCalls() []struct {
	ID   int64
	Name string
} {
	var calls []struct {
		ID   int64
		Name string
	}
	mock.lockAddApprovedUser.RLock()
	calls = mock.calls.AddApprovedUser
	mock.lockAddApprovedUser.RUnlock()
	return calls
}

// ResetAddApprovedUserCalls reset all the calls that were made to AddApprovedUser.
func (mock *BotMock) ResetAddApprovedUserCalls() {
	mock.lockAddApprovedUser.Lock()
	mock.calls.AddApprovedUser = nil
	mock.lockAddApprovedUser.Unlock()
}

// IsApprovedUser calls IsApprovedUserFunc.
func (mock *BotMock) IsApprovedUser(userID int64) bool {
	if mock.IsApprovedUserFunc == nil {
		panic("BotMock.IsApprovedUserFunc: method is nil but Bot.IsApprovedUser was just called")
	}
	callInfo := struct {
		UserID int64
	}{
		UserID: userID,
	}
	mock.lockIsApprovedUser.Lock()
	mock.calls.IsApprovedUser = append(mock.calls.IsApprovedUser, callInfo)
	mock.lockIsApprovedUser.Unlock()
	return mock.IsApprovedUserFunc(userID)
}

// IsApprovedUserCalls gets all the calls that were made to IsApprovedUser.
// Check the length with:
//
//	len(mockedBot.IsApprovedUserCalls())
func (mock *BotMock) IsApprovedUserCalls() []struct {
	UserID int64
} {
	var calls []struct {
		UserID int64
	}
	mock.lockIsApprovedUser.RLock()
	calls = mock.calls.IsApprovedUser
	mock.lockIsApprovedUser.RUnlock()
	return calls
}

// ResetIsApprovedUserCalls reset all the calls that were made to IsApprovedUser.
func (mock *BotMock) ResetIsApprovedUserCalls() {
	mock.lockIsApprovedUser.Lock()
	mock.calls.IsApprovedUser = nil
	mock.lockIsApprovedUser.Unlock()
}

// OnMessage calls OnMessageFunc.
func (mock *BotMock) OnMessage(msg bot.Message) bot.Response {
	if mock.OnMessageFunc == nil {
		panic("BotMock.OnMessageFunc: method is nil but Bot.OnMessage was just called")
	}
	callInfo := struct {
		Msg bot.Message
	}{
		Msg: msg,
	}
	mock.lockOnMessage.Lock()
	mock.calls.OnMessage = append(mock.calls.OnMessage, callInfo)
	mock.lockOnMessage.Unlock()
	return mock.OnMessageFunc(msg)
}

// OnMessageCalls gets all the calls that were made to OnMessage.
// Check the length with:
//
//	len(mockedBot.OnMessageCalls())
func (mock *BotMock) OnMessageCalls() []struct {
	Msg bot.Message
} {
	var calls []struct {
		Msg bot.Message
	}
	mock.lockOnMessage.RLock()
	calls = mock.calls.OnMessage
	mock.lockOnMessage.RUnlock()
	return calls
}

// ResetOnMessageCalls reset all the calls that were made to OnMessage.
func (mock *BotMock) ResetOnMessageCalls() {
	mock.lockOnMessage.Lock()
	mock.calls.OnMessage = nil
	mock.lockOnMessage.Unlock()
}

// RemoveApprovedUser calls RemoveApprovedUserFunc.
func (mock *BotMock) RemoveApprovedUser(id int64) error {
	if mock.RemoveApprovedUserFunc == nil {
		panic("BotMock.RemoveApprovedUserFunc: method is nil but Bot.RemoveApprovedUser was just called")
	}
	callInfo := struct {
		ID int64
	}{
		ID: id,
	}
	mock.lockRemoveApprovedUser.Lock()
	mock.calls.RemoveApprovedUser = append(mock.calls.RemoveApprovedUser, callInfo)
	mock.lockRemoveApprovedUser.Unlock()
	return mock.RemoveApprovedUserFunc(id)
}

// RemoveApprovedUserCalls gets all the calls that were made to RemoveApprovedUser.
// Check the length with:
//
//	len(mockedBot.RemoveApprovedUserCalls())
func (mock *BotMock) RemoveApprovedUserCalls() []struct {
	ID int64
} {
	var calls []struct {
		ID int64
	}
	mock.lockRemoveApprovedUser.RLock()
	calls = mock.calls.RemoveApprovedUser
	mock.lockRemoveApprovedUser.RUnlock()
	return calls
}

// ResetRemoveApprovedUserCalls reset all the calls that were made to RemoveApprovedUser.
func (mock *BotMock) ResetRemoveApprovedUserCalls() {
	mock.lockRemoveApprovedUser.Lock()
	mock.calls.RemoveApprovedUser = nil
	mock.lockRemoveApprovedUser.Unlock()
}

// UpdateHam calls UpdateHamFunc.
func (mock *BotMock) UpdateHam(msg string) error {
	if mock.UpdateHamFunc == nil {
		panic("BotMock.UpdateHamFunc: method is nil but Bot.UpdateHam was just called")
	}
	callInfo := struct {
		Msg string
	}{
		Msg: msg,
	}
	mock.lockUpdateHam.Lock()
	mock.calls.UpdateHam = append(mock.calls.UpdateHam, callInfo)
	mock.lockUpdateHam.Unlock()
	return mock.UpdateHamFunc(msg)
}

// UpdateHamCalls gets all the calls that were made to UpdateHam.
// Check the length with:
//
//	len(mockedBot.UpdateHamCalls())
func (mock *BotMock) UpdateHamCalls() []struct {
	Msg string
} {
	var calls []struct {
		Msg string
	}
	mock.lockUpdateHam.RLock()
	calls = mock.calls.UpdateHam
	mock.lockUpdateHam.RUnlock()
	return calls
}

// ResetUpdateHamCalls reset all the calls that were made to UpdateHam.
func (mock *BotMock) ResetUpdateHamCalls() {
	mock.lockUpdateHam.Lock()
	mock.calls.UpdateHam = nil
	mock.lockUpdateHam.Unlock()
}

// UpdateSpam calls UpdateSpamFunc.
func (mock *BotMock) UpdateSpam(msg string) error {
	if mock.UpdateSpamFunc == nil {
		panic("BotMock.UpdateSpamFunc: method is nil but Bot.UpdateSpam was just called")
	}
	callInfo := struct {
		Msg string
	}{
		Msg: msg,
	}
	mock.lockUpdateSpam.Lock()
	mock.calls.UpdateSpam = append(mock.calls.UpdateSpam, callInfo)
	mock.lockUpdateSpam.Unlock()
	return mock.UpdateSpamFunc(msg)
}

// UpdateSpamCalls gets all the calls that were made to UpdateSpam.
// Check the length with:
//
//	len(mockedBot.UpdateSpamCalls())
func (mock *BotMock) UpdateSpamCalls() []struct {
	Msg string
} {
	var calls []struct {
		Msg string
	}
	mock.lockUpdateSpam.RLock()
	calls = mock.calls.UpdateSpam
	mock.lockUpdateSpam.RUnlock()
	return calls
}

// ResetUpdateSpamCalls reset all the calls that were made to UpdateSpam.
func (mock *BotMock) ResetUpdateSpamCalls() {
	mock.lockUpdateSpam.Lock()
	mock.calls.UpdateSpam = nil
	mock.lockUpdateSpam.Unlock()
}

// ResetCalls reset all the calls that were made to all mocked methods.
func (mock *BotMock) ResetCalls() {
	mock.lockAddApprovedUser.Lock()
	mock.calls.AddApprovedUser = nil
	mock.lockAddApprovedUser.Unlock()

	mock.lockIsApprovedUser.Lock()
	mock.calls.IsApprovedUser = nil
	mock.lockIsApprovedUser.Unlock()

	mock.lockOnMessage.Lock()
	mock.calls.OnMessage = nil
	mock.lockOnMessage.Unlock()

	mock.lockRemoveApprovedUser.Lock()
	mock.calls.RemoveApprovedUser = nil
	mock.lockRemoveApprovedUser.Unlock()

	mock.lockUpdateHam.Lock()
	mock.calls.UpdateHam = nil
	mock.lockUpdateHam.Unlock()

	mock.lockUpdateSpam.Lock()
	mock.calls.UpdateSpam = nil
	mock.lockUpdateSpam.Unlock()
}
