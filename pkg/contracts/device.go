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

package contracts

import (
	"errors"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// OperatingState is an indication of the operations of the device.
type OperatingState string

const (
	UP          OperatingState = "Up"
	DOWN        OperatingState = "Down"
	UNKNOWN     OperatingState = "Unknown"
	REACHABLE   OperatingState = "Reachable"
	UNREACHABLE OperatingState = "Unreachable"
)

func (s OperatingState) String() string {
	return string(s)
}

// Device contains the necessary information of a device.
type Device struct {
	Name           string                               `json:"name"`
	Protocols      map[string]models.ProtocolProperties `json:"protocols"`
	OperatingState OperatingState                       `json:"operating_state,omitempty"`
	Message        string                               `json:"message,omitempty"`
}

func WrapDevice(name string, protocols map[string]models.ProtocolProperties) *Device {
	return &Device{Name: name, Protocols: protocols}
}

// GetProtocolByName returns the protocol specified by name.
func (d *Device) GetProtocolByName(name string) (models.ProtocolProperties, bool) {
	protocol, exist := d.Protocols[name]
	return protocol, exist
}

// SetStateUp set the device state to UP.
func (d *Device) SetStateUp() {
	d.OperatingState = UP
}

// SetStateDown set the device state to DOWN.
func (d *Device) SetStateDown() {
	d.OperatingState = DOWN
}

// SetOperatingState supports to set the device state customarily.
func (d *Device) SetOperatingState(state OperatingState, message string) {
	d.OperatingState, d.Message = state, message
}

func (d *Device) UpdateStateByError(raw error) {
	var err *Error
	if errors.As(raw, &err) {
		d.OperatingState = DOWN
		d.Message = err.Error()
	}
}
