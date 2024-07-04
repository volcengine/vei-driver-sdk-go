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

package media

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	server = "http://test-media-server"
	secret = "035c73f7-bb6b-4889-a715-d9eb2d1925cc"
	vhost  = "__defaultVhost__"
)

func MockMediaConfig() {
	if config == nil {
		conf := map[string]string{
			"MediaServer": server,
			"MediaSecret": secret,
			"MediaVhost":  vhost,
		}
		_ = InitializeConfig(conf, "mock")
	}
}

func TestInitializeMediaConfig(t *testing.T) {
	defer func() {
		config = nil
	}()

	type args struct {
		configs map[string]string
		app     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "nil config", args: args{configs: nil, app: "mock"}, wantErr: false},
		{name: "empty config", args: args{configs: map[string]string{}, app: "mock"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitializeConfig(tt.args.configs, tt.args.app); (err != nil) != tt.wantErr {
				t.Errorf("InitializeConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMedia(t *testing.T) {
	MockMediaConfig()
	require.Equal(t, server, Media().Server)
	require.Equal(t, secret, Media().Secret)
	require.Equal(t, vhost, Media().VHost)
}
