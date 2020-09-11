package iotmakerdocker

import (
	"errors"
	"github.com/docker/docker/api/types"
	"strings"
)

func (el *DockerSystem) ContainerFindIdByNameContains(
	containsName string,
) (
	list []NameAndId,
	err error,
) {

	list = make([]NameAndId, 0)
	var listOfContainers []types.Container

	listOfContainers, err = el.ContainerListAll()
	if err != nil {
		return
	}

	for _, containerData := range listOfContainers {
		for _, containerName := range containerData.Names {
			if strings.Contains(containerName, containsName) == true {
				list = append(list, NameAndId{
					ID:   containerData.ID,
					Name: containerName,
				})
			}
		}
	}

	if len(list) == 0 {
		err = errors.New("container name not found")
	}

	return
}
