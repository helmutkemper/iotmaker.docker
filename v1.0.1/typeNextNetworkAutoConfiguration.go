package iotmakerdocker

import (
	"errors"
	"github.com/docker/docker/api/types/network"
)

type NextNetworkAutoConfiguration struct {
	ip      IPv4Generator
	id      string
	name    string
	gateway string
	err     error
}

// init a network for new container
// nInit("test", 10, 0, 0, 1)
// before use this function, use whaleAquarium.Docker.NetworkCreate("test")
func (el *NextNetworkAutoConfiguration) Init(id, name, gateway string, a, b, c, d byte) {
	el.id = id
	el.name = name
	el.gateway = gateway
	el.ip.Init(a, b, c, d)

	el.err = errors.New("run GetNext() function before get a valid ip address")
}

func (el *NextNetworkAutoConfiguration) GetNext() (*network.NetworkingConfig, error) {
	el.err = el.ip.Inc()

	newIp := el.ip.String()
	return &network.NetworkingConfig{
			EndpointsConfig: map[string]*network.EndpointSettings{
				el.name: {
					NetworkID: el.id,
					Gateway:   el.gateway,
					IPAMConfig: &network.EndpointIPAMConfig{
						IPv4Address: newIp,
					},
					IPAddress: newIp,
				},
			},
		},
		el.err
}

func (el *NextNetworkAutoConfiguration) GetCurrentIpAddress() (IP string, err error) {
	return el.ip.String(), el.err
}
