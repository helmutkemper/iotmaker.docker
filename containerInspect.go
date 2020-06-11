package iotmakerDocker

import "github.com/docker/docker/api/types"

func (el *DockerSystem) ContainerInspect(id string) (error, types.ContainerJSON) {
	var err error
	var inspect types.ContainerJSON

	inspect, err = el.cli.ContainerInspect(el.ctx, id)

	return err, inspect
}
