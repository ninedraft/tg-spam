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
//			OnMessageFunc: func(msg bot.Message) bot.Response {
//				panic("mock out the OnMessage method")
//			},
//		}
//
//		// use mockedBot in code that requires events.Bot
//		// and then make assertions.
//
//	}
type BotMock struct {
	// OnMessageFunc mocks the OnMessage method.
	OnMessageFunc func(msg bot.Message) bot.Response

	// calls tracks calls to the methods.
	calls struct {
		// OnMessage holds details about calls to the OnMessage method.
		OnMessage []struct {
			// Msg is the msg argument value.
			Msg bot.Message
		}
	}
	lockOnMessage sync.RWMutex
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
