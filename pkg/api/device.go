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

package api

import (
	"github.com/edgexfoundry/device-sdk-go/v2/pkg/service"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// Devices return all managed Devices from cache.
func Devices() []models.Device {
	return service.RunningService().Devices()
}

// GetDeviceByName returns the Device by its name if it exists in the cache, or returns an error.
func GetDeviceByName(name string) (models.Device, error) {
	return service.RunningService().GetDeviceByName(name)
}

// AddDevice adds a new Device to the Device Service and Core Metadata.
// Returns new Device id or non-nil error.
func AddDevice(device models.Device) (string, error) {
	return service.RunningService().AddDevice(device)
}

// UpdateDevice updates the Device in the cache and ensures that the
// copy in Core Metadata is also updated.
func UpdateDevice(device models.Device) error {
	return service.RunningService().UpdateDevice(device)
}

// RemoveDeviceByName removes the specified Device by name from the cache and ensures that the
// instance in Core Metadata is also removed.
func RemoveDeviceByName(name string) error {
	return service.RunningService().RemoveDeviceByName(name)
}
