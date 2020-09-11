package iotmakerdocker

import (
	"github.com/docker/go-connections/nat"
)

func (el *DockerSystem) convertPort(
	in nat.PortMap,
) (
	out nat.PortSet,
) {

	out = make(map[nat.Port]struct{})
	for k := range in {
		out[k] = struct{}{}
	}

	return
}
