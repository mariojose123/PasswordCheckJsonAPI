// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\mario\go\pkg\ProjectStudioSol\interfaces\serviceInterface.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	context "context"
	structJson "passwordcheck/internal/structJson"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPWService is a mock of PWService interface.
type MockPWService struct {
	ctrl     *gomock.Controller
	recorder *MockPWServiceMockRecorder
}

// MockPWServiceMockRecorder is the mock recorder for MockPWService.
type MockPWServiceMockRecorder struct {
	mock *MockPWService
}

// NewMockPWService creates a new mock instance.
func NewMockPWService(ctrl *gomock.Controller) *MockPWService {
	mock := &MockPWService{ctrl: ctrl}
	mock.recorder = &MockPWServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPWService) EXPECT() *MockPWServiceMockRecorder {
	return m.recorder
}

// CheckPW mocks base method.
func (m *MockPWService) CheckPW(ctx context.Context, jsonStructure structJson.PSReceiveStructure) (bool, []string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPW", ctx, jsonStructure)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]string)
	return ret0, ret1
}

// CheckPW indicates an expected call of CheckPW.
func (mr *MockPWServiceMockRecorder) CheckPW(ctx, jsonStructure interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPW", reflect.TypeOf((*MockPWService)(nil).CheckPW), ctx, jsonStructure)
}