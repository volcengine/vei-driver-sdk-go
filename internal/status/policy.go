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
	"fmt"
	"time"
)

type Policy string

const (
	ExceedConsecutiveErrorNum     Policy = "ExceedConsecutiveErrorNum"
	ExceedContinuousErrorDuration Policy = "ExceedContinuousErrorDuration"
)

type OfflineDecision struct {
	policy    Policy
	threshold int64
}

func NewOfflineDecision(policy Policy, threshold int64) OfflineDecision {
	return OfflineDecision{policy: policy, threshold: threshold}
}

func ValidateOfflineDecision(decision OfflineDecision) error {
	switch decision.policy {
	case ExceedConsecutiveErrorNum:
		if decision.threshold <= 0 {
			return fmt.Errorf("the specified consecutive error number cannot be zero or negative")
		}
	case ExceedContinuousErrorDuration:
		if decision.threshold < int64(time.Minute.Seconds()) {
			return fmt.Errorf("the specified continuous error duration cannot be less than 60 seconds")
		}
	default:
		return fmt.Errorf("unsupported offline decision policy")
	}
	return nil
}
