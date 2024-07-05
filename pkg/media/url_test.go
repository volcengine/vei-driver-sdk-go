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

func TestStreamURL(t *testing.T) {
	MockMediaConfig()
	url := URL()
	stream := "test"

	u := url.RTSP(stream)
	require.Equal(t, "rtsp://test-media-server.default:554/mock/test", u)

	u = url.RTMP(stream)
	require.Equal(t, "rtmp://test-media-server.default:1935/mock/test", u)

	u = url.HLS(stream)
	require.Equal(t, "http://test-media-server.default/mock/test/hls.m3u8", u)

	u = url.FLV(stream)
	require.Equal(t, "http://test-media-server.default/mock/test.live.flv", u)

	u = url.TS(stream)
	require.Equal(t, "http://test-media-server.default/mock/test.live.ts", u)

	u = url.FMP4(stream)
	require.Equal(t, "http://test-media-server.default/mock/test.live.mp4", u)
}
