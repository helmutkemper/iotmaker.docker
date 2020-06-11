package iotmakerDocker

import "github.com/docker/docker/api/types"

func (el *DockerSystem) ContainerForceRemoveWithLinksAndVolumes(id string) error {
	return el.cli.ContainerRemove(el.ctx, id, types.ContainerRemoveOptions{RemoveLinks: true, RemoveVolumes: true, Force: true})
}
