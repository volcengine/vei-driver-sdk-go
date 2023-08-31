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
	"encoding/json"
	"strings"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/interfaces"
	"github.com/rcrowley/go-metrics"
)

type DeviceStatus string

const (
	Online  DeviceStatus = "Online"
	Offline DeviceStatus = "Offline"
	Unknown DeviceStatus = "Unknown"
)

const LabelPrefix = "vei-status:"

type ManagedDevice struct {
	Name                string          `json:"name"`
	Status              DeviceStatus    `json:"status"`
	OnlineTime          int64           `json:"online_time"`
	OfflineTime         int64           `json:"offline_time"`
	ConsecutiveErrorNum metrics.Counter `json:"-"`
}

func NewManagedDevice(deviceName string, ds interfaces.DeviceServiceSDK) *ManagedDevice {
	managedDevice := &ManagedDevice{Name: deviceName, Status: Unknown, ConsecutiveErrorNum: metrics.NewCounter()}

	cachedDevice, err := ds.GetDeviceByName(deviceName)
	if err != nil {
		return managedDevice
	}

	for _, label := range cachedDevice.Labels {
		if strings.HasPrefix(label, LabelPrefix) {
			if err = json.Unmarshal([]byte(label[len(LabelPrefix):]), managedDevice); err != nil {
				ds.GetLoggingClient().Errorf("failed to unmarshal device status in label")
				return managedDevice
			}
		}
	}
	return managedDevice
}
