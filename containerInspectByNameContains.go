package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) ContainerInspectByNameContains(
	name string,
) (
	err error,
	list []types.ContainerJSON,
) {

	list = make([]types.ContainerJSON, 0)
	var inspect types.ContainerJSON
	var listOfContainers []NameAndId

	err, listOfContainers = el.ContainerFindIdByNameContains(name)
	if err != nil {
		return
	}

	for _, v := range listOfContainers {
		err, inspect = el.ContainerInspect(v.ID)
		if err != nil {
			return
		}

		list = append(list, inspect)
	}

	return
}
