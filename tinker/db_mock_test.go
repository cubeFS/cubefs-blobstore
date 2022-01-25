// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cubefs/blobstore/tinker/db (interfaces: IOrphanedShardTbl)

// Package tinker is a generated GoMock package.
package tinker

import (
	context "context"
	reflect "reflect"

	db "github.com/cubefs/blobstore/tinker/db"
	gomock "github.com/golang/mock/gomock"
)

// MockOrphanedShardTbl is a mock of IOrphanedShardTbl interface.
type MockOrphanedShardTbl struct {
	ctrl     *gomock.Controller
	recorder *MockOrphanedShardTblMockRecorder
}

// MockOrphanedShardTblMockRecorder is the mock recorder for MockOrphanedShardTbl.
type MockOrphanedShardTblMockRecorder struct {
	mock *MockOrphanedShardTbl
}

// NewMockOrphanedShardTbl creates a new mock instance.
func NewMockOrphanedShardTbl(ctrl *gomock.Controller) *MockOrphanedShardTbl {
	mock := &MockOrphanedShardTbl{ctrl: ctrl}
	mock.recorder = &MockOrphanedShardTblMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrphanedShardTbl) EXPECT() *MockOrphanedShardTblMockRecorder {
	return m.recorder
}

// SaveOrphanedShard mocks base method.
func (m *MockOrphanedShardTbl) SaveOrphanedShard(arg0 context.Context, arg1 db.ShardInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveOrphanedShard", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveOrphanedShard indicates an expected call of SaveOrphanedShard.
func (mr *MockOrphanedShardTblMockRecorder) SaveOrphanedShard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveOrphanedShard", reflect.TypeOf((*MockOrphanedShardTbl)(nil).SaveOrphanedShard), arg0, arg1)
}
