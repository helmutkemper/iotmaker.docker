package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) ImageRemoveByName(name string) error {
	var err error
	var id string

	err, id = el.ImageFindIdByName(name)

	if err != nil {
		return err
	}

	_, err = el.cli.ImageRemove(el.ctx, id, types.ImageRemoveOptions{})

	return err
}
