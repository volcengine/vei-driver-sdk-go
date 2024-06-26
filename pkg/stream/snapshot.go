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

package stream

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"

	"github.com/volcengine/vei-driver-sdk-go/gen/zlm"
	"github.com/volcengine/vei-driver-sdk-go/pkg/utils"
)

const (
	SnapshotResultHeader = "data:image/jpeg;base64,"
)

type SnapshotResponse struct {
	Result string `json:"result"`
}

func NewSnapshotResponse(snapshot []byte) *SnapshotResponse {
	encoded := base64.StdEncoding.EncodeToString(snapshot)
	return &SnapshotResponse{Result: SnapshotResultHeader + encoded}
}

func (s *Stream) Snapshot(ctx context.Context) (*SnapshotResponse, error) {
	expire := 1
	timeout := 10
	localUrl := fmt.Sprintf("rtsp://127.0.0.1:554/%s/%s", media.App, s.name)
	stream := utils.Ternary(s.onDemand, s.url, localUrl)

	resp, err := media.Client.Native.GetSnapWithResponse(ctx, &zlm.GetSnapParams{
		Secret:     &media.Secret,
		Url:        &stream,
		TimeoutSec: &timeout,
		ExpireSec:  &expire,
	})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(resp.Body))
	}
	return NewSnapshotResponse(resp.Body), nil
}
