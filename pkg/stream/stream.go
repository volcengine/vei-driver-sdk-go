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
	"fmt"
	"net"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/volcengine/vei-driver-sdk-go/gen/zlm"
	"github.com/volcengine/vei-driver-sdk-go/pkg/logger"
)

const (
	DefaultHealthCheckInterval  = time.Minute
	DefaultHealthCheckThreshold = time.Second * 30
)

type HealthCheckCallback func(name string, err error)

func NOOPHealthCheckCallback(name string, err error) {}

type Stream struct {
	name   string
	url    string
	schema string
	host   string

	onDemand      bool          // 是否按需拉流，默认开启
	probeEnabled  bool          // 是否开启健康检查，默认开启
	probeInterval time.Duration // 健康检查间隔，默认1分钟
	healthyCb     HealthCheckCallback
	healthy       bool

	mutex    sync.Mutex
	shutdown context.CancelFunc
}

func NewStream(deviceName string, streamUrl string, opts ...func(*Stream)) (*Stream, error) {
	u, err := url.Parse(streamUrl)
	if err != nil {
		return nil, err
	}

	stream := &Stream{
		name:          deviceName,
		url:           streamUrl,
		schema:        u.Scheme,
		host:          u.Host,
		onDemand:      true,
		probeEnabled:  true,
		probeInterval: DefaultHealthCheckInterval,
		healthyCb:     NOOPHealthCheckCallback,
		healthy:       false,
		mutex:         sync.Mutex{},
	}

	for _, opt := range opts {
		opt(stream)
	}

	return stream, nil
}

func WithOnDemand(onDemand bool) func(*Stream) {
	return func(stream *Stream) {
		stream.onDemand = onDemand
	}
}

func WithProbeEnabled(probeEnabled bool) func(*Stream) {
	return func(stream *Stream) {
		stream.probeEnabled = probeEnabled
	}
}

func WithProbeInterval(probeInterval time.Duration) func(*Stream) {
	return func(stream *Stream) {
		if probeInterval > DefaultHealthCheckThreshold {
			stream.probeInterval = probeInterval
		}
	}
}

func WithHealthCheckCallback(cb HealthCheckCallback) func(*Stream) {
	return func(stream *Stream) {
		if cb != nil {
			stream.healthyCb = cb
		}
	}
}

func (s *Stream) Start() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	ctx, cancel := context.WithCancel(context.Background())
	s.shutdown = cancel

	if s.probeEnabled {
		go s.healthCheck(ctx)
	}

	if err := s.start(ctx); err != nil {
		return err
	}

	s.healthy = true
	return nil
}

func (s *Stream) start(ctx context.Context) error {
	logger.D.Infof("add stream proxy for device %s, url=%s", s.name, s.url)
	resp, err := media.Client.AddStreamProxy(ctx, &zlm.AddStreamProxyParams{
		Secret: &media.Secret,
		Vhost:  &media.VHost,
		App:    &media.App,
		Stream: &s.name,
		Url:    &s.url,
	})
	if err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf("invoke AddStreamProxy failed: %s", resp.Message)
	}
	return nil
}

func (s *Stream) Stop() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.shutdown != nil {
		s.shutdown()
	}

	return s.stop(context.Background())
}

func (s *Stream) stop(ctx context.Context) error {
	logger.D.Infof("delete stream proxy for device %s", s.name)
	key := strings.Join([]string{media.VHost, media.App, s.name}, "/")
	resp, err := media.Client.DelStreamProxy(ctx, &zlm.DelStreamProxyParams{
		Secret: &media.Secret,
		Key:    &key,
	})
	if err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf("invoke DelStreamProxy failed: %s", resp.Message)
	}
	return nil
}

func (s *Stream) healthCheck(ctx context.Context) {
	ticker := time.NewTicker(s.probeInterval)
	defer ticker.Stop()
	logger.D.Infof("starting health check for stream %s", s.name)

	for {
		select {
		case <-ticker.C:
			// 若按需拉流，通过TCP探测判断流是否健康；否则检查流是否在线，若不在线需要重新添加拉流代理
			if s.onDemand {
				err := s.tcpProbe(ctx)
				s.healthy = err == nil
				s.healthyCb(s.name, err)
			} else {
				err := s.checkMediaOnline(ctx)
				s.healthy = err == nil
				s.healthyCb(s.name, err)
			}
		case <-ctx.Done():
			logger.D.Infof("stopping health check for stream %s", s.name)
			return
		}
	}
}

func (s *Stream) tcpProbe(ctx context.Context) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	conn, err := net.DialTimeout("tcp", s.host, time.Second*2)
	if err != nil {
		logger.D.Errorf("connect to %s failed when using simple tcp dial: %v", s.url, err)
		return err
	}
	_ = conn.Close()

	return nil
}

func (s *Stream) checkMediaOnline(ctx context.Context) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	online, err := s.IsMediaOnline(ctx)
	if err != nil {
		logger.D.Errorf("failed to check if stream %s is online: %v", s.name, err)
		return err
	}

	if !online {
		if err = s.start(ctx); err != nil {
			logger.D.Errorf("failed to restart stream %s: %v", s.name, err)
			return err
		}
	}

	return nil
}

func (s *Stream) IsMediaOnline(ctx context.Context) (bool, error) {
	resp, err := media.Client.IsMediaOnline(ctx, &zlm.IsMediaOnlineParams{
		Secret: &media.Secret,
		Vhost:  &media.VHost,
		App:    &media.App,
		Schema: &s.schema,
		Stream: &s.name,
	})
	if err != nil {
		return false, err
	}
	if resp.Code != 0 {
		return false, fmt.Errorf("invoke IsMediaOnline failed: %s", resp.Message)
	}
	return resp.Online, nil
}

func (s *Stream) Healthy() bool {
	return s.healthy
}
