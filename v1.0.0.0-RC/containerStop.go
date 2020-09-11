package iotmakerdocker

import (
	"time"
)

func (el *DockerSystem) ContainerStop(
	id string,
) (
	err error,
) {

	var timeout = time.Microsecond * 1000
	return el.cli.ContainerStop(el.ctx, id, &timeout)
}
