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

package contracts

// ValueType indicates the type of reading value
type ValueType string

const (
	Bool         ValueType = "Bool"
	String       ValueType = "String"
	Uint8        ValueType = "Uint8"
	Uint16       ValueType = "Uint16"
	Uint32       ValueType = "Uint32"
	Uint64       ValueType = "Uint64"
	Int8         ValueType = "Int8"
	Int16        ValueType = "Int16"
	Int32        ValueType = "Int32"
	Int64        ValueType = "Int64"
	Float32      ValueType = "Float32"
	Float64      ValueType = "Float64"
	Binary       ValueType = "Binary"
	BoolArray    ValueType = "BoolArray"
	StringArray  ValueType = "StringArray"
	Uint8Array   ValueType = "Uint8Array"
	Uint16Array  ValueType = "Uint16Array"
	Uint32Array  ValueType = "Uint32Array"
	Uint64Array  ValueType = "Uint64Array"
	Int8Array    ValueType = "Int8Array"
	Int16Array   ValueType = "Int16Array"
	Int32Array   ValueType = "Int32Array"
	Int64Array   ValueType = "Int64Array"
	Float32Array ValueType = "Float32Array"
	Float64Array ValueType = "Float64Array"
	Object       ValueType = "Object"
)

func (t ValueType) String() string {
	return string(t)
}
