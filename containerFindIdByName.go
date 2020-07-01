package iotmakerDocker

import (
	"errors"
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) ContainerFindIdByName(
	name string,
) (
	err error,
	ID string,
) {

	var list []types.Container

	err, list = el.ContainerListAll()
	for _, containerData := range list {
		for _, containerName := range containerData.Names {
			if containerName == name || containerName == "/"+name {
				ID = containerData.ID
				return
			}
		}
	}

	err = errors.New("container name not found")

	return
}
