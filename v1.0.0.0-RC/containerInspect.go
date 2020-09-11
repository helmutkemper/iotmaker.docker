package iotmakerdocker

import (
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) ContainerInspect(
	id string,
) (
	inspect types.ContainerJSON,
	err error,
) {

	inspect, err = el.cli.ContainerInspect(el.ctx, id)
	if err != nil {
		return
	}

	return
}
