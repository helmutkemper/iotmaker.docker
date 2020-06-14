package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) ImageRemove(id string) error {
	var err error
	_, err = el.cli.ImageRemove(el.ctx, id, types.ImageRemoveOptions{})

	return err
}
