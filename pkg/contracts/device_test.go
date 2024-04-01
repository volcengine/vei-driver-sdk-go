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
	"reflect"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
	"github.com/stretchr/testify/require"
)

func TestOperatingState_String(t *testing.T) {
	tests := []struct {
		name string
		s    OperatingState
		want string
	}{
		{s: REACHABLE, want: "Reachable"},
		{s: UNREACHABLE, want: "Unreachable"},
		{s: UP, want: "Up"},
		{s: DOWN, want: "Down"},
		{s: UNKNOWN, want: "Unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrapDevice(t *testing.T) {
	type args struct {
		name      string
		protocols map[string]models.ProtocolProperties
	}
	tests := []struct {
		name string
		args args
		want *Device
	}{
		{args: args{name: "device", protocols: map[string]models.ProtocolProperties{}}, want: &Device{Name: "device", Protocols: map[string]models.ProtocolProperties{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WrapDevice(tt.args.name, tt.args.protocols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapDevice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDevice_GetProtocolByName(t *testing.T) {
	device := WrapDevice("device", map[string]models.ProtocolProperties{
		"modbus-tcp": map[string]string{},
	})
	protocol, exist := device.GetProtocolByName("modbus-tcp")
	require.True(t, exist)
	require.NotNil(t, protocol)
	protocol, exist = device.GetProtocolByName("modbus-rtu")
	require.False(t, exist)
	require.Nil(t, protocol)
}

func TestDevice_SetOperatingState(t *testing.T) {
	device := WrapDevice("device", nil)

	device.SetStateUp()
	require.Equal(t, UP, device.OperatingState)

	device.SetStateDown()
	require.Equal(t, DOWN, device.OperatingState)

	device.SetStateReachable()
	require.Equal(t, REACHABLE, device.OperatingState)

	device.SetStateUnreachable()
	require.Equal(t, UNREACHABLE, device.OperatingState)

	state, message := UP, "error message"
	device.SetOperatingState(state, message)
	require.Equal(t, state, device.OperatingState)
	require.Equal(t, message, device.Message)
}
