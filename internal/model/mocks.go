// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package model

import (
	"context"
	"database/sql"
	"sync"
)

// Ensure, that DBMock does implement DB.
// If this is not the case, regenerate this file with moq.
var _ DB = &DBMock{}

// DBMock is a mock implementation of DB.
//
// 	func TestSomethingThatUsesDB(t *testing.T) {
//
// 		// make and configure a mocked DB
// 		mockedDB := &DBMock{
// 			CreateFunc: func(ctx context.Context, model any) (sql.Result, error) {
// 				panic("mock out the Create method")
// 			},
// 			FindFunc: func(ctx context.Context, dest any, where string, args ...any) error {
// 				panic("mock out the Find method")
// 			},
// 		}
//
// 		// use mockedDB in code that requires DB
// 		// and then make assertions.
//
// 	}
type DBMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, model any) (sql.Result, error)

	// FindFunc mocks the Find method.
	FindFunc func(ctx context.Context, dest any, where string, args ...any) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Model is the model argument value.
			Model any
		}
		// Find holds details about calls to the Find method.
		Find []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Dest is the dest argument value.
			Dest any
			// Where is the where argument value.
			Where string
			// Args is the args argument value.
			Args []any
		}
	}
	lockCreate sync.RWMutex
	lockFind   sync.RWMutex
}

// Create calls CreateFunc.
func (mock *DBMock) Create(ctx context.Context, model any) (sql.Result, error) {
	if mock.CreateFunc == nil {
		panic("DBMock.CreateFunc: method is nil but DB.Create was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Model any
	}{
		Ctx:   ctx,
		Model: model,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, model)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedDB.CreateCalls())
func (mock *DBMock) CreateCalls() []struct {
	Ctx   context.Context
	Model any
} {
	var calls []struct {
		Ctx   context.Context
		Model any
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Find calls FindFunc.
func (mock *DBMock) Find(ctx context.Context, dest any, where string, args ...any) error {
	if mock.FindFunc == nil {
		panic("DBMock.FindFunc: method is nil but DB.Find was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Dest  any
		Where string
		Args  []any
	}{
		Ctx:   ctx,
		Dest:  dest,
		Where: where,
		Args:  args,
	}
	mock.lockFind.Lock()
	mock.calls.Find = append(mock.calls.Find, callInfo)
	mock.lockFind.Unlock()
	return mock.FindFunc(ctx, dest, where, args...)
}

// FindCalls gets all the calls that were made to Find.
// Check the length with:
//     len(mockedDB.FindCalls())
func (mock *DBMock) FindCalls() []struct {
	Ctx   context.Context
	Dest  any
	Where string
	Args  []any
} {
	var calls []struct {
		Ctx   context.Context
		Dest  any
		Where string
		Args  []any
	}
	mock.lockFind.RLock()
	calls = mock.calls.Find
	mock.lockFind.RUnlock()
	return calls
}
