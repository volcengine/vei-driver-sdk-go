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

package hook

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"

	"github.com/volcengine/vei-driver-sdk-go/pkg/interfaces"
	"github.com/volcengine/vei-driver-sdk-go/pkg/logger"
)

type StreamRequest struct {
	App    string `json:"app"`
	Schema string `json:"schema"`
	Stream string `json:"stream"`
	Vhost  string `json:"vhost"`
}

func OnStreamNotFound(webhook interfaces.Webhook) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		if webhook == nil {
			edgexErr := errors.NewCommonEdgeX(errors.KindNotImplemented, "Please implement the webhook interface", nil)
			WriteErrorResponse(writer, edgexErr)
			return
		}

		defer request.Body.Close()
		body, err := io.ReadAll(request.Body)
		if err != nil {
			edgexErr := errors.NewCommonEdgeX(errors.KindServerError, "failed to read request body", err)
			WriteErrorResponse(writer, edgexErr)
			return
		}

		param := &StreamRequest{}
		if err = json.Unmarshal(body, param); err != nil {
			edgexErr := errors.NewCommonEdgeX(errors.KindServerError, "failed to parse request body", err)
			WriteErrorResponse(writer, edgexErr)
			return
		}

		ctx := request.Context()
		if err = webhook.OnStreamNotFound(ctx, param.Schema, param.Stream); err != nil {
			edgexErr := errors.NewCommonEdgeX(errors.KindServerError, "failed to execute webhook", err)
			WriteErrorResponse(writer, edgexErr)
			return
		}

		writer.WriteHeader(http.StatusNoContent)
	}
}

func OnStreamNoneReader(webhook interfaces.Webhook) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		if webhook == nil {
			edgexErr := errors.NewCommonEdgeX(errors.KindNotImplemented, "Please implement the webhook interface", nil)
			WriteErrorResponse(writer, edgexErr)
			return
		}

		defer request.Body.Close()
		body, err := io.ReadAll(request.Body)
		if err != nil {
			edgexErr := errors.NewCommonEdgeX(errors.KindServerError, "failed to read request body", err)
			WriteErrorResponse(writer, edgexErr)
			return
		}

		param := &StreamRequest{}
		if err = json.Unmarshal(body, param); err != nil {
			edgexErr := errors.NewCommonEdgeX(errors.KindServerError, "failed to parse request body", err)
			WriteErrorResponse(writer, edgexErr)
			return
		}

		ctx := request.Context()
		if err = webhook.OnStreamNoneReader(ctx, param.Schema, param.Stream); err != nil {
			edgexErr := errors.NewCommonEdgeX(errors.KindServerError, "failed to execute webhook", err)
			WriteErrorResponse(writer, edgexErr)
			return
		}

		writer.WriteHeader(http.StatusNoContent)
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
