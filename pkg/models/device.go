/*
 * Copyright 2023 Beijing Volcano Engine Technology Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package models

import (
	contracts "github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// OperatingState is an indication of the operations of the device.
type OperatingState string

const (
	REACHABLE   OperatingState = "Reachable"
	UNREACHABLE OperatingState = "Unreachable"
	UP          OperatingState = "Up"
	DOWN        OperatingState = "Down"
	UNKNOWN     OperatingState = "Unknown"
)

// Device contains the necessary information of a device.
type Device struct {
	Name           string
	Protocols      map[string]contracts.ProtocolProperties
	OperatingState OperatingState
	Message        string
}

func WrapDevice(name string, protocols map[string]contracts.ProtocolProperties) *Device {
	return &Device{Name: name, Protocols: protocols}
}

// GetProtocolByName returns the protocol specified by name.
func (d *Device) GetProtocolByName(name string) (contracts.ProtocolProperties, bool) {
	protocol, exist := d.Protocols[name]
	return protocol, exist
}

// SetDeviceReachable set the device state to REACHABLE.
func (d *Device) SetDeviceReachable() {
	d.OperatingState = REACHABLE
}

// SetDeviceUnreachable set the device state to UNREACHABLE.
func (d *Device) SetDeviceUnreachable() {
	d.OperatingState = UNREACHABLE
}

// SetDeviceOperatingState supports to set the customized device state.
func (d *Device) SetDeviceOperatingState(state OperatingState, message string) {
	d.OperatingState, d.Message = state, message
}
