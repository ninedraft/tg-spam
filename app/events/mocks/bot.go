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
//			AddApprovedUsersFunc: func(id int64, ids ...int64)  {
//				panic("mock out the AddApprovedUsers method")
//			},
//			OnMessageFunc: func(msg bot.Message) bot.Response {
//				panic("mock out the OnMessage method")
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
	// AddApprovedUsersFunc mocks the AddApprovedUsers method.
	AddApprovedUsersFunc func(id int64, ids ...int64)

	// OnMessageFunc mocks the OnMessage method.
	OnMessageFunc func(msg bot.Message) bot.Response

	// UpdateHamFunc mocks the UpdateHam method.
	UpdateHamFunc func(msg string) error

	// UpdateSpamFunc mocks the UpdateSpam method.
	UpdateSpamFunc func(msg string) error

	// calls tracks calls to the methods.
	calls struct {
		// AddApprovedUsers holds details about calls to the AddApprovedUsers method.
		AddApprovedUsers []struct {
			// ID is the id argument value.
			ID int64
			// Ids is the ids argument value.
			Ids []int64
		}
		// OnMessage holds details about calls to the OnMessage method.
		OnMessage []struct {
			// Msg is the msg argument value.
			Msg bot.Message
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
	lockAddApprovedUsers sync.RWMutex
	lockOnMessage        sync.RWMutex
	lockUpdateHam        sync.RWMutex
	lockUpdateSpam       sync.RWMutex
}

// AddApprovedUsers calls AddApprovedUsersFunc.
func (mock *BotMock) AddApprovedUsers(id int64, ids ...int64) {
	if mock.AddApprovedUsersFunc == nil {
		panic("BotMock.AddApprovedUsersFunc: method is nil but Bot.AddApprovedUsers was just called")
	}
	callInfo := struct {
		ID  int64
		Ids []int64
	}{
		ID:  id,
		Ids: ids,
	}
	mock.lockAddApprovedUsers.Lock()
	mock.calls.AddApprovedUsers = append(mock.calls.AddApprovedUsers, callInfo)
	mock.lockAddApprovedUsers.Unlock()
	mock.AddApprovedUsersFunc(id, ids...)
}

// AddApprovedUsersCalls gets all the calls that were made to AddApprovedUsers.
// check the length with:
//
//	len(mockedBot.AddApprovedUsersCalls())
func (mock *BotMock) AddApprovedUsersCalls() []struct {
	ID  int64
	Ids []int64
} {
	var calls []struct {
		ID  int64
		Ids []int64
	}
	mock.lockAddApprovedUsers.RLock()
	calls = mock.calls.AddApprovedUsers
	mock.lockAddApprovedUsers.RUnlock()
	return calls
}

// ResetAddApprovedUsersCalls reset all the calls that were made to AddApprovedUsers.
func (mock *BotMock) ResetAddApprovedUsersCalls() {
	mock.lockAddApprovedUsers.Lock()
	mock.calls.AddApprovedUsers = nil
	mock.lockAddApprovedUsers.Unlock()
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
// check the length with:
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
// check the length with:
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
// check the length with:
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
	mock.lockAddApprovedUsers.Lock()
	mock.calls.AddApprovedUsers = nil
	mock.lockAddApprovedUsers.Unlock()

	mock.lockOnMessage.Lock()
	mock.calls.OnMessage = nil
	mock.lockOnMessage.Unlock()

	mock.lockUpdateHam.Lock()
	mock.calls.UpdateHam = nil
	mock.lockUpdateHam.Unlock()

	mock.lockUpdateSpam.Lock()
	mock.calls.UpdateSpam = nil
	mock.lockUpdateSpam.Unlock()
}
