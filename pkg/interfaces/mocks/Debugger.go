// Code generated by mockery v2.40.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Debugger is an autogenerated mock type for the Debugger type
type Debugger struct {
	mock.Mock
}

// Debug provides a mock function with given fields:
func (_m *Debugger) Debug() {
	_m.Called()
}

// NewDebugger creates a new instance of Debugger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDebugger(t interface {
	mock.TestingT
	Cleanup(func())
}) *Debugger {
	mock := &Debugger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
