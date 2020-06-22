package iotmakerDocker

import (
	"errors"
	"github.com/docker/docker/api/types"
	"strings"
)

func (el *DockerSystem) ContainerFindIdByNameContains(containsName string) (err error, id string) {
	var list []types.Container

	err, list = el.ContainerListAll()
	for _, containerData := range list {
		for _, containerName := range containerData.Names {
			if strings.Contains(containerName, containsName) == true {
				id = containerData.ID
				return
			}
		}
	}

	err = errors.New("container name not found")

	return
}
