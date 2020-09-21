// Code generated by MockGen. DO NOT EDIT.
// Source: adder.go

// Package adder is a generated GoMock package.
package adder

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAdder is a mock of Adder interface
type MockAdder struct {
	ctrl     *gomock.Controller
	recorder *MockAdderMockRecorder
}

// MockAdderMockRecorder is the mock recorder for MockAdder
type MockAdderMockRecorder struct {
	mock *MockAdder
}

// NewMockAdder creates a new mock instance
func NewMockAdder(ctrl *gomock.Controller) *MockAdder {
	mock := &MockAdder{ctrl: ctrl}
	mock.recorder = &MockAdderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdder) EXPECT() *MockAdderMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockAdder) Add(ctx context.Context, in *AddReq) (*AddResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, in)
	ret0, _ := ret[0].(*AddResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockAdderMockRecorder) Add(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockAdder)(nil).Add), ctx, in)
}