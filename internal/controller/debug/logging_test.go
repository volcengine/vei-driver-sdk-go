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

package debug

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/stretchr/testify/require"
)

func TestSetDefaultLogLevel(t *testing.T) {
	tests := []struct {
		name           string
		level          string
		wantStatusCode int
	}{
		{name: "parse failed", level: "err", wantStatusCode: http.StatusBadRequest},
		{name: "debug", level: "debug", wantStatusCode: http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, common.ApiBase+"/logging", http.NoBody)
			query := req.URL.Query()
			query.Add(LogLevel, tt.level)
			req.URL.RawQuery = query.Encode()
			require.NoError(t, err)

			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(SetDefaultLogLevel)
			handler.ServeHTTP(recorder, req)
			require.Equal(t, tt.wantStatusCode, recorder.Result().StatusCode, "HTTP status code not as expected")
		})
	}
}
