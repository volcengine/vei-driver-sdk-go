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

func TestNewStream(t *testing.T) {
	type args struct {
		deviceName string
		streamUrl  string
		opts       []func(*Stream)
	}
	tests := []struct {
		name    string
		args    args
		want    *Stream
		wantErr bool
	}{
		{
			name:    "invalid url",
			args:    args{deviceName: "device", streamUrl: "https://example.com:abc", opts: []func(*Stream){}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "default",
			args: args{deviceName: "device", streamUrl: "rtsp://127.0.0.1:554/live/test", opts: []func(*Stream){}},
			want: &Stream{
				name:          "device",
				url:           "rtsp://127.0.0.1:554/live/test",
				schema:        "rtsp",
				host:          "127.0.0.1:554",
				onDemand:      true,
				probeEnabled:  true,
				probeInterval: DefaultHealthCheckInterval,
				healthyCb:     NOOPHealthCheckCallback,
				mutex:         sync.Mutex{},
			},
			wantErr: false,
		},
		{
			name: "with options",
			args: args{deviceName: "device", streamUrl: "rtsp://127.0.0.1:554/live/test", opts: []func(*Stream){
				WithOnDemand(false),
				WithProbeEnabled(false),
				WithProbeInterval(time.Minute * 2),
				WithHealthCheckCallback(NOOPHealthCheckCallback),
			}},
			want: &Stream{
				name:          "device",
				url:           "rtsp://127.0.0.1:554/live/test",
				schema:        "rtsp",
				host:          "127.0.0.1:554",
				onDemand:      false,
				probeEnabled:  false,
				probeInterval: DefaultHealthCheckInterval * 2,
				healthyCb:     NOOPHealthCheckCallback,
				mutex:         sync.Mutex{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStream(tt.args.deviceName, tt.args.streamUrl, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && got == nil {
				require.Equalf(t, tt.want.name, got.name, "NewStream() name is %v, want %v", got.name, tt.want.name)
				require.Equalf(t, tt.want.url, got.url, "NewStream() url is %v, want %v", got.url, tt.want.url)
				require.Equalf(t, tt.want.schema, got.schema, "NewStream() schema is %v, want %v", got.schema, tt.want.schema)
				require.Equalf(t, tt.want.host, got.host, "NewStream() host is %v, want %v", got.host, tt.want.host)
				require.Equalf(t, tt.want.onDemand, got.onDemand, "NewStream() onDemand is %v, want %v", got.onDemand, tt.want.onDemand)
				require.Equalf(t, tt.want.probeEnabled, got.probeEnabled, "NewStream() probe is %v, want %v", got.probeEnabled, tt.want.probeEnabled)
				require.Equalf(t, tt.want.probeInterval, got.probeInterval, "NewStream() probe interval is %v, want %v", got.probeInterval, tt.want.probeInterval)
			}
		})
	}
}

func TestStreamLifecycle(t *testing.T) {
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

	err := stream.Start()
	require.Error(t, err)
	require.False(t, stream.Healthy())

	time.Sleep(time.Second * 7)
	err = stream.Stop()
	require.Error(t, err)
}
