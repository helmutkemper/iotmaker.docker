package iotmakerDocker

import (
	"errors"
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) ContainerStatisticsOneShotByName(name string) (error, types.Stats) {
	var err error
	var list []types.Container
	var ret types.Stats
	var pass bool
	var id string

	err, list = el.ContainerListAll()
	if err != nil {
		return err, ret
	}

	for _, containerData := range list {
		for _, containerName := range containerData.Names {
			if containerName == name {
				pass = true
				id = containerData.ID
				break
			}
		}
	}

	if pass == false {
		return errors.New("container name not found"), ret
	}

	return el.ContainerStatisticsOneShot(id)
}
