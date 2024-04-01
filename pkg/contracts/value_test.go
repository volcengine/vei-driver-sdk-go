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
	"encoding/json"
	"testing"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/stretchr/testify/require"
)

func TestAsyncValues_Transform(t *testing.T) {
	value := &AsyncValues{}
	transformed := value.Transform()
	v1, _ := json.Marshal(value)
	v2, _ := json.Marshal(transformed)
	require.Equal(t, v1, v2)

	value = &AsyncValues{
		DeviceName:    "device",
		SourceName:    "source",
		CommandValues: []*models.CommandValue{{DeviceResourceName: "resource", Type: common.ValueTypeInt32, Value: 100}},
	}
	transformed = value.Transform()
	v1, _ = json.Marshal(value)
	v2, _ = json.Marshal(transformed)
	require.Equal(t, v1, v2)
}
