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
	"encoding/json"
	"net/url"

	"github.com/volcengine/vei-driver-sdk-go/pkg/clients"
)

var config *Config

type Config struct {
	Server   string             `json:"MediaServer"`
	Secret   string             `json:"MediaSecret"`
	VHost    string             `json:"MediaVHost"`
	HostName string             `json:"HostName"`
	App      string             `json:"-"`
	Client   *clients.ZLMClient `json:"-"`
}

func InitializeConfig(configs map[string]string, app string) error {
	data, err := json.Marshal(configs)
	if err != nil {
		return err
	}

	config = &Config{}
	if err = json.Unmarshal(data, config); err != nil {
		return err
	}

	_url, err := url.Parse(config.Server)
	if err != nil {
		return err
	}
	if config.HostName == "" {
		config.HostName = _url.Hostname()
	}

	client, err := clients.NewZLMClient(config.Server)
	if err != nil {
		return err
	}

	config.Client = client
	config.App = app
	return nil
}

func Media() *Config {
	return config
}

func Server() string {
	return config.Server
}

func Secret() string {
	return config.Secret
}

func VHost() string {
	return config.VHost
}

func HostName() string {
	return config.HostName
}

func App() string {
	return config.App
}

func Client() *clients.ZLMClient {
	return config.Client
}
