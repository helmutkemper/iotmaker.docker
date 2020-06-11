package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
)

// list image exposed ports by id
func (el *DockerSystem) ImageListExposedPorts(id string) (error, []string) {
	var err error
	var imageData types.ImageInspect
	var ret = make([]string, 0)

	imageData, _, err = el.cli.ImageInspectWithRaw(el.ctx, id)
	if err != nil {
		return err, []string{}
	}
	for port := range imageData.ContainerConfig.ExposedPorts {
		ret = append(ret, port.Port()+"/"+port.Proto())
	}

	return nil, ret
}
