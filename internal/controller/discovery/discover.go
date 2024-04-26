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

package discovery

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"

	"github.com/volcengine/vei-driver-sdk-go/pkg/contracts"
	"github.com/volcengine/vei-driver-sdk-go/pkg/interfaces"
	"github.com/volcengine/vei-driver-sdk-go/pkg/logger"
	"github.com/volcengine/vei-driver-sdk-go/pkg/utils"
)

const (
	DefaultMaxDurationSeconds = 30
	DefaultDeviceChannelNum   = 10
)

var mutex sync.Mutex

func Discover(discovery interfaces.Discovery) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		if discovery == nil {
			edgexErr := errors.NewCommonEdgeX(errors.KindNotImplemented, "Please implement the discovery interface", nil)
			WriteErrorResponse(writer, edgexErr)
			return
		}

		if !mutex.TryLock() {
			edgexErr := errors.NewCommonEdgeX(errors.KindContractInvalid, "Please wait for the last operation complete", nil)
			WriteErrorResponse(writer, edgexErr)
			return
		}

		defer mutex.Unlock()
		defer request.Body.Close()

		body, err := io.ReadAll(request.Body)
		if err != nil {
			edgexErr := errors.NewCommonEdgeX(errors.KindServerError, "failed to read request body", err)
			WriteErrorResponse(writer, edgexErr)
			return
		}

		param := &contracts.DiscoveryParameter{}
		if err = json.Unmarshal(body, param); err != nil {
			edgexErr := errors.NewCommonEdgeX(errors.KindServerError, "failed to parse request body", err)
			WriteErrorResponse(writer, edgexErr)
			return
		}

		param.MaxDurationTime = utils.Ternary(param.MaxDurationTime == 0, time.Second*DefaultMaxDurationSeconds, param.MaxDurationTime)
		ctx, cancel := context.WithTimeout(request.Context(), param.MaxDurationTime)
		defer cancel()

		deviceChan := make(chan *contracts.Device, DefaultDeviceChannelNum)
		go discovery.Discover(ctx, param, deviceChan)

		for {
			select {
			case <-ctx.Done():
				return
			case device, ok := <-deviceChan:
				if !ok {
					return
				}
				if device == nil {
					continue
				}
				bytes, err := json.Marshal(device)
				if err != nil {
					logger.D.Errorf("marshal device in json format failed: %v", err)
					continue
				}
				if _, err = writer.Write(append(bytes, '\n')); err != nil {
					logger.D.Errorf("write response failed: %v", err)
					continue
				}
				writer.(http.Flusher).Flush()
			}
		}
	}
}

func WriteErrorResponse(w http.ResponseWriter, edgexErr errors.EdgeX) {
	logger.D.Errorf("%v", edgexErr.Error())
	responses := common.NewBaseResponse("", edgexErr.Error(), edgexErr.Code())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(edgexErr.Code())

	enc := json.NewEncoder(w)
	err := enc.Encode(responses)
	if err != nil {
		logger.D.Errorf("Error encoding the data: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
