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

package status

import (
	"errors"
	"reflect"
	"testing"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/interfaces"
	"github.com/edgexfoundry/device-sdk-go/v2/pkg/interfaces/mocks"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
	"github.com/rcrowley/go-metrics"
)

func TestNewManagedDevice(t *testing.T) {
	mockDeviceService := &mocks.DeviceServiceSDK{}
	mockDeviceService.On("GetLoggingClient").Return(logger.NewMockClient())
	mockDeviceService.On("GetDeviceByName", "not-found").Return(models.Device{}, errors.New("device not found"))
	mockDeviceService.On("GetDeviceByName", "invalid-label").Return(models.Device{Labels: []string{LabelPrefix + "..."}}, nil)
	mockDeviceService.On("GetDeviceByName", "online").Return(models.Device{Labels: []string{LabelPrefix + `{"status":"Online"}`}}, nil)
	mockDeviceService.On("GetDeviceByName", "offline").Return(models.Device{Labels: []string{LabelPrefix + `{"status":"Offline"}`}}, nil)

	type args struct {
		deviceName string
		ds         interfaces.DeviceServiceSDK
	}
	tests := []struct {
		name string
		args args
		want *ManagedDevice
	}{
		{name: "not-found", args: args{deviceName: "not-found", ds: mockDeviceService}, want: &ManagedDevice{Name: "not-found", Status: Unknown, ConsecutiveErrorNum: metrics.NewCounter()}},
		{name: "unknown", args: args{deviceName: "invalid-label", ds: mockDeviceService}, want: &ManagedDevice{Name: "invalid-label", Status: Unknown, ConsecutiveErrorNum: metrics.NewCounter()}},
		{name: "online", args: args{deviceName: "online", ds: mockDeviceService}, want: &ManagedDevice{Name: "online", Status: Online, ConsecutiveErrorNum: metrics.NewCounter()}},
		{name: "offline", args: args{deviceName: "offline", ds: mockDeviceService}, want: &ManagedDevice{Name: "offline", Status: Offline, ConsecutiveErrorNum: metrics.NewCounter()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewManagedDevice(tt.args.deviceName, tt.args.ds)
			t.Logf("%+v", got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewManagedDevice() = %v, want %v", got, tt.want)
			}
		})
	}
}
