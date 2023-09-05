// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-ebpf-sdk-go/pkg/xdp (interfaces: BpfXdp)

// Package mock_xdp is a generated GoMock package.
package mock_xdp

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBpfXdp is a mock of BpfXdp interface.
type MockBpfXdp struct {
	ctrl     *gomock.Controller
	recorder *MockBpfXdpMockRecorder
}

// MockBpfXdpMockRecorder is the mock recorder for MockBpfXdp.
type MockBpfXdpMockRecorder struct {
	mock *MockBpfXdp
}

// NewMockBpfXdp creates a new mock instance.
func NewMockBpfXdp(ctrl *gomock.Controller) *MockBpfXdp {
	mock := &MockBpfXdp{ctrl: ctrl}
	mock.recorder = &MockBpfXdpMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBpfXdp) EXPECT() *MockBpfXdpMockRecorder {
	return m.recorder
}

// XDPAttach mocks base method.
func (m *MockBpfXdp) XDPAttach(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XDPAttach", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// XDPAttach indicates an expected call of XDPAttach.
func (mr *MockBpfXdpMockRecorder) XDPAttach(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XDPAttach", reflect.TypeOf((*MockBpfXdp)(nil).XDPAttach), arg0)
}

// XDPDetach mocks base method.
func (m *MockBpfXdp) XDPDetach() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XDPDetach")
	ret0, _ := ret[0].(error)
	return ret0
}

// XDPDetach indicates an expected call of XDPDetach.
func (mr *MockBpfXdpMockRecorder) XDPDetach() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XDPDetach", reflect.TypeOf((*MockBpfXdp)(nil).XDPDetach))
}