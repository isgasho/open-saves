// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/metadb/metadb.go

// Package mock_metadb is a generated GoMock package.
package mock_metadb

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	metadb "github.com/googleforgames/triton/internal/pkg/metadb"
	reflect "reflect"
	time "time"
)

// MockDriver is a mock of Driver interface
type MockDriver struct {
	ctrl     *gomock.Controller
	recorder *MockDriverMockRecorder
}

// MockDriverMockRecorder is the mock recorder for MockDriver
type MockDriverMockRecorder struct {
	mock *MockDriver
}

// NewMockDriver creates a new mock instance
func NewMockDriver(ctrl *gomock.Controller) *MockDriver {
	mock := &MockDriver{ctrl: ctrl}
	mock.recorder = &MockDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDriver) EXPECT() *MockDriverMockRecorder {
	return m.recorder
}

// Connect mocks base method
func (m *MockDriver) Connect(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect
func (mr *MockDriverMockRecorder) Connect(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockDriver)(nil).Connect), ctx)
}

// Disconnect mocks base method
func (m *MockDriver) Disconnect(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disconnect", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnect indicates an expected call of Disconnect
func (mr *MockDriverMockRecorder) Disconnect(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockDriver)(nil).Disconnect), ctx)
}

// CreateStore mocks base method
func (m *MockDriver) CreateStore(ctx context.Context, store *metadb.Store) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStore", ctx, store)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStore indicates an expected call of CreateStore
func (mr *MockDriverMockRecorder) CreateStore(ctx, store interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStore", reflect.TypeOf((*MockDriver)(nil).CreateStore), ctx, store)
}

// GetStore mocks base method
func (m *MockDriver) GetStore(ctx context.Context, key string) (*metadb.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStore", ctx, key)
	ret0, _ := ret[0].(*metadb.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStore indicates an expected call of GetStore
func (mr *MockDriverMockRecorder) GetStore(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStore", reflect.TypeOf((*MockDriver)(nil).GetStore), ctx, key)
}

// FindStoreByName mocks base method
func (m *MockDriver) FindStoreByName(ctx context.Context, name string) (*metadb.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindStoreByName", ctx, name)
	ret0, _ := ret[0].(*metadb.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindStoreByName indicates an expected call of FindStoreByName
func (mr *MockDriverMockRecorder) FindStoreByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindStoreByName", reflect.TypeOf((*MockDriver)(nil).FindStoreByName), ctx, name)
}

// DeleteStore mocks base method
func (m *MockDriver) DeleteStore(ctx context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStore", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStore indicates an expected call of DeleteStore
func (mr *MockDriverMockRecorder) DeleteStore(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStore", reflect.TypeOf((*MockDriver)(nil).DeleteStore), ctx, key)
}

// InsertRecord mocks base method
func (m *MockDriver) InsertRecord(ctx context.Context, storeKey string, record *metadb.Record) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertRecord", ctx, storeKey, record)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertRecord indicates an expected call of InsertRecord
func (mr *MockDriverMockRecorder) InsertRecord(ctx, storeKey, record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertRecord", reflect.TypeOf((*MockDriver)(nil).InsertRecord), ctx, storeKey, record)
}

// UpdateRecord mocks base method
func (m *MockDriver) UpdateRecord(ctx context.Context, storeKey string, record *metadb.Record) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRecord", ctx, storeKey, record)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRecord indicates an expected call of UpdateRecord
func (mr *MockDriverMockRecorder) UpdateRecord(ctx, storeKey, record interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRecord", reflect.TypeOf((*MockDriver)(nil).UpdateRecord), ctx, storeKey, record)
}

// GetRecord mocks base method
func (m *MockDriver) GetRecord(ctx context.Context, storeKey, key string) (*metadb.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecord", ctx, storeKey, key)
	ret0, _ := ret[0].(*metadb.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecord indicates an expected call of GetRecord
func (mr *MockDriverMockRecorder) GetRecord(ctx, storeKey, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecord", reflect.TypeOf((*MockDriver)(nil).GetRecord), ctx, storeKey, key)
}

// DeleteRecord mocks base method
func (m *MockDriver) DeleteRecord(ctx context.Context, storeKey, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecord", ctx, storeKey, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRecord indicates an expected call of DeleteRecord
func (mr *MockDriverMockRecorder) DeleteRecord(ctx, storeKey, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecord", reflect.TypeOf((*MockDriver)(nil).DeleteRecord), ctx, storeKey, key)
}

// TimestampPrecision mocks base method
func (m *MockDriver) TimestampPrecision() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimestampPrecision")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// TimestampPrecision indicates an expected call of TimestampPrecision
func (mr *MockDriverMockRecorder) TimestampPrecision() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimestampPrecision", reflect.TypeOf((*MockDriver)(nil).TimestampPrecision))
}
