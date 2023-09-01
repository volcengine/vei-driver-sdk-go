// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	logger "github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
	interfaces "github.com/volcengine/vei-driver-sdk-go/pkg/interfaces"

	mock "github.com/stretchr/testify/mock"

	models "github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	pkgmodels "github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
)

// Driver is an autogenerated mock type for the Driver type
type Driver struct {
	mock.Mock
}

// HandleReadCommands provides a mock function with given fields: deviceName, protocols, reqs
func (_m *Driver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []pkgmodels.CommandRequest) ([]*pkgmodels.CommandValue, error) {
	ret := _m.Called(deviceName, protocols, reqs)

	var r0 []*pkgmodels.CommandValue
	if rf, ok := ret.Get(0).(func(string, map[string]models.ProtocolProperties, []pkgmodels.CommandRequest) []*pkgmodels.CommandValue); ok {
		r0 = rf(deviceName, protocols, reqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pkgmodels.CommandValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, map[string]models.ProtocolProperties, []pkgmodels.CommandRequest) error); ok {
		r1 = rf(deviceName, protocols, reqs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleServiceCall provides a mock function with given fields: deviceName, protocols, req, data
func (_m *Driver) HandleServiceCall(deviceName string, protocols map[string]models.ProtocolProperties, req pkgmodels.CommandRequest, data []byte) (*pkgmodels.CommandValue, error) {
	ret := _m.Called(deviceName, protocols, req, data)

	var r0 *pkgmodels.CommandValue
	if rf, ok := ret.Get(0).(func(string, map[string]models.ProtocolProperties, pkgmodels.CommandRequest, []byte) *pkgmodels.CommandValue); ok {
		r0 = rf(deviceName, protocols, req, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pkgmodels.CommandValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, map[string]models.ProtocolProperties, pkgmodels.CommandRequest, []byte) error); ok {
		r1 = rf(deviceName, protocols, req, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleWriteCommands provides a mock function with given fields: deviceName, protocols, reqs, params
func (_m *Driver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []pkgmodels.CommandRequest, params []*pkgmodels.CommandValue) error {
	ret := _m.Called(deviceName, protocols, reqs, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]models.ProtocolProperties, []pkgmodels.CommandRequest, []*pkgmodels.CommandValue) error); ok {
		r0 = rf(deviceName, protocols, reqs, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Initialize provides a mock function with given fields: lc, asyncCh, deviceCh, eventCallback
func (_m *Driver) Initialize(lc logger.LoggingClient, asyncCh chan<- *pkgmodels.AsyncValues, deviceCh chan<- []pkgmodels.DiscoveredDevice, eventCallback interfaces.EventCallback) error {
	ret := _m.Called(lc, asyncCh, deviceCh, eventCallback)

	var r0 error
	if rf, ok := ret.Get(0).(func(logger.LoggingClient, chan<- *pkgmodels.AsyncValues, chan<- []pkgmodels.DiscoveredDevice, interfaces.EventCallback) error); ok {
		r0 = rf(lc, asyncCh, deviceCh, eventCallback)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields: force
func (_m *Driver) Stop(force bool) error {
	ret := _m.Called(force)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(force)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDriver interface {
	mock.TestingT
	Cleanup(func())
}

// NewDriver creates a new instance of Driver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDriver(t mockConstructorTestingTNewDriver) *Driver {
	mock := &Driver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
