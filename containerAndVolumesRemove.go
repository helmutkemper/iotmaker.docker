package iotmaker_docker

import "github.com/docker/docker/api/types"

func (el *DockerSystem) ContainerAndVolumesRemove(id string) error {
	return el.cli.ContainerRemove(el.ctx, id, types.ContainerRemoveOptions{RemoveVolumes: true})
}
