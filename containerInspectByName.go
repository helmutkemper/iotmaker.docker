package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) ContainerInspectByName(name string) (err error, inspect types.ContainerJSON) {
	var id string

	err, id = el.ContainerFindIdByName(name)
	if err != nil {
		return
	}

	err, inspect = el.ContainerInspect(id)

	return err, inspect
}
