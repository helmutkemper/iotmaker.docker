package iotmakerdocker

import (
	"github.com/docker/go-connections/nat"
)

// Mount nat por list by image config
func (el *DockerSystem) ImageMountNatPortList(
	imageId string,
) (
	nat.PortMap,
	error,
) {

	var err error
	var portList []nat.Port
	var ret nat.PortMap = make(map[nat.Port][]nat.PortBinding)

	portList, err = el.ImageListExposedPorts(imageId)
	if err != nil {
		return nat.PortMap{}, err
	}

	for _, port := range portList {
		ret[port] = []nat.PortBinding{
			{
				HostPort: port.Port() + "/" + port.Proto(),
			},
		}
	}

	return ret, err
}
