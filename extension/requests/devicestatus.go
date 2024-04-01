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

package requests

import (
	"encoding/json"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"

	"github.com/volcengine/vei-driver-sdk-go/extension/dtos"
)

type UpdateDeviceStatusRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	Status                dtos.UpdateDeviceStatus `json:"status"`
}

func (d *UpdateDeviceStatusRequest) Validate() error {
	err := common.Validate(d)
	return err
}

func (d *UpdateDeviceStatusRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		Status dtos.UpdateDeviceStatus
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*d = UpdateDeviceStatusRequest(alias)

	// validate UpdateDeviceStatusRequest DTO
	if err := d.Validate(); err != nil {
		return err
	}
	return nil
}

func NewUpdateDeviceStatusRequest(dto dtos.UpdateDeviceStatus) UpdateDeviceStatusRequest {
	return UpdateDeviceStatusRequest{
		BaseRequest: dtoCommon.NewBaseRequest(),
		Status:      dto,
	}
}
