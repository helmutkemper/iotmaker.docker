package iotmakerDocker

import (
	"errors"
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) NetworkFindIdByName(
	name string,
) (
	err error,
	id string,
) {

	var list []types.NetworkResource

	err, list = el.NetworkList()
	if err != nil {
		return
	}

	for _, data := range list {
		if data.Name == name {
			id = data.ID
			return
		}
	}

	err = errors.New("network not found")

	return
}
