package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
)

// list image exposed ports by id
func (el *DockerSystem) ImageListExposedPorts(
	id string,
) (
	portList []nat.Port,
	err error,
) {

	var imageData types.ImageInspect

	imageData, _, err = el.cli.ImageInspectWithRaw(el.ctx, id)
	if err != nil {
		return
	}
	for port := range imageData.ContainerConfig.ExposedPorts {
		portList = append(portList, port)
	}

	return
}
