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
	"fmt"
	"testing"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/stretchr/testify/require"

	"github.com/volcengine/vei-driver-sdk-go/pkg/utils"
)

func MockCommandRequest(resourceName string, valueType string) models.CommandRequest {
	return models.CommandRequest{
		DeviceResourceName: resourceName,
		Attributes:         map[string]interface{}{"scale": 1},
		Type:               valueType,
	}
}

func TestNewReadRequest(t *testing.T) {
	resourceName, valueType := "temperature", Float32
	cr := MockCommandRequest(resourceName, string(valueType))

	req := NewReadRequest(cr)
	require.Equal(t, DefaultModule, req.Module())
	require.Equal(t, resourceName, req.Resource())
	require.Equal(t, valueType, req.ValueType())
}

func TestNewWriteRequest(t *testing.T) {
	resourceName, valueType := "temperature", Float32
	param := &models.CommandValue{}
	cr := MockCommandRequest(resourceName, string(valueType))

	req := NewWriteRequest(cr, param)
	require.Equal(t, DefaultModule, req.Module())
	require.Equal(t, resourceName, req.Resource())
	require.Equal(t, valueType, req.ValueType())
	require.Equal(t, param, req.Param())
}

func TestNewCallRequest(t *testing.T) {
	resourceName, valueType := "temperature", Float32
	cr := MockCommandRequest(resourceName, string(valueType))

	// invalid raw query
	cr.Attributes[utils.URLRawQuery] = ";;;"
	req, err := NewCallRequest(cr)
	require.Error(t, err)
	require.Nil(t, req)

	// normal
	cr.Attributes[utils.URLRawQuery] = utils.ServiceParams + "=eyJ4IjoxMCwieSI6MjB9"
	req, err = NewCallRequest(cr)
	require.NoError(t, err)
	require.Equal(t, `{"x":10,"y":20}`, string(req.Payload()))
}

func TestNewRequest(t *testing.T) {
	resourceName, valueType := "temperature", Float32
	cr := MockCommandRequest(resourceName, string(valueType))

	req := newRequest(cr)
	require.NotNil(t, req.Native())
	require.Equal(t, DefaultModule, req.Module())
	require.Equal(t, resourceName, req.Resource())
	require.Equal(t, valueType, req.ValueType())
	require.NotNil(t, req.Attributes())

	// set result
	req.SetResult(NewSimpleResult(1.0))
	require.NotNil(t, req.Result())

	// failed
	req.Failed(fmt.Errorf("handle failed"))
	require.Error(t, req.Error())

	// skip
	req.Skip()
	require.True(t, req.Skipped())
}
