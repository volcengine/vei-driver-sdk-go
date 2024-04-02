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
	"github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
)

// AsyncValues is the struct for sending Device readings asynchronously via ProtocolDrivers
type AsyncValues struct {
	DeviceName    string
	SourceName    string
	CommandValues []*models.CommandValue
}

func (v *AsyncValues) Transform() *models.AsyncValues {
	return &models.AsyncValues{
		DeviceName:    v.DeviceName,
		SourceName:    v.SourceName,
		CommandValues: v.CommandValues,
	}
}