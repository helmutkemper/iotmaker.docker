package iotmakerdocker

import (
	"github.com/docker/docker/api/types"
)

// ContainerStart (English): start a container by id
//   id: string container id
//
// ContainerStart (Português): inicia um container por id
//   id: string container id
func (el *DockerSystem) ContainerStart(
	id string,
) (
	err error,
) {

	return el.cli.ContainerStart(el.ctx, id, types.ContainerStartOptions{})
}
