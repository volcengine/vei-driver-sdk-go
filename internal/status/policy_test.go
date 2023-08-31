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

package status

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewOfflineDecision(t *testing.T) {
	decision := NewOfflineDecision(ExceedConsecutiveErrorNum, 10)
	require.NotEmpty(t, decision)
}

func TestValidateOfflineDecision(t *testing.T) {
	type args struct {
		decision OfflineDecision
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "unsupported policy", args: args{decision: OfflineDecision{}}, wantErr: true},
		{name: "invalid error num", args: args{decision: OfflineDecision{policy: ExceedConsecutiveErrorNum, threshold: 0}}, wantErr: true},
		{name: "valid error num", args: args{decision: OfflineDecision{policy: ExceedConsecutiveErrorNum, threshold: 1}}, wantErr: false},
		{name: "invalid error duration", args: args{decision: OfflineDecision{policy: ExceedContinuousErrorDuration, threshold: 0}}, wantErr: true},
		{name: "valid error duration", args: args{decision: OfflineDecision{policy: ExceedContinuousErrorDuration, threshold: int64(time.Minute.Seconds())}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateOfflineDecision(tt.args.decision); (err != nil) != tt.wantErr {
				t.Errorf("ParseOfflineDecision() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
