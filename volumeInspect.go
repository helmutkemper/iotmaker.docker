package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) VolumeInspect(ID string) (err error, inspect types.Volume) {
	inspect, err = el.cli.VolumeInspect(el.ctx, ID)

	return
}
