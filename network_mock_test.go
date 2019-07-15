// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/choria-io/go-network-broker (interfaces: ChoriaFramework,BuildInfoProvider)

// Package network is a generated GoMock package.
package network

import (
	tls "crypto/tls"
	go_config "github.com/choria-io/go-config"
	go_srvcache "github.com/choria-io/go-srvcache"
	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
	reflect "reflect"
)

// MockChoriaFramework is a mock of ChoriaFramework interface
type MockChoriaFramework struct {
	ctrl     *gomock.Controller
	recorder *MockChoriaFrameworkMockRecorder
}

// MockChoriaFrameworkMockRecorder is the mock recorder for MockChoriaFramework
type MockChoriaFrameworkMockRecorder struct {
	mock *MockChoriaFramework
}

// NewMockChoriaFramework creates a new mock instance
func NewMockChoriaFramework(ctrl *gomock.Controller) *MockChoriaFramework {
	mock := &MockChoriaFramework{ctrl: ctrl}
	mock.recorder = &MockChoriaFrameworkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChoriaFramework) EXPECT() *MockChoriaFrameworkMockRecorder {
	return m.recorder
}

// Configuration mocks base method
func (m *MockChoriaFramework) Configuration() *go_config.Config {
	ret := m.ctrl.Call(m, "Configuration")
	ret0, _ := ret[0].(*go_config.Config)
	return ret0
}

// Configuration indicates an expected call of Configuration
func (mr *MockChoriaFrameworkMockRecorder) Configuration() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configuration", reflect.TypeOf((*MockChoriaFramework)(nil).Configuration))
}

// Logger mocks base method
func (m *MockChoriaFramework) Logger(arg0 string) *logrus.Entry {
	ret := m.ctrl.Call(m, "Logger", arg0)
	ret0, _ := ret[0].(*logrus.Entry)
	return ret0
}

// Logger indicates an expected call of Logger
func (mr *MockChoriaFrameworkMockRecorder) Logger(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockChoriaFramework)(nil).Logger), arg0)
}

// NetworkBrokerPeers mocks base method
func (m *MockChoriaFramework) NetworkBrokerPeers() (go_srvcache.Servers, error) {
	ret := m.ctrl.Call(m, "NetworkBrokerPeers")
	ret0, _ := ret[0].(go_srvcache.Servers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkBrokerPeers indicates an expected call of NetworkBrokerPeers
func (mr *MockChoriaFrameworkMockRecorder) NetworkBrokerPeers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkBrokerPeers", reflect.TypeOf((*MockChoriaFramework)(nil).NetworkBrokerPeers))
}

// TLSConfig mocks base method
func (m *MockChoriaFramework) TLSConfig() (*tls.Config, error) {
	ret := m.ctrl.Call(m, "TLSConfig")
	ret0, _ := ret[0].(*tls.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TLSConfig indicates an expected call of TLSConfig
func (mr *MockChoriaFrameworkMockRecorder) TLSConfig() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TLSConfig", reflect.TypeOf((*MockChoriaFramework)(nil).TLSConfig))
}

// ValidateSecurity mocks base method
func (m *MockChoriaFramework) ValidateSecurity() ([]string, bool) {
	ret := m.ctrl.Call(m, "ValidateSecurity")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ValidateSecurity indicates an expected call of ValidateSecurity
func (mr *MockChoriaFrameworkMockRecorder) ValidateSecurity() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateSecurity", reflect.TypeOf((*MockChoriaFramework)(nil).ValidateSecurity))
}

// MockBuildInfoProvider is a mock of BuildInfoProvider interface
type MockBuildInfoProvider struct {
	ctrl     *gomock.Controller
	recorder *MockBuildInfoProviderMockRecorder
}

// MockBuildInfoProviderMockRecorder is the mock recorder for MockBuildInfoProvider
type MockBuildInfoProviderMockRecorder struct {
	mock *MockBuildInfoProvider
}

// NewMockBuildInfoProvider creates a new mock instance
func NewMockBuildInfoProvider(ctrl *gomock.Controller) *MockBuildInfoProvider {
	mock := &MockBuildInfoProvider{ctrl: ctrl}
	mock.recorder = &MockBuildInfoProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBuildInfoProvider) EXPECT() *MockBuildInfoProviderMockRecorder {
	return m.recorder
}

// MaxBrokerClients mocks base method
func (m *MockBuildInfoProvider) MaxBrokerClients() int {
	ret := m.ctrl.Call(m, "MaxBrokerClients")
	ret0, _ := ret[0].(int)
	return ret0
}

// MaxBrokerClients indicates an expected call of MaxBrokerClients
func (mr *MockBuildInfoProviderMockRecorder) MaxBrokerClients() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxBrokerClients", reflect.TypeOf((*MockBuildInfoProvider)(nil).MaxBrokerClients))
}
