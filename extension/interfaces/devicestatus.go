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

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"

	"github.com/volcengine/vei-driver-sdk-go/extension/requests"
	"github.com/volcengine/vei-driver-sdk-go/extension/responses"
)

// DeviceStatusClient defines the interface for interactions with the DeviceStatus on the EdgeX Foundry core-metadata service.
type DeviceStatusClient interface {
	// Update updates the device status.
	Update(ctx context.Context, req requests.UpdateDeviceStatusRequest) (common.BaseResponse, errors.EdgeX)
	// AllDeviceStatus returns all device status.
	// The result can be limited in a certain range by specifying the offset and limit parameters.
	// offset: The number of items to skip before starting to collect the result set. Default is 0.
	// limit: The number of items to return. Specify -1 will return all remaining items after offset.
	// The maximum will be the MaxResultCount as defined in the configuration of service. Default is 20.
	AllDeviceStatus(ctx context.Context, offset int, limit int) (responses.MultiDeviceStatusResponse, errors.EdgeX)
	// DeviceStatusByName returns a device status by device name.
	DeviceStatusByName(ctx context.Context, name string) (responses.DeviceStatusResponse, errors.EdgeX)
}
