// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package usecase

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/sknv/passwordless-verifier/internal/model"
)

// Ensure, that StoreMock does implement Store.
// If this is not the case, regenerate this file with moq.
var _ Store = &StoreMock{}

// StoreMock is a mock implementation of Store.
//
// 	func TestSomethingThatUsesStore(t *testing.T) {
//
// 		// make and configure a mocked Store
// 		mockedStore := &StoreMock{
// 			CreateVerificationFunc: func(ctx context.Context, verification *model.Verification) error {
// 				panic("mock out the CreateVerification method")
// 			},
// 			FindLatestVerificationByChatIDFunc: func(ctx context.Context, chatID int64) (*model.Verification, error) {
// 				panic("mock out the FindLatestVerificationByChatID method")
// 			},
// 			FindVerificationByIDFunc: func(ctx context.Context, id uuid.UUID) (*model.Verification, error) {
// 				panic("mock out the FindVerificationByID method")
// 			},
// 			FindVerificationByIDWithSessionFunc: func(ctx context.Context, id uuid.UUID) (*model.Verification, error) {
// 				panic("mock out the FindVerificationByIDWithSession method")
// 			},
// 			UpdateVerificationFunc: func(ctx context.Context, verification *model.Verification) error {
// 				panic("mock out the UpdateVerification method")
// 			},
// 			UpdateVerificationAndCreateSessionFunc: func(ctx context.Context, verification *model.Verification) error {
// 				panic("mock out the UpdateVerificationAndCreateSession method")
// 			},
// 		}
//
// 		// use mockedStore in code that requires Store
// 		// and then make assertions.
//
// 	}
type StoreMock struct {
	// CreateVerificationFunc mocks the CreateVerification method.
	CreateVerificationFunc func(ctx context.Context, verification *model.Verification) error

	// FindLatestVerificationByChatIDFunc mocks the FindLatestVerificationByChatID method.
	FindLatestVerificationByChatIDFunc func(ctx context.Context, chatID int64) (*model.Verification, error)

	// FindVerificationByIDFunc mocks the FindVerificationByID method.
	FindVerificationByIDFunc func(ctx context.Context, id uuid.UUID) (*model.Verification, error)

	// FindVerificationByIDWithSessionFunc mocks the FindVerificationByIDWithSession method.
	FindVerificationByIDWithSessionFunc func(ctx context.Context, id uuid.UUID) (*model.Verification, error)

	// UpdateVerificationFunc mocks the UpdateVerification method.
	UpdateVerificationFunc func(ctx context.Context, verification *model.Verification) error

	// UpdateVerificationAndCreateSessionFunc mocks the UpdateVerificationAndCreateSession method.
	UpdateVerificationAndCreateSessionFunc func(ctx context.Context, verification *model.Verification) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateVerification holds details about calls to the CreateVerification method.
		CreateVerification []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Verification is the verification argument value.
			Verification *model.Verification
		}
		// FindLatestVerificationByChatID holds details about calls to the FindLatestVerificationByChatID method.
		FindLatestVerificationByChatID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ChatID is the chatID argument value.
			ChatID int64
		}
		// FindVerificationByID holds details about calls to the FindVerificationByID method.
		FindVerificationByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// FindVerificationByIDWithSession holds details about calls to the FindVerificationByIDWithSession method.
		FindVerificationByIDWithSession []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// UpdateVerification holds details about calls to the UpdateVerification method.
		UpdateVerification []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Verification is the verification argument value.
			Verification *model.Verification
		}
		// UpdateVerificationAndCreateSession holds details about calls to the UpdateVerificationAndCreateSession method.
		UpdateVerificationAndCreateSession []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Verification is the verification argument value.
			Verification *model.Verification
		}
	}
	lockCreateVerification                 sync.RWMutex
	lockFindLatestVerificationByChatID     sync.RWMutex
	lockFindVerificationByID               sync.RWMutex
	lockFindVerificationByIDWithSession    sync.RWMutex
	lockUpdateVerification                 sync.RWMutex
	lockUpdateVerificationAndCreateSession sync.RWMutex
}

// CreateVerification calls CreateVerificationFunc.
func (mock *StoreMock) CreateVerification(ctx context.Context, verification *model.Verification) error {
	if mock.CreateVerificationFunc == nil {
		panic("StoreMock.CreateVerificationFunc: method is nil but Store.CreateVerification was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		Verification *model.Verification
	}{
		Ctx:          ctx,
		Verification: verification,
	}
	mock.lockCreateVerification.Lock()
	mock.calls.CreateVerification = append(mock.calls.CreateVerification, callInfo)
	mock.lockCreateVerification.Unlock()
	return mock.CreateVerificationFunc(ctx, verification)
}

// CreateVerificationCalls gets all the calls that were made to CreateVerification.
// Check the length with:
//     len(mockedStore.CreateVerificationCalls())
func (mock *StoreMock) CreateVerificationCalls() []struct {
	Ctx          context.Context
	Verification *model.Verification
} {
	var calls []struct {
		Ctx          context.Context
		Verification *model.Verification
	}
	mock.lockCreateVerification.RLock()
	calls = mock.calls.CreateVerification
	mock.lockCreateVerification.RUnlock()
	return calls
}

// FindLatestVerificationByChatID calls FindLatestVerificationByChatIDFunc.
func (mock *StoreMock) FindLatestVerificationByChatID(ctx context.Context, chatID int64) (*model.Verification, error) {
	if mock.FindLatestVerificationByChatIDFunc == nil {
		panic("StoreMock.FindLatestVerificationByChatIDFunc: method is nil but Store.FindLatestVerificationByChatID was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		ChatID int64
	}{
		Ctx:    ctx,
		ChatID: chatID,
	}
	mock.lockFindLatestVerificationByChatID.Lock()
	mock.calls.FindLatestVerificationByChatID = append(mock.calls.FindLatestVerificationByChatID, callInfo)
	mock.lockFindLatestVerificationByChatID.Unlock()
	return mock.FindLatestVerificationByChatIDFunc(ctx, chatID)
}

// FindLatestVerificationByChatIDCalls gets all the calls that were made to FindLatestVerificationByChatID.
// Check the length with:
//     len(mockedStore.FindLatestVerificationByChatIDCalls())
func (mock *StoreMock) FindLatestVerificationByChatIDCalls() []struct {
	Ctx    context.Context
	ChatID int64
} {
	var calls []struct {
		Ctx    context.Context
		ChatID int64
	}
	mock.lockFindLatestVerificationByChatID.RLock()
	calls = mock.calls.FindLatestVerificationByChatID
	mock.lockFindLatestVerificationByChatID.RUnlock()
	return calls
}

// FindVerificationByID calls FindVerificationByIDFunc.
func (mock *StoreMock) FindVerificationByID(ctx context.Context, id uuid.UUID) (*model.Verification, error) {
	if mock.FindVerificationByIDFunc == nil {
		panic("StoreMock.FindVerificationByIDFunc: method is nil but Store.FindVerificationByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockFindVerificationByID.Lock()
	mock.calls.FindVerificationByID = append(mock.calls.FindVerificationByID, callInfo)
	mock.lockFindVerificationByID.Unlock()
	return mock.FindVerificationByIDFunc(ctx, id)
}

// FindVerificationByIDCalls gets all the calls that were made to FindVerificationByID.
// Check the length with:
//     len(mockedStore.FindVerificationByIDCalls())
func (mock *StoreMock) FindVerificationByIDCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockFindVerificationByID.RLock()
	calls = mock.calls.FindVerificationByID
	mock.lockFindVerificationByID.RUnlock()
	return calls
}

// FindVerificationByIDWithSession calls FindVerificationByIDWithSessionFunc.
func (mock *StoreMock) FindVerificationByIDWithSession(ctx context.Context, id uuid.UUID) (*model.Verification, error) {
	if mock.FindVerificationByIDWithSessionFunc == nil {
		panic("StoreMock.FindVerificationByIDWithSessionFunc: method is nil but Store.FindVerificationByIDWithSession was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockFindVerificationByIDWithSession.Lock()
	mock.calls.FindVerificationByIDWithSession = append(mock.calls.FindVerificationByIDWithSession, callInfo)
	mock.lockFindVerificationByIDWithSession.Unlock()
	return mock.FindVerificationByIDWithSessionFunc(ctx, id)
}

// FindVerificationByIDWithSessionCalls gets all the calls that were made to FindVerificationByIDWithSession.
// Check the length with:
//     len(mockedStore.FindVerificationByIDWithSessionCalls())
func (mock *StoreMock) FindVerificationByIDWithSessionCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockFindVerificationByIDWithSession.RLock()
	calls = mock.calls.FindVerificationByIDWithSession
	mock.lockFindVerificationByIDWithSession.RUnlock()
	return calls
}

// UpdateVerification calls UpdateVerificationFunc.
func (mock *StoreMock) UpdateVerification(ctx context.Context, verification *model.Verification) error {
	if mock.UpdateVerificationFunc == nil {
		panic("StoreMock.UpdateVerificationFunc: method is nil but Store.UpdateVerification was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		Verification *model.Verification
	}{
		Ctx:          ctx,
		Verification: verification,
	}
	mock.lockUpdateVerification.Lock()
	mock.calls.UpdateVerification = append(mock.calls.UpdateVerification, callInfo)
	mock.lockUpdateVerification.Unlock()
	return mock.UpdateVerificationFunc(ctx, verification)
}

// UpdateVerificationCalls gets all the calls that were made to UpdateVerification.
// Check the length with:
//     len(mockedStore.UpdateVerificationCalls())
func (mock *StoreMock) UpdateVerificationCalls() []struct {
	Ctx          context.Context
	Verification *model.Verification
} {
	var calls []struct {
		Ctx          context.Context
		Verification *model.Verification
	}
	mock.lockUpdateVerification.RLock()
	calls = mock.calls.UpdateVerification
	mock.lockUpdateVerification.RUnlock()
	return calls
}

// UpdateVerificationAndCreateSession calls UpdateVerificationAndCreateSessionFunc.
func (mock *StoreMock) UpdateVerificationAndCreateSession(ctx context.Context, verification *model.Verification) error {
	if mock.UpdateVerificationAndCreateSessionFunc == nil {
		panic("StoreMock.UpdateVerificationAndCreateSessionFunc: method is nil but Store.UpdateVerificationAndCreateSession was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		Verification *model.Verification
	}{
		Ctx:          ctx,
		Verification: verification,
	}
	mock.lockUpdateVerificationAndCreateSession.Lock()
	mock.calls.UpdateVerificationAndCreateSession = append(mock.calls.UpdateVerificationAndCreateSession, callInfo)
	mock.lockUpdateVerificationAndCreateSession.Unlock()
	return mock.UpdateVerificationAndCreateSessionFunc(ctx, verification)
}

// UpdateVerificationAndCreateSessionCalls gets all the calls that were made to UpdateVerificationAndCreateSession.
// Check the length with:
//     len(mockedStore.UpdateVerificationAndCreateSessionCalls())
func (mock *StoreMock) UpdateVerificationAndCreateSessionCalls() []struct {
	Ctx          context.Context
	Verification *model.Verification
} {
	var calls []struct {
		Ctx          context.Context
		Verification *model.Verification
	}
	mock.lockUpdateVerificationAndCreateSession.RLock()
	calls = mock.calls.UpdateVerificationAndCreateSession
	mock.lockUpdateVerificationAndCreateSession.RUnlock()
	return calls
}
