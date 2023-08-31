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

package resource

import (
	sdkmodels "github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
)

type Category int

const (
	Property Category = 1
	Service  Category = 2
	Event    Category = 3
)

func (c Category) String() string {
	switch c {
	case Property:
		return "property"
	case Service:
		return "service"
	case Event:
		return "event"
	default:
		return "other"
	}
}

const (
	CategoryKey = "category"
)

func GetCategory(req sdkmodels.CommandRequest) Category {
	if req.Attributes == nil || req.Attributes[CategoryKey] == nil {
		return Property
	}
	c := req.Attributes[CategoryKey]
	c, ok := c.(string)
	if !ok {
		return Property
	}
	switch c {
	case Property.String():
		return Property
	case Service.String():
		return Service
	case Event.String():
		return Event
	default:
		return Property
	}
}

func SetCategory(attributes map[string]interface{}, category Category) map[string]interface{} {
	if attributes == nil {
		attributes = make(map[string]interface{}, 0)
	}
	attributes[CategoryKey] = category.String()
	return attributes
}

func IsProperty(req sdkmodels.CommandRequest) bool {
	return GetCategory(req) == Property
}

func IsService(req sdkmodels.CommandRequest) bool {
	return GetCategory(req) == Service
}

func IsEvent(req sdkmodels.CommandRequest) bool {
	return GetCategory(req) == Event
}
