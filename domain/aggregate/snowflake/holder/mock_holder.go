// Code generated by MockGen. DO NOT EDIT.
// Source: holder.go

// Package snowflake is a generated GoMock package.
package holder

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHolder is a mock of Holder interface
type MockHolder struct {
	ctrl     *gomock.Controller
	recorder *MockHolderMockRecorder
}

// MockHolderMockRecorder is the mock recorder for MockHolder
type MockHolderMockRecorder struct {
	mock *MockHolder
}

// NewMockHolder creates a new mock instance
func NewMockHolder(ctrl *gomock.Controller) *MockHolder {
	mock := &MockHolder{ctrl: ctrl}
	mock.recorder = &MockHolderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHolder) EXPECT() *MockHolderMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockHolder) Init(addrs []string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", addrs)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockHolderMockRecorder) Init(addrs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockHolder)(nil).Init), addrs)
}

// GetWorkerId mocks base method
func (m *MockHolder) GetWorkerId(ip string) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkerId", ip)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetWorkerId indicates an expected call of GetWorkerId
func (mr *MockHolderMockRecorder) GetWorkerId(ip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkerId", reflect.TypeOf((*MockHolder)(nil).GetWorkerId), ip)
}
