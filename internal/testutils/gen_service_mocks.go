// Code generated by MockGen. DO NOT EDIT.
// Source: ../../pkg/service/service.go

// Package testutils is a generated GoMock package.
package testutils

import (
	reflect "reflect"

	ssm "github.com/aws/aws-sdk-go/service/ssm"
	gomock "github.com/golang/mock/gomock"
)

// MockSSMClient is a mock of SSMClient interface.
type MockSSMClient struct {
	ctrl     *gomock.Controller
	recorder *MockSSMClientMockRecorder
}

// MockSSMClientMockRecorder is the mock recorder for MockSSMClient.
type MockSSMClientMockRecorder struct {
	mock *MockSSMClient
}

// NewMockSSMClient creates a new mock instance.
func NewMockSSMClient(ctrl *gomock.Controller) *MockSSMClient {
	mock := &MockSSMClient{ctrl: ctrl}
	mock.recorder = &MockSSMClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSSMClient) EXPECT() *MockSSMClientMockRecorder {
	return m.recorder
}

// GetParametersByPathPages mocks base method.
func (m *MockSSMClient) GetParametersByPathPages(getParametersByPathInput *ssm.GetParametersByPathInput, fn func(*ssm.GetParametersByPathOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParametersByPathPages", getParametersByPathInput, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetParametersByPathPages indicates an expected call of GetParametersByPathPages.
func (mr *MockSSMClientMockRecorder) GetParametersByPathPages(getParametersByPathInput, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParametersByPathPages", reflect.TypeOf((*MockSSMClient)(nil).GetParametersByPathPages), getParametersByPathInput, fn)
}

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetParameters mocks base method.
func (m *MockService) GetParameters(searchPath string, recursive bool) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParameters", searchPath, recursive)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParameters indicates an expected call of GetParameters.
func (mr *MockServiceMockRecorder) GetParameters(searchPath, recursive interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParameters", reflect.TypeOf((*MockService)(nil).GetParameters), searchPath, recursive)
}
