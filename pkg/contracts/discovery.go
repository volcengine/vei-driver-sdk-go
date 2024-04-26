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

import (
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
}

// MulticastParameter is the input configuration for a Multicast Discovery.
type MulticastParameter struct {
	// The target ethernet interface for multicast discovering
	EthernetInterface string
}

// NetScanParameter is the input configuration for a NetScan Discovery.
type NetScanParameter struct {
	// Protocol is the type of probe to make: tcp, udp, etc.
	Protocol string
	// List of IPv4 subnets to perform netscan discovery on, in CIDR format (X.X.X.X/Y).
	Subnets []string
	// ScanPorts is a slice of ports to scan for on each host.
	ScanPorts []string
	// Maximum simultaneous network probes when running netscan discovery.
	ProbeAsyncLimit int
	// Maximum amount of milliseconds to wait for each IP probe before timing out.
	ProbeTimeout time.Duration
}

// SerialScanParameter is the input configuration for a SerialScan Discovery.
type SerialScanParameter struct {
	// The serial port bitrate, e.g. 9600, 19200, ...
	BaudRate int
	// Size of the character (must be 5, 6, 7 or 8)
	DataBits int
	// The serial port stop bits setting (must be 1 or 2)
	StopBits int
	// The serial port parity setting (must be none, even or odd)
	Parity string
}
