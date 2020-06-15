package iotmakerDocker

import (
	"errors"
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) ContainerInspectByName(name string) (err error, inspect types.ContainerJSON) {
	var id string

	err, id = el.ContainerFindIdByName(name)
	if err != nil {
		return
	}

	inspect, err = el.cli.ContainerInspect(el.ctx, id)

	return err, inspect
}
