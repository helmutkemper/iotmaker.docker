package util

import (
	"github.com/docker/docker/api/types/network"
)

type NetworkGenerator struct {
	ip   IPv4Generator
	name string
}

// init a network for new container
// nInit("test", 10, 0, 0, 1)
// before use this function, use whaleAquarium.Docker.NetworkCreate("test")
func (el *NetworkGenerator) Init(name string, a, b, c, d byte) {
	el.name = name
	el.ip.Init(a, b, c, d)
}

func (el *NetworkGenerator) GetNext() (error, *network.NetworkingConfig) {
	var err = el.ip.Inc()
	return err, &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			el.name: {
				IPAddress: el.ip.String(),
			},
		},
	}
}
