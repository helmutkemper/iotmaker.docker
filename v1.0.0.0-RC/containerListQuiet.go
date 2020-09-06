package iotmakerDocker

import "github.com/docker/docker/api/types"

func (el *DockerSystem) ContainerListQuiet() (
	list []types.Container,
	err error,
) {

	list, err = el.cli.ContainerList(el.ctx, types.ContainerListOptions{Quiet: true})

	return
}
