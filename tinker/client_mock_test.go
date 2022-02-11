// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cubefs/blobstore/tinker/client (interfaces: ClusterMgrAPI,IScheduler,BlobnodeAPI,IWorker)

// Package tinker is a generated GoMock package.
package tinker

import (
	context "context"
	reflect "reflect"

	proto "github.com/cubefs/blobstore/common/proto"
	client "github.com/cubefs/blobstore/tinker/client"
	gomock "github.com/golang/mock/gomock"
)

// MockClusterMgrAPI is a mock of ClusterMgrAPI interface.
type MockClusterMgrAPI struct {
	ctrl     *gomock.Controller
	recorder *MockClusterMgrAPIMockRecorder
}

// MockClusterMgrAPIMockRecorder is the mock recorder for MockClusterMgrAPI.
type MockClusterMgrAPIMockRecorder struct {
	mock *MockClusterMgrAPI
}

// NewMockClusterMgrAPI creates a new mock instance.
func NewMockClusterMgrAPI(ctrl *gomock.Controller) *MockClusterMgrAPI {
	mock := &MockClusterMgrAPI{ctrl: ctrl}
	mock.recorder = &MockClusterMgrAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterMgrAPI) EXPECT() *MockClusterMgrAPIMockRecorder {
	return m.recorder
}

// GetConfig mocks base method.
func (m *MockClusterMgrAPI) GetConfig(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockClusterMgrAPIMockRecorder) GetConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockClusterMgrAPI)(nil).GetConfig), arg0, arg1)
}

// GetVolInfo mocks base method.
func (m *MockClusterMgrAPI) GetVolInfo(arg0 context.Context, arg1 proto.Vid) (client.VolInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolInfo", arg0, arg1)
	ret0, _ := ret[0].(client.VolInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolInfo indicates an expected call of GetVolInfo.
func (mr *MockClusterMgrAPIMockRecorder) GetVolInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolInfo", reflect.TypeOf((*MockClusterMgrAPI)(nil).GetVolInfo), arg0, arg1)
}

// ListVolume mocks base method.
func (m *MockClusterMgrAPI) ListVolume(arg0 context.Context, arg1 proto.Vid, arg2 int) ([]client.VolInfo, proto.Vid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].([]client.VolInfo)
	ret1, _ := ret[1].(proto.Vid)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListVolume indicates an expected call of ListVolume.
func (mr *MockClusterMgrAPIMockRecorder) ListVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVolume", reflect.TypeOf((*MockClusterMgrAPI)(nil).ListVolume), arg0, arg1, arg2)
}

// MockScheduler is a mock of IScheduler interface.
type MockScheduler struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerMockRecorder
}

// MockSchedulerMockRecorder is the mock recorder for MockScheduler.
type MockSchedulerMockRecorder struct {
	mock *MockScheduler
}

// NewMockScheduler creates a new mock instance.
func NewMockScheduler(ctrl *gomock.Controller) *MockScheduler {
	mock := &MockScheduler{ctrl: ctrl}
	mock.recorder = &MockSchedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduler) EXPECT() *MockSchedulerMockRecorder {
	return m.recorder
}

// DeleteService mocks base method.
func (m *MockScheduler) DeleteService(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService.
func (mr *MockSchedulerMockRecorder) DeleteService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockScheduler)(nil).DeleteService), arg0, arg1)
}

// ListService mocks base method.
func (m *MockScheduler) ListService(arg0 context.Context, arg1 proto.ClusterID, arg2 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListService", arg0, arg1, arg2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListService indicates an expected call of ListService.
func (mr *MockSchedulerMockRecorder) ListService(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListService", reflect.TypeOf((*MockScheduler)(nil).ListService), arg0, arg1, arg2)
}

// Register mocks base method.
func (m *MockScheduler) Register(arg0 context.Context, arg1 proto.ClusterID, arg2, arg3, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockSchedulerMockRecorder) Register(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockScheduler)(nil).Register), arg0, arg1, arg2, arg3, arg4)
}

// MockBlobnodeAPI is a mock of BlobnodeAPI interface.
type MockBlobnodeAPI struct {
	ctrl     *gomock.Controller
	recorder *MockBlobnodeAPIMockRecorder
}

// MockBlobnodeAPIMockRecorder is the mock recorder for MockBlobnodeAPI.
type MockBlobnodeAPIMockRecorder struct {
	mock *MockBlobnodeAPI
}

// NewMockBlobnodeAPI creates a new mock instance.
func NewMockBlobnodeAPI(ctrl *gomock.Controller) *MockBlobnodeAPI {
	mock := &MockBlobnodeAPI{ctrl: ctrl}
	mock.recorder = &MockBlobnodeAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlobnodeAPI) EXPECT() *MockBlobnodeAPIMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockBlobnodeAPI) Delete(arg0 context.Context, arg1 proto.VunitLocation, arg2 proto.BlobID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBlobnodeAPIMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBlobnodeAPI)(nil).Delete), arg0, arg1, arg2)
}

// MarkDelete mocks base method.
func (m *MockBlobnodeAPI) MarkDelete(arg0 context.Context, arg1 proto.VunitLocation, arg2 proto.BlobID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkDelete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkDelete indicates an expected call of MarkDelete.
func (mr *MockBlobnodeAPIMockRecorder) MarkDelete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkDelete", reflect.TypeOf((*MockBlobnodeAPI)(nil).MarkDelete), arg0, arg1, arg2)
}

// MockWorkerCli is a mock of IWorker interface.
type MockWorkerCli struct {
	ctrl     *gomock.Controller
	recorder *MockWorkerCliMockRecorder
}

// MockWorkerCliMockRecorder is the mock recorder for MockWorkerCli.
type MockWorkerCliMockRecorder struct {
	mock *MockWorkerCli
}

// NewMockWorkerCli creates a new mock instance.
func NewMockWorkerCli(ctrl *gomock.Controller) *MockWorkerCli {
	mock := &MockWorkerCli{ctrl: ctrl}
	mock.recorder = &MockWorkerCliMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkerCli) EXPECT() *MockWorkerCliMockRecorder {
	return m.recorder
}

// RepairShard mocks base method.
func (m *MockWorkerCli) RepairShard(arg0 context.Context, arg1 string, arg2 proto.ShardRepairTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RepairShard", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RepairShard indicates an expected call of RepairShard.
func (mr *MockWorkerCliMockRecorder) RepairShard(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RepairShard", reflect.TypeOf((*MockWorkerCli)(nil).RepairShard), arg0, arg1, arg2)
}
