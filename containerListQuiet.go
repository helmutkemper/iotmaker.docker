package iotmakerDocker

import "github.com/docker/docker/api/types"

func (el *DockerSystem) ContainerListQuiet() (
	err error,
	list []types.Container,
) {

	list, err = el.cli.ContainerList(el.ctx, types.ContainerListOptions{Quiet: true})

	return
}
