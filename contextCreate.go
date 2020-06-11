package iotmakerDocker

import "context"

func (el *DockerSystem) contextCreate() {
	el.ctx = context.Background()
}
