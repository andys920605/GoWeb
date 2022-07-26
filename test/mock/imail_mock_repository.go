// Code generated by MockGen. DO NOT EDIT.
// Source: GoWeb/repository/interface (interfaces: IMailRep)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
)

// MockIMailRep is a mock of IMailRep interface.
type MockIMailRep struct {
	ctrl     *gomock.Controller
	recorder *MockIMailRepMockRecorder
}

// MockIMailRepMockRecorder is the mock recorder for MockIMailRep.
type MockIMailRepMockRecorder struct {
	mock *MockIMailRep
}

// NewMockIMailRep creates a new mock instance.
func NewMockIMailRep(ctrl *gomock.Controller) *MockIMailRep {
	mock := &MockIMailRep{ctrl: ctrl}
	mock.recorder = &MockIMailRepMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMailRep) EXPECT() *MockIMailRepMockRecorder {
	return m.recorder
}
