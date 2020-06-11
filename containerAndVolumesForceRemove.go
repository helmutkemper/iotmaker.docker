package iotmaker_docker

import "github.com/docker/docker/api/types"

func (el *DockerSystem) ContainerAndVolumesForceRemove(id string) error {
	return el.cli.ContainerRemove(el.ctx, id, types.ContainerRemoveOptions{RemoveVolumes: true, Force: true})
}
