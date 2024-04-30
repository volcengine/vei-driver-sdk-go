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

package requests

import (
	"testing"
)

func TestDiscoveryParameter_Valid(t *testing.T) {
	tests := []struct {
		name       string
		mode       DiscoveryMode
		multicast  *MulticastParameter
		netscan    *NetScanParameter
		serialscan *SerialScanParameter
		wantErr    bool
	}{
		{name: "unsupported mode", mode: "other", wantErr: true},
		{name: "invalid multicast", mode: Multicast, multicast: &MulticastParameter{EthernetInterface: ""}, wantErr: true},
		{name: "valid multicast", mode: Multicast, multicast: &MulticastParameter{EthernetInterface: "eth0"}, wantErr: false},
		{name: "nil netscan", mode: NetScan, netscan: nil, wantErr: true},
		{name: "invalid netscan", mode: NetScan, netscan: &NetScanParameter{}, wantErr: true},
		{name: "valid netscan", mode: NetScan, netscan: &NetScanParameter{Subnets: []string{"192.168.1.0/24"}, ScanPorts: []string{"3702"}}, wantErr: false},
		{name: "nil serialscan", mode: SerialScan, serialscan: nil, wantErr: true},
		{name: "invalid serialscan", mode: SerialScan, serialscan: &SerialScanParameter{}, wantErr: true},
		{name: "valid serialscan", mode: SerialScan, serialscan: &SerialScanParameter{BaudRate: 9600, DataBits: 8, StopBits: 1, Parity: "none"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dp := &DiscoveryParameter{
				DiscoveryMode:       tt.mode,
				MulticastParameter:  tt.multicast,
				NetScanParameter:    tt.netscan,
				SerialScanParameter: tt.serialscan,
			}
			if err := dp.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDiscoveryParameter_String(t *testing.T) {
	dp := &DiscoveryParameter{
		DiscoveryMode:      NetScan,
		MaxDurationTime:    30,
		MulticastParameter: &MulticastParameter{EthernetInterface: "eth0"},
		NetScanParameter: &NetScanParameter{
			Protocol:        "udp",
			Subnets:         []string{"192.168.1.0/24"},
			ScanPorts:       []string{"3702"},
			ProbeAsyncLimit: 4000,
			ProbeTimeout:    2000,
		},
		SerialScanParameter: &SerialScanParameter{
			BaudRate: 9600,
			DataBits: 7,
			StopBits: 1,
			Parity:   "none",
		},
	}
	t.Log(dp.String())
}
