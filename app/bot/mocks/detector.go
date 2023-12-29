// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"github.com/umputun/tg-spam/lib/approved"
	"github.com/umputun/tg-spam/lib/spamcheck"
	"github.com/umputun/tg-spam/lib/tgspam"
	"io"
	"sync"
)

// DetectorMock is a mock implementation of bot.Detector.
//
//	func TestSomethingThatUsesDetector(t *testing.T) {
//
//		// make and configure a mocked bot.Detector
//		mockedDetector := &DetectorMock{
//			AddApprovedUserFunc: func(user approved.UserInfo) error {
//				panic("mock out the AddApprovedUser method")
//			},
//			ApprovedUsersFunc: func() []approved.UserInfo {
//				panic("mock out the ApprovedUsers method")
//			},
//			CheckFunc: func(request spamcheck.Request) (bool, []spamcheck.Response) {
//				panic("mock out the Check method")
//			},
//			IsApprovedUserFunc: func(userID string) bool {
//				panic("mock out the IsApprovedUser method")
//			},
//			LoadSamplesFunc: func(exclReader io.Reader, spamReaders []io.Reader, hamReaders []io.Reader) (tgspam.LoadResult, error) {
//				panic("mock out the LoadSamples method")
//			},
//			LoadStopWordsFunc: func(readers ...io.Reader) (tgspam.LoadResult, error) {
//				panic("mock out the LoadStopWords method")
//			},
//			RemoveApprovedUserFunc: func(id string) error {
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
//		// use mockedDetector in code that requires bot.Detector
//		// and then make assertions.
//
//	}
type DetectorMock struct {
	// AddApprovedUserFunc mocks the AddApprovedUser method.
	AddApprovedUserFunc func(user approved.UserInfo) error

	// ApprovedUsersFunc mocks the ApprovedUsers method.
	ApprovedUsersFunc func() []approved.UserInfo

	// CheckFunc mocks the Check method.
	CheckFunc func(request spamcheck.Request) (bool, []spamcheck.Response)

	// IsApprovedUserFunc mocks the IsApprovedUser method.
	IsApprovedUserFunc func(userID string) bool

	// LoadSamplesFunc mocks the LoadSamples method.
	LoadSamplesFunc func(exclReader io.Reader, spamReaders []io.Reader, hamReaders []io.Reader) (tgspam.LoadResult, error)

	// LoadStopWordsFunc mocks the LoadStopWords method.
	LoadStopWordsFunc func(readers ...io.Reader) (tgspam.LoadResult, error)

	// RemoveApprovedUserFunc mocks the RemoveApprovedUser method.
	RemoveApprovedUserFunc func(id string) error

	// UpdateHamFunc mocks the UpdateHam method.
	UpdateHamFunc func(msg string) error

	// UpdateSpamFunc mocks the UpdateSpam method.
	UpdateSpamFunc func(msg string) error

	// calls tracks calls to the methods.
	calls struct {
		// AddApprovedUser holds details about calls to the AddApprovedUser method.
		AddApprovedUser []struct {
			// User is the user argument value.
			User approved.UserInfo
		}
		// ApprovedUsers holds details about calls to the ApprovedUsers method.
		ApprovedUsers []struct {
		}
		// Check holds details about calls to the Check method.
		Check []struct {
			// Request is the request argument value.
			Request spamcheck.Request
		}
		// IsApprovedUser holds details about calls to the IsApprovedUser method.
		IsApprovedUser []struct {
			// UserID is the userID argument value.
			UserID string
		}
		// LoadSamples holds details about calls to the LoadSamples method.
		LoadSamples []struct {
			// ExclReader is the exclReader argument value.
			ExclReader io.Reader
			// SpamReaders is the spamReaders argument value.
			SpamReaders []io.Reader
			// HamReaders is the hamReaders argument value.
			HamReaders []io.Reader
		}
		// LoadStopWords holds details about calls to the LoadStopWords method.
		LoadStopWords []struct {
			// Readers is the readers argument value.
			Readers []io.Reader
		}
		// RemoveApprovedUser holds details about calls to the RemoveApprovedUser method.
		RemoveApprovedUser []struct {
			// ID is the id argument value.
			ID string
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
	lockApprovedUsers      sync.RWMutex
	lockCheck              sync.RWMutex
	lockIsApprovedUser     sync.RWMutex
	lockLoadSamples        sync.RWMutex
	lockLoadStopWords      sync.RWMutex
	lockRemoveApprovedUser sync.RWMutex
	lockUpdateHam          sync.RWMutex
	lockUpdateSpam         sync.RWMutex
}

// AddApprovedUser calls AddApprovedUserFunc.
func (mock *DetectorMock) AddApprovedUser(user approved.UserInfo) error {
	if mock.AddApprovedUserFunc == nil {
		panic("DetectorMock.AddApprovedUserFunc: method is nil but Detector.AddApprovedUser was just called")
	}
	callInfo := struct {
		User approved.UserInfo
	}{
		User: user,
	}
	mock.lockAddApprovedUser.Lock()
	mock.calls.AddApprovedUser = append(mock.calls.AddApprovedUser, callInfo)
	mock.lockAddApprovedUser.Unlock()
	return mock.AddApprovedUserFunc(user)
}

// AddApprovedUserCalls gets all the calls that were made to AddApprovedUser.
// Check the length with:
//
//	len(mockedDetector.AddApprovedUserCalls())
func (mock *DetectorMock) AddApprovedUserCalls() []struct {
	User approved.UserInfo
} {
	var calls []struct {
		User approved.UserInfo
	}
	mock.lockAddApprovedUser.RLock()
	calls = mock.calls.AddApprovedUser
	mock.lockAddApprovedUser.RUnlock()
	return calls
}

// ResetAddApprovedUserCalls reset all the calls that were made to AddApprovedUser.
func (mock *DetectorMock) ResetAddApprovedUserCalls() {
	mock.lockAddApprovedUser.Lock()
	mock.calls.AddApprovedUser = nil
	mock.lockAddApprovedUser.Unlock()
}

// ApprovedUsers calls ApprovedUsersFunc.
func (mock *DetectorMock) ApprovedUsers() []approved.UserInfo {
	if mock.ApprovedUsersFunc == nil {
		panic("DetectorMock.ApprovedUsersFunc: method is nil but Detector.ApprovedUsers was just called")
	}
	callInfo := struct {
	}{}
	mock.lockApprovedUsers.Lock()
	mock.calls.ApprovedUsers = append(mock.calls.ApprovedUsers, callInfo)
	mock.lockApprovedUsers.Unlock()
	return mock.ApprovedUsersFunc()
}

// ApprovedUsersCalls gets all the calls that were made to ApprovedUsers.
// Check the length with:
//
//	len(mockedDetector.ApprovedUsersCalls())
func (mock *DetectorMock) ApprovedUsersCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockApprovedUsers.RLock()
	calls = mock.calls.ApprovedUsers
	mock.lockApprovedUsers.RUnlock()
	return calls
}

// ResetApprovedUsersCalls reset all the calls that were made to ApprovedUsers.
func (mock *DetectorMock) ResetApprovedUsersCalls() {
	mock.lockApprovedUsers.Lock()
	mock.calls.ApprovedUsers = nil
	mock.lockApprovedUsers.Unlock()
}

// Check calls CheckFunc.
func (mock *DetectorMock) Check(request spamcheck.Request) (bool, []spamcheck.Response) {
	if mock.CheckFunc == nil {
		panic("DetectorMock.CheckFunc: method is nil but Detector.Check was just called")
	}
	callInfo := struct {
		Request spamcheck.Request
	}{
		Request: request,
	}
	mock.lockCheck.Lock()
	mock.calls.Check = append(mock.calls.Check, callInfo)
	mock.lockCheck.Unlock()
	return mock.CheckFunc(request)
}

// CheckCalls gets all the calls that were made to Check.
// Check the length with:
//
//	len(mockedDetector.CheckCalls())
func (mock *DetectorMock) CheckCalls() []struct {
	Request spamcheck.Request
} {
	var calls []struct {
		Request spamcheck.Request
	}
	mock.lockCheck.RLock()
	calls = mock.calls.Check
	mock.lockCheck.RUnlock()
	return calls
}

// ResetCheckCalls reset all the calls that were made to Check.
func (mock *DetectorMock) ResetCheckCalls() {
	mock.lockCheck.Lock()
	mock.calls.Check = nil
	mock.lockCheck.Unlock()
}

// IsApprovedUser calls IsApprovedUserFunc.
func (mock *DetectorMock) IsApprovedUser(userID string) bool {
	if mock.IsApprovedUserFunc == nil {
		panic("DetectorMock.IsApprovedUserFunc: method is nil but Detector.IsApprovedUser was just called")
	}
	callInfo := struct {
		UserID string
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
//	len(mockedDetector.IsApprovedUserCalls())
func (mock *DetectorMock) IsApprovedUserCalls() []struct {
	UserID string
} {
	var calls []struct {
		UserID string
	}
	mock.lockIsApprovedUser.RLock()
	calls = mock.calls.IsApprovedUser
	mock.lockIsApprovedUser.RUnlock()
	return calls
}

// ResetIsApprovedUserCalls reset all the calls that were made to IsApprovedUser.
func (mock *DetectorMock) ResetIsApprovedUserCalls() {
	mock.lockIsApprovedUser.Lock()
	mock.calls.IsApprovedUser = nil
	mock.lockIsApprovedUser.Unlock()
}

// LoadSamples calls LoadSamplesFunc.
func (mock *DetectorMock) LoadSamples(exclReader io.Reader, spamReaders []io.Reader, hamReaders []io.Reader) (tgspam.LoadResult, error) {
	if mock.LoadSamplesFunc == nil {
		panic("DetectorMock.LoadSamplesFunc: method is nil but Detector.LoadSamples was just called")
	}
	callInfo := struct {
		ExclReader  io.Reader
		SpamReaders []io.Reader
		HamReaders  []io.Reader
	}{
		ExclReader:  exclReader,
		SpamReaders: spamReaders,
		HamReaders:  hamReaders,
	}
	mock.lockLoadSamples.Lock()
	mock.calls.LoadSamples = append(mock.calls.LoadSamples, callInfo)
	mock.lockLoadSamples.Unlock()
	return mock.LoadSamplesFunc(exclReader, spamReaders, hamReaders)
}

// LoadSamplesCalls gets all the calls that were made to LoadSamples.
// Check the length with:
//
//	len(mockedDetector.LoadSamplesCalls())
func (mock *DetectorMock) LoadSamplesCalls() []struct {
	ExclReader  io.Reader
	SpamReaders []io.Reader
	HamReaders  []io.Reader
} {
	var calls []struct {
		ExclReader  io.Reader
		SpamReaders []io.Reader
		HamReaders  []io.Reader
	}
	mock.lockLoadSamples.RLock()
	calls = mock.calls.LoadSamples
	mock.lockLoadSamples.RUnlock()
	return calls
}

// ResetLoadSamplesCalls reset all the calls that were made to LoadSamples.
func (mock *DetectorMock) ResetLoadSamplesCalls() {
	mock.lockLoadSamples.Lock()
	mock.calls.LoadSamples = nil
	mock.lockLoadSamples.Unlock()
}

// LoadStopWords calls LoadStopWordsFunc.
func (mock *DetectorMock) LoadStopWords(readers ...io.Reader) (tgspam.LoadResult, error) {
	if mock.LoadStopWordsFunc == nil {
		panic("DetectorMock.LoadStopWordsFunc: method is nil but Detector.LoadStopWords was just called")
	}
	callInfo := struct {
		Readers []io.Reader
	}{
		Readers: readers,
	}
	mock.lockLoadStopWords.Lock()
	mock.calls.LoadStopWords = append(mock.calls.LoadStopWords, callInfo)
	mock.lockLoadStopWords.Unlock()
	return mock.LoadStopWordsFunc(readers...)
}

// LoadStopWordsCalls gets all the calls that were made to LoadStopWords.
// Check the length with:
//
//	len(mockedDetector.LoadStopWordsCalls())
func (mock *DetectorMock) LoadStopWordsCalls() []struct {
	Readers []io.Reader
} {
	var calls []struct {
		Readers []io.Reader
	}
	mock.lockLoadStopWords.RLock()
	calls = mock.calls.LoadStopWords
	mock.lockLoadStopWords.RUnlock()
	return calls
}

// ResetLoadStopWordsCalls reset all the calls that were made to LoadStopWords.
func (mock *DetectorMock) ResetLoadStopWordsCalls() {
	mock.lockLoadStopWords.Lock()
	mock.calls.LoadStopWords = nil
	mock.lockLoadStopWords.Unlock()
}

// RemoveApprovedUser calls RemoveApprovedUserFunc.
func (mock *DetectorMock) RemoveApprovedUser(id string) error {
	if mock.RemoveApprovedUserFunc == nil {
		panic("DetectorMock.RemoveApprovedUserFunc: method is nil but Detector.RemoveApprovedUser was just called")
	}
	callInfo := struct {
		ID string
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
//	len(mockedDetector.RemoveApprovedUserCalls())
func (mock *DetectorMock) RemoveApprovedUserCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockRemoveApprovedUser.RLock()
	calls = mock.calls.RemoveApprovedUser
	mock.lockRemoveApprovedUser.RUnlock()
	return calls
}

// ResetRemoveApprovedUserCalls reset all the calls that were made to RemoveApprovedUser.
func (mock *DetectorMock) ResetRemoveApprovedUserCalls() {
	mock.lockRemoveApprovedUser.Lock()
	mock.calls.RemoveApprovedUser = nil
	mock.lockRemoveApprovedUser.Unlock()
}

// UpdateHam calls UpdateHamFunc.
func (mock *DetectorMock) UpdateHam(msg string) error {
	if mock.UpdateHamFunc == nil {
		panic("DetectorMock.UpdateHamFunc: method is nil but Detector.UpdateHam was just called")
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
//	len(mockedDetector.UpdateHamCalls())
func (mock *DetectorMock) UpdateHamCalls() []struct {
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
func (mock *DetectorMock) ResetUpdateHamCalls() {
	mock.lockUpdateHam.Lock()
	mock.calls.UpdateHam = nil
	mock.lockUpdateHam.Unlock()
}

// UpdateSpam calls UpdateSpamFunc.
func (mock *DetectorMock) UpdateSpam(msg string) error {
	if mock.UpdateSpamFunc == nil {
		panic("DetectorMock.UpdateSpamFunc: method is nil but Detector.UpdateSpam was just called")
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
//	len(mockedDetector.UpdateSpamCalls())
func (mock *DetectorMock) UpdateSpamCalls() []struct {
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
func (mock *DetectorMock) ResetUpdateSpamCalls() {
	mock.lockUpdateSpam.Lock()
	mock.calls.UpdateSpam = nil
	mock.lockUpdateSpam.Unlock()
}

// ResetCalls reset all the calls that were made to all mocked methods.
func (mock *DetectorMock) ResetCalls() {
	mock.lockAddApprovedUser.Lock()
	mock.calls.AddApprovedUser = nil
	mock.lockAddApprovedUser.Unlock()

	mock.lockApprovedUsers.Lock()
	mock.calls.ApprovedUsers = nil
	mock.lockApprovedUsers.Unlock()

	mock.lockCheck.Lock()
	mock.calls.Check = nil
	mock.lockCheck.Unlock()

	mock.lockIsApprovedUser.Lock()
	mock.calls.IsApprovedUser = nil
	mock.lockIsApprovedUser.Unlock()

	mock.lockLoadSamples.Lock()
	mock.calls.LoadSamples = nil
	mock.lockLoadSamples.Unlock()

	mock.lockLoadStopWords.Lock()
	mock.calls.LoadStopWords = nil
	mock.lockLoadStopWords.Unlock()

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
