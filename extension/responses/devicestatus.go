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

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"

	"github.com/volcengine/vei-driver-sdk-go/extension/dtos"
)

type DeviceStatusResponse struct {
	common.BaseResponse `json:",inline"`
	Status              dtos.DeviceStatus `json:"status"`
}

func NewDeviceStatusResponse(requestId string, message string, statusCode int, status dtos.DeviceStatus) DeviceStatusResponse {
	return DeviceStatusResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Status:       status,
	}
}

type MultiDeviceStatusResponse struct {
	common.BaseWithTotalCountResponse `json:",inline"`
	Status                            []dtos.DeviceStatus `json:"status"`
}

func NewMultiDeviceStatusResponse(requestId string, message string, statusCode int, totalCount uint32, status []dtos.DeviceStatus) MultiDeviceStatusResponse {
	return MultiDeviceStatusResponse{
		BaseWithTotalCountResponse: common.NewBaseWithTotalCountResponse(requestId, message, statusCode, totalCount),
		Status:                     status,
	}
}
