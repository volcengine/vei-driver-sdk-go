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
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestStream_SetOnDemand(t *testing.T) {
	stream := &Stream{}
	stream.SetOnDemand(true)
	require.Equal(t, true, stream.onDemand)
}

func TestStream_UpdateURL(t *testing.T) {
	stream := &Stream{
		name:          "device",
		url:           "rtsp://127.0.0.1:554/live/test",
		schema:        "rtsp",
		host:          "127.0.0.1:554",
		onDemand:      false,
		probeEnabled:  true,
		probeInterval: time.Second * 3,
		healthyCb:     NOOPHealthCheckCallback,
		mutex:         sync.Mutex{},
	}

	newUrl := "rtsp://127.0.0.1:554/live/test"
	err := stream.UpdateURL(newUrl)
	require.NoError(t, err)

	invalidUrl := "https://example.com:abc"
	err = stream.UpdateURL(invalidUrl)
	require.Error(t, err)
}
