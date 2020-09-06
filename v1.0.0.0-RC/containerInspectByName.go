package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) ContainerInspectByName(
	name string,
) (
	inspect types.ContainerJSON,
	err error,
) {

	var id string

	id, err = el.ContainerFindIdByName(name)
	if err != nil {
		return
	}

	inspect, err = el.ContainerInspect(id)

	return inspect, err
}
