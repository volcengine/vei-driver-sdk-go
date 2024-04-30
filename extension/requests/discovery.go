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
	"encoding/json"
	"fmt"
	"time"
)

type DiscoveryMode string

const (
	Multicast  DiscoveryMode = "multicast"
	NetScan    DiscoveryMode = "netscan"
	SerialScan DiscoveryMode = "serialscan"
)

type DiscoveryParameter struct {
	DiscoveryMode       DiscoveryMode        `json:"discovery_mode"`
	MaxDurationTime     time.Duration        `json:"max_duration_time"`
	MulticastParameter  *MulticastParameter  `json:"multicast_parameter,omitempty"`
	NetScanParameter    *NetScanParameter    `json:"net_scan_parameter,omitempty"`
	SerialScanParameter *SerialScanParameter `json:"serial_scan_parameter,omitempty"`
	Credentials         map[string]string    `json:"credentials,omitempty"`
}

// MulticastParameter is the input configuration for a Multicast Discovery.
type MulticastParameter struct {
	// The target ethernet interface for multicast discovering
	EthernetInterface string `json:"ethernet_interface"`
}

// NetScanParameter is the input configuration for a NetScan Discovery.
type NetScanParameter struct {
	// Protocol is the type of probe to make: tcp, udp, etc.
	Protocol string `json:"protocol"`
	// List of IPv4 subnets to perform netscan discovery on, in CIDR format (X.X.X.X/Y).
	Subnets []string `json:"subnets"`
	// ScanPorts is a slice of ports to scan for on each host.
	ScanPorts []string `json:"scan_ports"`
	// Maximum simultaneous network probes when running netscan discovery.
	ProbeAsyncLimit int `json:"probe_async_limit"`
	// Maximum amount of milliseconds to wait for each IP probe before timing out.
	ProbeTimeout time.Duration `json:"probe_timeout"`
}

// SerialScanParameter is the input configuration for a SerialScan Discovery.
type SerialScanParameter struct {
	// The serial port bitrate, e.g. 9600, 19200, ...
	BaudRate int `json:"baud_rate"`
	// Size of the character (must be 5, 6, 7 or 8)
	DataBits int `json:"data_bits"`
	// The serial port stop bits setting (must be 1 or 2)
	StopBits int `json:"stop_bits"`
	// The serial port parity setting (must be none, even or odd)
	Parity string `json:"parity"`
}

func (dp *DiscoveryParameter) Valid() error {
	switch dp.DiscoveryMode {
	case Multicast:
		if dp.MulticastParameter == nil || dp.MulticastParameter.EthernetInterface == "" {
			return fmt.Errorf("multicast parameter invalid")
		}
	case NetScan:
		if dp.NetScanParameter == nil {
			return fmt.Errorf("netscan parameter invalid")
		}
		if len(dp.NetScanParameter.Subnets) == 0 || len(dp.NetScanParameter.ScanPorts) == 0 {
			return fmt.Errorf("the subnets or scanports in netscan parameter can not be empty")
		}
		if dp.NetScanParameter.Protocol == "" {
			dp.NetScanParameter.Protocol = "udp"
		}
		if dp.NetScanParameter.ProbeAsyncLimit == 0 {
			dp.NetScanParameter.ProbeAsyncLimit = 4000
		}
		if dp.NetScanParameter.ProbeTimeout == 0 {
			dp.NetScanParameter.ProbeTimeout = 2000
		}
	case SerialScan:
		if dp.SerialScanParameter == nil || dp.SerialScanParameter.BaudRate == 0 ||
			dp.SerialScanParameter.DataBits == 0 || dp.SerialScanParameter.StopBits == 0 ||
			dp.SerialScanParameter.Parity == "" {
			return fmt.Errorf("serialscan parameter invalid")
		}
	default:
		return fmt.Errorf("discovery mode %s not supported", dp.DiscoveryMode)
	}
	return nil
}

func (dp *DiscoveryParameter) String() string {
	data, _ := json.Marshal(dp)
	return string(data)
}
