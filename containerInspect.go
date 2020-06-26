package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) ContainerInspect(
	id string,
) (err error, inspect types.ContainerJSON) {

	inspect, err = el.cli.ContainerInspect(el.ctx, id)
	if err != nil {
		return
	}

	return
}
