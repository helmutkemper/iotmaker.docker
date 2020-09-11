package iotmakerdocker

import "github.com/docker/docker/api/types"

func (el *DockerSystem) ContainerListAll() (
	list []types.Container,
	err error,
) {

	list, err = el.cli.ContainerList(el.ctx, types.ContainerListOptions{All: true})
	return
}
