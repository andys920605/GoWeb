// Code generated by MockGen. DO NOT EDIT.
// Source: GoWeb/repository/interface (interfaces: IMemberRep)

// Package mock is a generated GoMock package.
package mock

import (
	repository "GoWeb/models/repository"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIMemberRep is a mock of IMemberRep interface.
type MockIMemberRep struct {
	ctrl     *gomock.Controller
	recorder *MockIMemberRepMockRecorder
}

// MockIMemberRepMockRecorder is the mock recorder for MockIMemberRep.
type MockIMemberRepMockRecorder struct {
	mock *MockIMemberRep
}

// NewMockIMemberRep creates a new mock instance.
func NewMockIMemberRep(ctrl *gomock.Controller) *MockIMemberRep {
	mock := &MockIMemberRep{ctrl: ctrl}
	mock.recorder = &MockIMemberRepMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMemberRep) EXPECT() *MockIMemberRepMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIMemberRep) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockIMemberRepMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIMemberRep)(nil).Close))
}

// Disable mocks base method.
func (m *MockIMemberRep) Disable(arg0 context.Context, arg1 *repository.UpdateMember) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disable", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Disable indicates an expected call of Disable.
func (mr *MockIMemberRepMockRecorder) Disable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*MockIMemberRep)(nil).Disable), arg0, arg1)
}

// Find mocks base method.
func (m *MockIMemberRep) Find(arg0 context.Context, arg1, arg2 *string) (*repository.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1, arg2)
	ret0, _ := ret[0].(*repository.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockIMemberRepMockRecorder) Find(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockIMemberRep)(nil).Find), arg0, arg1, arg2)
}

// FindAll mocks base method.
func (m *MockIMemberRep) FindAll(arg0 context.Context) (*[]repository.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].(*[]repository.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockIMemberRepMockRecorder) FindAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockIMemberRep)(nil).FindAll), arg0)
}

// Insert mocks base method.
func (m *MockIMemberRep) Insert(arg0 context.Context, arg1 *repository.Member) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockIMemberRepMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIMemberRep)(nil).Insert), arg0, arg1)
}

// Updates mocks base method.
func (m *MockIMemberRep) Updates(arg0 context.Context, arg1 *repository.UpdateMember) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Updates", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Updates indicates an expected call of Updates.
func (mr *MockIMemberRepMockRecorder) Updates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Updates", reflect.TypeOf((*MockIMemberRep)(nil).Updates), arg0, arg1)
}
