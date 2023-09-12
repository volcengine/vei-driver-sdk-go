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

package driver

import (
	"encoding/json"
	"fmt"
	"math/rand"

	sdkmodels "github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/volcengine/vei-driver-sdk-go/pkg/interfaces"
)

type MinimalDriver struct {
	lc       logger.LoggingClient
	asyncCh  chan<- *sdkmodels.AsyncValues
	reporter interfaces.EventReporter
}

var _ interfaces.Driver = (*MinimalDriver)(nil)

func (m *MinimalDriver) Initialize(lc logger.LoggingClient, asyncCh chan<- *sdkmodels.AsyncValues, eventReporter interfaces.EventReporter) error {
	m.lc = lc
	m.asyncCh = asyncCh
	m.reporter = eventReporter
	return nil
}

func (m *MinimalDriver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []sdkmodels.CommandRequest) ([]*sdkmodels.CommandValue, error) {
	result := make([]*sdkmodels.CommandValue, len(reqs))
	for i, req := range reqs {
		var cv *sdkmodels.CommandValue
		var err error
		switch req.Type {
		case common.ValueTypeInt32:
			cv, err = sdkmodels.NewCommandValue(req.DeviceResourceName, common.ValueTypeInt32, rand.Int())
		case common.ValueTypeFloat32:
			cv, err = sdkmodels.NewCommandValue(req.DeviceResourceName, common.ValueTypeFloat32, rand.Int())
		case common.ValueTypeBool:
			cv, err = sdkmodels.NewCommandValue(req.DeviceResourceName, common.ValueTypeBool, true)
		case common.ValueTypeString:
			cv, err = sdkmodels.NewCommandValue(req.DeviceResourceName, common.ValueTypeString, "this is a test message")
		default:
			return nil, errors.NewCommonEdgeX(errors.KindServerError, "unsupported value type", nil)
		}
		if err != nil {
			return result, errors.NewCommonEdgeX(errors.KindServerError, fmt.Sprintf("failed to get %s value ", req.DeviceResourceName), err)
		} else {
			result[i] = cv
		}
	}

	return result, nil
}

func (m *MinimalDriver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []sdkmodels.CommandRequest, params []*sdkmodels.CommandValue) error {
	return nil
}

func (m *MinimalDriver) HandleServiceCall(deviceName string, protocols map[string]models.ProtocolProperties, req sdkmodels.CommandRequest, data []byte) (*sdkmodels.CommandValue, error) {
	type Request struct {
		X float32 `json:"x"`
		Y float32 `json:"y"`
	}

	type Response struct {
		Result float32 `json:"result"`
	}

	request, response := &Request{}, &Response{}
	if err := json.Unmarshal(data, request); err != nil {
		return nil, err
	}

	switch req.DeviceResourceName {
	case "Add":
		response.Result = request.X + request.Y
	case "Sub":
		response.Result = request.X - request.Y
	case "Multiply":
		response.Result = request.X * request.Y
	case "Divide":
		if request.Y == 0 {
			return nil, fmt.Errorf("the divisor cannot be 0")
		}
		response.Result = request.X / request.Y
	default:
		return nil, fmt.Errorf("unsupported service: %s", req.DeviceResourceName)
	}

	return sdkmodels.NewCommandValue(req.DeviceResourceName, common.ValueTypeObject, response)
}

func (m *MinimalDriver) Stop(force bool) error {
	return nil
}
