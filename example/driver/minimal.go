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

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"

	"github.com/volcengine/vei-driver-sdk-go/pkg/common"
	"github.com/volcengine/vei-driver-sdk-go/pkg/contracts"
	"github.com/volcengine/vei-driver-sdk-go/pkg/interfaces"
	"github.com/volcengine/vei-driver-sdk-go/pkg/logger"
)

type MinimalDriver struct {
	logger  logger.Logger
	asyncCh chan<- *contracts.AsyncValues
}

var _ interfaces.Driver = (*MinimalDriver)(nil)

func (m *MinimalDriver) Initialize(logger logger.Logger, asyncCh chan<- *contracts.AsyncValues) error {
	m.logger = logger
	m.asyncCh = asyncCh
	return nil
}

func (m *MinimalDriver) ReadProperty(device *contracts.Device, reqs []contracts.ReadRequest) error {
	for _, req := range reqs {
		m.logger.Infof("%v %v %v %v", req.Module(), req.Resource(), req.ValueType(), req.Attributes())
		switch req.ValueType() {
		case common.Int32:
			req.SetResult(contracts.NewSimpleResult(int32(rand.Intn(100))))
		case common.Float32:
			req.SetResult(contracts.NewSimpleResult(float32(rand.Intn(100))))
		case common.Bool:
			req.SetResult(contracts.NewSimpleResult(true))
		case common.String:
			req.SetResult(contracts.NewSimpleResult("this is a test message"))
		default:
			req.Failed(errors.NewCommonEdgeX(errors.KindServerError, "unsupported value type", nil))
		}
	}
	return nil
}

func (m *MinimalDriver) WriteProperty(device *contracts.Device, reqs []contracts.WriteRequest) error {
	return nil
}

func (m *MinimalDriver) CallService(device *contracts.Device, reqs []contracts.CallRequest) error {
	type Request struct {
		X float32 `json:"x"`
		Y float32 `json:"y"`
	}

	type Response struct {
		Result float32 `json:"result"`
	}

	for _, req := range reqs {
		request, response := &Request{}, &Response{}
		if err := json.Unmarshal(req.Payload(), request); err != nil {
			req.Failed(err)
		}

		switch req.Resource() {
		case "Add":
			response.Result = request.X + request.Y
		case "Sub":
			response.Result = request.X - request.Y
		case "Multiply":
			response.Result = request.X * request.Y
		case "Divide":
			if request.Y == 0 {
				req.Failed(fmt.Errorf("the divisor cannot be 0"))
			}
			response.Result = request.X / request.Y
		default:
			req.Failed(fmt.Errorf("unsupported service: %s", req.Resource()))
		}

		req.SetResult(contracts.NewSimpleResult(response))
	}

	return nil
}

func (m *MinimalDriver) Stop(force bool) error {
	return nil
}
