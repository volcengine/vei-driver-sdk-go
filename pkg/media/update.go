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
	"net/url"
)

func (s *Stream) SetOnDemand(onDemand bool) {
	s.onDemand = onDemand
}

func (s *Stream) UpdateURL(newUrl string) error {
	u, err := url.Parse(newUrl)
	if err != nil {
		return err
	}

	oldUrl := s.url
	s.url = newUrl
	s.schema = u.Scheme
	s.host = u.Host

	if oldUrl != newUrl {
		if err = s.Stop(); err != nil {
			return err
		}
		if err = s.Start(); err != nil {
			return err
		}
	}

	return nil
}
