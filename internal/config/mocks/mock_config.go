// Code generated by MockGen. DO NOT EDIT.
// Source: internal/config/interface.go

// Package mock_config is a generated GoMock package.
package mock_config

import (
	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
	reflect "reflect"
)

// MockGeneralConfig is a mock of GeneralConfig interface
type MockGeneralConfig struct {
	ctrl     *gomock.Controller
	recorder *MockGeneralConfigMockRecorder
}

// MockGeneralConfigMockRecorder is the mock recorder for MockGeneralConfig
type MockGeneralConfigMockRecorder struct {
	mock *MockGeneralConfig
}

// NewMockGeneralConfig creates a new mock instance
func NewMockGeneralConfig(ctrl *gomock.Controller) *MockGeneralConfig {
	mock := &MockGeneralConfig{ctrl: ctrl}
	mock.recorder = &MockGeneralConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGeneralConfig) EXPECT() *MockGeneralConfigMockRecorder {
	return m.recorder
}

// GetKey mocks base method
func (m *MockGeneralConfig) GetKey(key string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKey", key)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// GetKey indicates an expected call of GetKey
func (mr *MockGeneralConfigMockRecorder) GetKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKey", reflect.TypeOf((*MockGeneralConfig)(nil).GetKey), key)
}

// GetLogLevel mocks base method
func (m *MockGeneralConfig) GetLogLevel() logrus.Level {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogLevel")
	ret0, _ := ret[0].(logrus.Level)
	return ret0
}

// GetLogLevel indicates an expected call of GetLogLevel
func (mr *MockGeneralConfigMockRecorder) GetLogLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogLevel", reflect.TypeOf((*MockGeneralConfig)(nil).GetLogLevel))
}