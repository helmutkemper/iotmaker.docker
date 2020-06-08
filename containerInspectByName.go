package iotmaker_docker

import (
	"errors"
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) ContainerInspectByName(name string) (error, types.ContainerJSON) {
	var err error
	var list []types.Container
	var inspect types.ContainerJSON
	var pass bool

	err, list = el.ContainerListAll()
	if err != nil {
		return err, inspect
	}

	for _, containerData := range list {
		for _, containerName := range containerData.Names {
			if containerName == name {
				pass = true
				inspect, err = el.cli.ContainerInspect(el.ctx, containerData.ID)
			}
		}
	}

	if pass == false {
		err = errors.New("container name not found")
	}

	return err, inspect
}
