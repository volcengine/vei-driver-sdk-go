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

type ErrorKind string

const (
	DeviceNotFound          ErrorKind = "设备未找到"
	DeviceLoadFailed        ErrorKind = "设备加载失败"
	ProtocolMissing         ErrorKind = "协议缺失"
	ProtocolParseFailed     ErrorKind = "协议解析失败"
	ProtocolUnsupported     ErrorKind = "协议暂不支持"
	AuthFailed              ErrorKind = "鉴权失败"
	LoginFailed             ErrorKind = "登录失败"
	WrongPassword           ErrorKind = "用户名密码错误"
	ConnectionFailed        ErrorKind = "无法建立连接"
	ConnectionTimeout       ErrorKind = "连接超时"
	NetworkUnreachable      ErrorKind = "网络不可达"
	ResourceNotFound        ErrorKind = "资源未找到"
	ResourceTypeUnsupported ErrorKind = "资源类型暂不支持"
	AttributeParseFailed    ErrorKind = "点表解析失败"
	ParameterParseFailed    ErrorKind = "参数解析失败"
	DataParseError          ErrorKind = "数据解析错误"
	ReadError               ErrorKind = "数据读取错误"
	ReadTimeout             ErrorKind = "数据读取超时"
	WriteError              ErrorKind = "数据写入错误"
	WriteTimeout            ErrorKind = "数据写入超时"
)

type Error struct {
	kind   ErrorKind
	reason string
}

func NewError(kind ErrorKind, err error) *Error {
	if err == nil {
		return NewErrorWithReason(kind, "Unknown")
	}
	return NewErrorWithReason(kind, err.Error())
}

func NewErrorWithReason(kind ErrorKind, reason string) *Error {
	return &Error{kind: kind, reason: reason}
}

func (e *Error) Error() string {
	return string(e.kind) + ": " + e.reason
}
