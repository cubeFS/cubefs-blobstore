// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cubefs/blobstore/tinker/base (interfaces: IConsumer,IProducer,IVolumeCache,IBaseMgr)

// Package tinker is a generated GoMock package.
package tinker

import (
	context "context"
	reflect "reflect"

	sarama "github.com/Shopify/sarama"
	proto "github.com/cubefs/blobstore/common/proto"
	client "github.com/cubefs/blobstore/tinker/client"
	gomock "github.com/golang/mock/gomock"
)

// MockConsumer is a mock of IConsumer interface.
type MockConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerMockRecorder
}

// MockConsumerMockRecorder is the mock recorder for MockConsumer.
type MockConsumerMockRecorder struct {
	mock *MockConsumer
}

// NewMockConsumer creates a new mock instance.
func NewMockConsumer(ctrl *gomock.Controller) *MockConsumer {
	mock := &MockConsumer{ctrl: ctrl}
	mock.recorder = &MockConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsumer) EXPECT() *MockConsumerMockRecorder {
	return m.recorder
}

// CommitOffset mocks base method.
func (m *MockConsumer) CommitOffset(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitOffset", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitOffset indicates an expected call of CommitOffset.
func (mr *MockConsumerMockRecorder) CommitOffset(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitOffset", reflect.TypeOf((*MockConsumer)(nil).CommitOffset), arg0)
}

// ConsumeMessages mocks base method.
func (m *MockConsumer) ConsumeMessages(arg0 context.Context, arg1 int) []*sarama.ConsumerMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumeMessages", arg0, arg1)
	ret0, _ := ret[0].([]*sarama.ConsumerMessage)
	return ret0
}

// ConsumeMessages indicates an expected call of ConsumeMessages.
func (mr *MockConsumerMockRecorder) ConsumeMessages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeMessages", reflect.TypeOf((*MockConsumer)(nil).ConsumeMessages), arg0, arg1)
}

// MockProducer is a mock of IProducer interface.
type MockProducer struct {
	ctrl     *gomock.Controller
	recorder *MockProducerMockRecorder
}

// MockProducerMockRecorder is the mock recorder for MockProducer.
type MockProducerMockRecorder struct {
	mock *MockProducer
}

// NewMockProducer creates a new mock instance.
func NewMockProducer(ctrl *gomock.Controller) *MockProducer {
	mock := &MockProducer{ctrl: ctrl}
	mock.recorder = &MockProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducer) EXPECT() *MockProducerMockRecorder {
	return m.recorder
}

// SendMessage mocks base method.
func (m *MockProducer) SendMessage(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockProducerMockRecorder) SendMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockProducer)(nil).SendMessage), arg0)
}

// SendMessages mocks base method.
func (m *MockProducer) SendMessages(arg0 [][]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessages", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessages indicates an expected call of SendMessages.
func (mr *MockProducerMockRecorder) SendMessages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessages", reflect.TypeOf((*MockProducer)(nil).SendMessages), arg0)
}

// MockVolumeCache is a mock of IVolumeCache interface.
type MockVolumeCache struct {
	ctrl     *gomock.Controller
	recorder *MockVolumeCacheMockRecorder
}

// MockVolumeCacheMockRecorder is the mock recorder for MockVolumeCache.
type MockVolumeCacheMockRecorder struct {
	mock *MockVolumeCache
}

// NewMockVolumeCache creates a new mock instance.
func NewMockVolumeCache(ctrl *gomock.Controller) *MockVolumeCache {
	mock := &MockVolumeCache{ctrl: ctrl}
	mock.recorder = &MockVolumeCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVolumeCache) EXPECT() *MockVolumeCacheMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockVolumeCache) Get(arg0 proto.Vid) (*client.VolInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*client.VolInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockVolumeCacheMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVolumeCache)(nil).Get), arg0)
}

// Load mocks base method.
func (m *MockVolumeCache) Load() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load")
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockVolumeCacheMockRecorder) Load() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockVolumeCache)(nil).Load))
}

// Update mocks base method.
func (m *MockVolumeCache) Update(arg0 proto.Vid) (*client.VolInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*client.VolInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockVolumeCacheMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVolumeCache)(nil).Update), arg0)
}

// MockBaseMgr is a mock of IBaseMgr interface.
type MockBaseMgr struct {
	ctrl     *gomock.Controller
	recorder *MockBaseMgrMockRecorder
}

// MockBaseMgrMockRecorder is the mock recorder for MockBaseMgr.
type MockBaseMgrMockRecorder struct {
	mock *MockBaseMgr
}

// NewMockBaseMgr creates a new mock instance.
func NewMockBaseMgr(ctrl *gomock.Controller) *MockBaseMgr {
	mock := &MockBaseMgr{ctrl: ctrl}
	mock.recorder = &MockBaseMgrMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBaseMgr) EXPECT() *MockBaseMgrMockRecorder {
	return m.recorder
}

// Enabled mocks base method.
func (m *MockBaseMgr) Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled.
func (mr *MockBaseMgrMockRecorder) Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockBaseMgr)(nil).Enabled))
}

// GetErrorStats mocks base method.
func (m *MockBaseMgr) GetErrorStats() ([]string, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetErrorStats")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// GetErrorStats indicates an expected call of GetErrorStats.
func (mr *MockBaseMgrMockRecorder) GetErrorStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetErrorStats", reflect.TypeOf((*MockBaseMgr)(nil).GetErrorStats))
}

// GetTaskStats mocks base method.
func (m *MockBaseMgr) GetTaskStats() ([20]int, [20]int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskStats")
	ret0, _ := ret[0].([20]int)
	ret1, _ := ret[1].([20]int)
	return ret0, ret1
}

// GetTaskStats indicates an expected call of GetTaskStats.
func (mr *MockBaseMgrMockRecorder) GetTaskStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskStats", reflect.TypeOf((*MockBaseMgr)(nil).GetTaskStats))
}

// RunTask mocks base method.
func (m *MockBaseMgr) RunTask() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RunTask")
}

// RunTask indicates an expected call of RunTask.
func (mr *MockBaseMgrMockRecorder) RunTask() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTask", reflect.TypeOf((*MockBaseMgr)(nil).RunTask))
}
