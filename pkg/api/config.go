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

// DriverConfigs retrieves the driver specific configuration
func DriverConfigs() map[string]string {
	return service.RunningService().DriverConfigs()
}

// DeviceCommand retrieves the specific DeviceCommand instance from cache according to
// the Device name and Command name
func DeviceCommand(deviceName string, commandName string) (models.DeviceCommand, bool) {
	return service.RunningService().DeviceCommand(deviceName, commandName)
}

// DeviceResource retrieves the specific DeviceResource instance from cache according to
// the Device name and Device Resource name
func DeviceResource(deviceName string, deviceResource string) (models.DeviceResource, bool) {
	return service.RunningService().DeviceResource(deviceName, deviceResource)
}
