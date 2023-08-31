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

// DeviceProfiles return all managed DeviceProfiles from cache
func DeviceProfiles() []models.DeviceProfile {
	return service.RunningService().DeviceProfiles()
}

// GetProfileByName returns the Profile by its name if it exists in the cache, or returns an error.
func GetProfileByName(name string) (models.DeviceProfile, error) {
	return service.RunningService().GetProfileByName(name)
}

// AddDeviceProfile adds a new DeviceProfile to the Device Service and Core Metadata.
// Returns new DeviceProfile id or non-nil error.
func AddDeviceProfile(profile models.DeviceProfile) (string, error) {
	return service.RunningService().AddDeviceProfile(profile)
}

// UpdateDeviceProfile updates the DeviceProfile in the cache and ensures that the
// copy in Core Metadata is also updated.
func UpdateDeviceProfile(profile models.DeviceProfile) error {
	return service.RunningService().UpdateDeviceProfile(profile)
}

// RemoveDeviceProfileByName removes the specified DeviceProfile by name from the cache and ensures that the
// instance in Core Metadata is also removed.
func RemoveDeviceProfileByName(name string) error {
	return service.RunningService().RemoveDeviceProfileByName(name)
}
