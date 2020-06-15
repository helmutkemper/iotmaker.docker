package iotmakerDocker

import (
	"time"
)

func (el *DockerSystem) ContainerStop(id string) error {
	var timeout = time.Microsecond * 1000
	return el.cli.ContainerStop(el.ctx, id, &timeout)
}
