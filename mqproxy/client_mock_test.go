// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cubefs/blobstore/mqproxy/client (interfaces: Register)

// Package mqproxy is a generated GoMock package.
package mqproxy

import (
	context "context"
	reflect "reflect"

	client "github.com/cubefs/blobstore/mqproxy/client"
	gomock "github.com/golang/mock/gomock"
)

// MockRegister is a mock of Register interface.
type MockRegister struct {
	ctrl     *gomock.Controller
	recorder *MockRegisterMockRecorder
}

// MockRegisterMockRecorder is the mock recorder for MockRegister.
type MockRegisterMockRecorder struct {
	mock *MockRegister
}

// NewMockRegister creates a new mock instance.
func NewMockRegister(ctrl *gomock.Controller) *MockRegister {
	mock := &MockRegister{ctrl: ctrl}
	mock.recorder = &MockRegisterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegister) EXPECT() *MockRegisterMockRecorder {
	return m.recorder
}

// Register mocks base method.
func (m *MockRegister) Register(arg0 context.Context, arg1 client.RegisterInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockRegisterMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRegister)(nil).Register), arg0, arg1)
}
