package iotmakerdocker

import (
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) ContainerStart(
	id string,
) (
	err error,
) {

	return el.cli.ContainerStart(el.ctx, id, types.ContainerStartOptions{})
}
