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

package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/volcengine/vei-driver-sdk-go/gen/zlm"
)

type ZLMClient struct {
	Native *zlm.ClientWithResponses
}

func NewZLMClient(server string, opts ...zlm.ClientOption) (*ZLMClient, error) {
	client, err := zlm.NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	native := &zlm.ClientWithResponses{ClientInterface: client}
	return &ZLMClient{Native: native}, nil
}

type CommonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type AddFFmpegSourceResponse struct {
	CommonResponse
	Data struct {
		Key string `json:"key"`
	} `json:"data"`
}

func (c *ZLMClient) AddFFmpegSource(ctx context.Context, params *zlm.AddFFmpegSourceParams) (*AddFFmpegSourceResponse, error) {
	rsp, err := c.Native.AddFFmpegSourceWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("AddFFmpegSource failed: %s", string(rsp.Body))
	}
	var response AddFFmpegSourceResponse
	err = json.Unmarshal(rsp.Body, &response)
	return &response, err
}

type AddStreamProxyResponse struct {
	CommonResponse
}

func (c *ZLMClient) AddStreamProxy(ctx context.Context, params *zlm.AddStreamProxyParams) (*AddStreamProxyResponse, error) {
	rsp, err := c.Native.AddStreamProxyWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("AddStreamProxy failed: %s", string(rsp.Body))
	}
	var response AddStreamProxyResponse
	err = json.Unmarshal(rsp.Body, &response)
	return &response, err
}

type AddStreamPusherProxyResponse struct {
	CommonResponse
	Data struct {
		Key string `json:"key"`
	} `json:"data"`
}

func (c *ZLMClient) AddStreamPusherProxy(ctx context.Context, params *zlm.AddStreamPusherProxyParams) (*AddStreamPusherProxyResponse, error) {
	rsp, err := c.Native.AddStreamPusherProxyWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("AddStreamPusherProxy failed: %s", string(rsp.Body))
	}
	var response AddStreamPusherProxyResponse
	err = json.Unmarshal(rsp.Body, &response)
	return &response, err
}

type CloseStreamResponse struct {
	CommonResponse
	Result int `json:"result"`
}

func (c *ZLMClient) CloseStream(ctx context.Context, params *zlm.CloseStreamParams) (*CloseStreamResponse, error) {
	rsp, err := c.Native.CloseStreamWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("CloseStream failed: %s", string(rsp.Body))
	}
	var response CloseStreamResponse
	err = json.Unmarshal(rsp.Body, &response)
	return &response, err
}

type CloseStreamsResponse struct {
	CommonResponse
	CountHit    int `json:"count_hit"`
	CountClosed int `json:"count_closed"`
}

func (c *ZLMClient) CloseStreams(ctx context.Context, params *zlm.CloseStreamsParams) (*CloseStreamsResponse, error) {
	rsp, err := c.Native.CloseStreamsWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("CloseStreams failed: %s", string(rsp.Body))
	}
	var response CloseStreamsResponse
	err = json.Unmarshal(rsp.Body, &response)
	return &response, err
}

type DelFFmpegSourceResponse struct {
	CommonResponse
	Data struct {
		Flag bool `json:"flag"`
	} `json:"data"`
}

func (c *ZLMClient) DelFFmpegSource(ctx context.Context, params *zlm.DelFFmpegSourceParams) (*DelFFmpegSourceResponse, error) {
	rsp, err := c.Native.DelFFmpegSourceWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("DelFFmpegSource failed: %s", string(rsp.Body))
	}
	var response DelFFmpegSourceResponse
	err = json.Unmarshal(rsp.Body, &response)
	return &response, err
}

type DelStreamProxyResponse struct {
	CommonResponse
	Data struct {
		Flag bool `json:"flag"`
	} `json:"data"`
}

func (c *ZLMClient) DelStreamProxy(ctx context.Context, params *zlm.DelStreamProxyParams) (*DelStreamProxyResponse, error) {
	rsp, err := c.Native.DelStreamProxyWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("DelStreamProxy failed: %s", string(rsp.Body))
	}
	var response DelStreamProxyResponse
	err = json.Unmarshal(rsp.Body, &response)
	return &response, err
}

type DelStreamPusherProxyResponse struct {
	CommonResponse
	Data struct {
		Flag bool `json:"flag"`
	} `json:"data"`
}

func (c *ZLMClient) DelStreamPusherProxy(ctx context.Context, params *zlm.DelStreamPusherProxyParams) (*DelStreamPusherProxyResponse, error) {
	rsp, err := c.Native.DelStreamPusherProxyWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("DelStreamPusherProxy failed: %s", string(rsp.Body))
	}
	var response DelStreamPusherProxyResponse
	err = json.Unmarshal(rsp.Body, &response)
	return &response, err
}

type GetMediaPlayerListResponse struct {
	CommonResponse
	Data []struct {
		Identifier string `json:"identifier"`
		LocalIp    string `json:"local_ip"`
		LocalPort  int    `json:"local_port"`
		PeerIp     string `json:"peer_ip"`
		PeerPort   int    `json:"peer_port"`
		Typeid     string `json:"typeid"`
	} `json:"data"`
}

func (c *ZLMClient) GetMediaPlayerList(ctx context.Context, params *zlm.GetMediaPlayerListParams) (*GetMediaPlayerListResponse, error) {
	rsp, err := c.Native.GetMediaPlayerListWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("GetMediaPlayerList failed: %s", string(rsp.Body))
	}
	var response GetMediaPlayerListResponse
	err = json.Unmarshal(rsp.Body, &response)
	return &response, err
}

type IsMediaOnlineResponse struct {
	CommonResponse
	Online bool `json:"online"`
}

func (c *ZLMClient) IsMediaOnline(ctx context.Context, params *zlm.IsMediaOnlineParams) (*IsMediaOnlineResponse, error) {
	rsp, err := c.Native.IsMediaOnlineWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("IsMediaOnline failed: %s", string(rsp.Body))
	}
	var response IsMediaOnlineResponse
	err = json.Unmarshal(rsp.Body, &response)
	return &response, err
}
