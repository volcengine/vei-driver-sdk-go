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
	"encoding/json"

	"github.com/volcengine/vei-driver-sdk-go/pkg/clients"
)

var media *MediaConfig

type MediaConfig struct {
	Server string             `json:"MediaServer"`
	Secret string             `json:"MediaSecret"`
	VHost  string             `json:"MediaVHost"`
	App    string             `json:"-"`
	Client *clients.ZLMClient `json:"-"`
}

func InitializeMediaConfig(configs map[string]string, app string) error {
	data, err := json.Marshal(configs)
	if err != nil {
		return err
	}

	media = &MediaConfig{}
	if err = json.Unmarshal(data, media); err != nil {
		return err
	}

	client, err := clients.NewZLMClient(media.Server)
	if err != nil {
		return err
	}

	media.Client = client
	media.App = app
	return nil
}

func Media() *MediaConfig {
	return media
}
