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

package interfaces

type Manager interface {
	// OnAddDevice is a callback function that is invoked when a new device is added
	OnAddDevice(deviceName string)
	// OnRemoveDevice is a callback function that is invoked when a device is removed
	OnRemoveDevice(deviceName string)
	// OnHandleCommandsFailed is a callback function that is invoked when failed to
	// handle read/write commands or call service.
	OnHandleCommandsFailed(deviceName string)
	// OnHandleCommandsSuccessfully is a callback function that is invoked when handling
	// read/write commands or calling service or reporting async data
	OnHandleCommandsSuccessfully(deviceName string)
	// SetDeviceOffline will set the specified device to offline status
	SetDeviceOffline(deviceName string)
	// SetDeviceOnline will set the specified device to online status
	SetDeviceOnline(deviceName string)
}
