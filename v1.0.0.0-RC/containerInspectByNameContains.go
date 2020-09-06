package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) ContainerInspectByNameContains(
	name string,
) (
	list []types.ContainerJSON,
	err error,
) {

	list = make([]types.ContainerJSON, 0)
	var inspect types.ContainerJSON
	var listOfContainers []NameAndId

	listOfContainers, err = el.ContainerFindIdByNameContains(name)
	if err != nil {
		return
	}

	for _, v := range listOfContainers {
		inspect, err = el.ContainerInspect(v.ID)
		if err != nil {
			return
		}

		list = append(list, inspect)
	}

	return
}
