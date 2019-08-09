/*
 * Copyright (C) 2018 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package service

import (
	"encoding/json"

	"gopkg.in/urfave/cli.v1"
	"gopkg.in/urfave/cli.v1/altsrc"

	"github.com/mysteriumnetwork/node/core/service"
)

// Options describes options which are required to start Openvpn service
type Options struct {
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
	Subnet   string `json:"subnet"`
	Netmask  string `json:"netmask"`
}

var (
	protocolFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  "openvpn.proto",
		Usage: "Openvpn protocol to use. Options: { udp, tcp }",
		Value: defaultOptions.Protocol,
	})
	portFlag = altsrc.NewIntFlag(cli.IntFlag{
		Name:  "openvpn.port",
		Usage: "Openvpn port to use. If not specified, random port will be used",
		Value: defaultOptions.Port,
	})
	subnetFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  "openvpn.subnet",
		Usage: "Openvpn subnet that will be used to connecting VPN clients",
		Value: defaultOptions.Subnet,
	})
	netmaskFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  "openvpn.netmask",
		Usage: "Openvpn subnet netmask ",
		Value: defaultOptions.Netmask,
	})
	defaultOptions = Options{
		Protocol: "udp",
		Port:     0,
		Subnet:   "10.8.0.0",
		Netmask:  "255.255.255.0",
	}
)

// RegisterFlags function register Openvpn flags to flag list
func RegisterFlags(flags *[]cli.Flag) {
	*flags = append(*flags, protocolFlag, portFlag, subnetFlag, netmaskFlag)
}

// ParseFlags function fills in Openvpn options from CLI context
func ParseFlags(ctx *cli.Context) service.Options {
	return Options{
		Protocol: ctx.String(protocolFlag.Name),
		Port:     ctx.Int(portFlag.Name),
		Subnet:   ctx.String(subnetFlag.Name),
		Netmask:  ctx.String(netmaskFlag.Name),
	}
}

// ParseJSONOptions function fills in Openvpn options from JSON request
func ParseJSONOptions(request *json.RawMessage) (service.Options, error) {
	if request == nil {
		return defaultOptions, nil
	}

	opts := defaultOptions
	err := json.Unmarshal(*request, &opts)
	return opts, err
}
