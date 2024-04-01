// Code generated by mockery v2.40.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// StatusManager is an autogenerated mock type for the StatusManager type
type StatusManager struct {
	mock.Mock
}

// OnAddDevice provides a mock function with given fields: deviceName
func (_m *StatusManager) OnAddDevice(deviceName string) {
	_m.Called(deviceName)
}

// OnHandleCommandsFailed provides a mock function with given fields: deviceName, n
func (_m *StatusManager) OnHandleCommandsFailed(deviceName string, n int64) {
	_m.Called(deviceName, n)
}

// OnHandleCommandsSuccessfully provides a mock function with given fields: deviceName, n
func (_m *StatusManager) OnHandleCommandsSuccessfully(deviceName string, n int64) {
	_m.Called(deviceName, n)
}

// OnRemoveDevice provides a mock function with given fields: deviceName
func (_m *StatusManager) OnRemoveDevice(deviceName string) {
	_m.Called(deviceName)
}

// SetDeviceOffline provides a mock function with given fields: deviceName, reason
func (_m *StatusManager) SetDeviceOffline(deviceName string, reason string) {
	_m.Called(deviceName, reason)
}

// SetDeviceOnline provides a mock function with given fields: deviceName
func (_m *StatusManager) SetDeviceOnline(deviceName string) {
	_m.Called(deviceName)
}

// UpdateDeviceStatus provides a mock function with given fields: deviceName, status, reason
func (_m *StatusManager) UpdateDeviceStatus(deviceName string, status string, reason string) {
	_m.Called(deviceName, status, reason)
}

// NewStatusManager creates a new instance of StatusManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStatusManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *StatusManager {
	mock := &StatusManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
