package zlmediakit

import (
	"context"
	"fmt"
	"net/url"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
)

type MediaServerClientInterface interface {
	GetSnap(streamURI string) ([]byte, error)
}

const (
	APIBase         = "/index/api"
	APIGetSnapRoute = APIBase + "/getSnap"
)

type Client struct {
	baseURL string
	secret  string
	lc      logger.LoggingClient
	ctx     context.Context
}

// NewClient creates an instance of zlmdiakit client
func NewClient(baseURL string, secret string, lc logger.LoggingClient) (MediaServerClientInterface, error) {
	parsedURL, err := url.ParseRequestURI(baseURL)
	if err != nil {
		lc.Errorf("parse url %s failed: %v", baseURL, err)
		return nil, err
	}
	return &Client{baseURL: parsedURL.String(), secret: secret, ctx: context.Background(), lc: lc}, nil
}

func (zc Client) GetSnap(streamURI string) ([]byte, error) {
	requestParams := url.Values{}
	requestParams.Set("url", streamURI)
	requestParams.Set("timeout_sec", "10")
	requestParams.Set("expire_sec", "1")
	if len(zc.secret) != 0 {
		requestParams.Set("secret", zc.secret)
	}

	response, contentType, err := utils.GetRequestAndReturnBinaryRes(zc.ctx, zc.baseURL, APIGetSnapRoute, requestParams)
	if err != nil {
		zc.lc.Errorf("call %s api error: %v, resp: %v", APIGetSnapRoute, err, response)
		return response, errors.NewCommonEdgeXWrapper(err)
	}

	if contentType != "image/jpeg" {
		zc.lc.Errorf("call %s api error: %v, resp: %v, contentType: %v", APIGetSnapRoute, err, response, contentType)
		return response, errors.NewCommonEdgeXWrapper(fmt.Errorf("contentType error: %v", contentType))
	}

	return response, nil
}
