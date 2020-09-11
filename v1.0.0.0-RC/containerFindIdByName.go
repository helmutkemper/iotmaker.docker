package iotmakerdocker

import (
	"errors"
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) ContainerFindIdByName(
	name string,
) (
	ID string,
	err error,
) {

	var list []types.Container

	list, err = el.ContainerListAll()
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
