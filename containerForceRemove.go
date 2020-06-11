package iotmakerDocker

import "github.com/docker/docker/api/types"

func (el *DockerSystem) ContainerForceRemove(id string) error {
	return el.cli.ContainerRemove(el.ctx, id, types.ContainerRemoveOptions{Force: true})
}
