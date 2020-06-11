package iotmakerDocker

import "github.com/docker/docker/api/types"

func (el *DockerSystem) ContainerAndLinksForceRemove(id string) error {
	return el.cli.ContainerRemove(el.ctx, id, types.ContainerRemoveOptions{RemoveLinks: true, Force: true})
}
