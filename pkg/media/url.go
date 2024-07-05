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
	"fmt"
)

type StreamURL interface {
	RTSP(stream string) string // rtsp
	RTMP(stream string) string // rtmp
	HLS(stream string) string  // hls-mpegts
	FLV(stream string) string  // http-flv
	TS(stream string) string   // http-ts
	FMP4(stream string) string // http-fmp4
}

type streamURL struct{}

func URL() StreamURL {
	return &streamURL{}
}

func (s *streamURL) RTSP(stream string) string {
	return fmt.Sprintf("rtsp://%s:554/%s/%s", Media().HostName, Media().App, stream)
}

func (s *streamURL) RTMP(stream string) string {
	return fmt.Sprintf("rtmp://%s:1935/%s/%s", Media().HostName, Media().App, stream)
}

func (s *streamURL) HLS(stream string) string {
	return fmt.Sprintf("http://%s/%s/%s/hls.m3u8", Media().HostName, Media().App, stream)
}

func (s *streamURL) FLV(stream string) string {
	return fmt.Sprintf("http://%s/%s/%s.live.flv", Media().HostName, Media().App, stream)
}

func (s *streamURL) TS(stream string) string {
	return fmt.Sprintf("http://%s/%s/%s.live.ts", Media().HostName, Media().App, stream)
}

func (s *streamURL) FMP4(stream string) string {
	return fmt.Sprintf("http://%s/%s/%s.live.mp4", Media().HostName, Media().App, stream)
}
