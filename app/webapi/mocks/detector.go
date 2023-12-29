// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"github.com/umputun/tg-spam/lib/approved"
	"github.com/umputun/tg-spam/lib/spamcheck"
	"sync"
)

// DetectorMock is a mock implementation of webapi.Detector.
//
//	func TestSomethingThatUsesDetector(t *testing.T) {
//
//		// make and configure a mocked webapi.Detector
//		mockedDetector := &DetectorMock{
//			AddApprovedUserFunc: func(user approved.UserInfo) error {
//				panic("mock out the AddApprovedUser method")
//			},
//			ApprovedUsersFunc: func() []approved.UserInfo {
//				panic("mock out the ApprovedUsers method")
//			},
//			CheckFunc: func(req spamcheck.Request) (bool, []spamcheck.Response) {
//				panic("mock out the Check method")
//			},
//			RemoveApprovedUserFunc: func(id string) error {
//				panic("mock out the RemoveApprovedUser method")
//			},
//		}
//
//		// use mockedDetector in code that requires webapi.Detector
//		// and then make assertions.
//
//	}
type DetectorMock struct {
	// AddApprovedUserFunc mocks the AddApprovedUser method.
	AddApprovedUserFunc func(user approved.UserInfo) error

	// ApprovedUsersFunc mocks the ApprovedUsers method.
	ApprovedUsersFunc func() []approved.UserInfo

	// CheckFunc mocks the Check method.
	CheckFunc func(req spamcheck.Request) (bool, []spamcheck.Response)

	// RemoveApprovedUserFunc mocks the RemoveApprovedUser method.
	RemoveApprovedUserFunc func(id string) error

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
			// Req is the req argument value.
			Req spamcheck.Request
		}
		// RemoveApprovedUser holds details about calls to the RemoveApprovedUser method.
		RemoveApprovedUser []struct {
			// ID is the id argument value.
			ID string
		}
	}
	lockAddApprovedUser    sync.RWMutex
	lockApprovedUsers      sync.RWMutex
	lockCheck              sync.RWMutex
	lockRemoveApprovedUser sync.RWMutex
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
func (mock *DetectorMock) Check(req spamcheck.Request) (bool, []spamcheck.Response) {
	if mock.CheckFunc == nil {
		panic("DetectorMock.CheckFunc: method is nil but Detector.Check was just called")
	}
	callInfo := struct {
		Req spamcheck.Request
	}{
		Req: req,
	}
	mock.lockCheck.Lock()
	mock.calls.Check = append(mock.calls.Check, callInfo)
	mock.lockCheck.Unlock()
	return mock.CheckFunc(req)
}

// CheckCalls gets all the calls that were made to Check.
// Check the length with:
//
//	len(mockedDetector.CheckCalls())
func (mock *DetectorMock) CheckCalls() []struct {
	Req spamcheck.Request
} {
	var calls []struct {
		Req spamcheck.Request
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

	mock.lockRemoveApprovedUser.Lock()
	mock.calls.RemoveApprovedUser = nil
	mock.lockRemoveApprovedUser.Unlock()
}
