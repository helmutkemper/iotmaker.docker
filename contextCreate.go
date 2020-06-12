package iotmakerDocker

import "context"

func (el *DockerSystem) ContextCreate() {
	el.ctx = context.Background()
}
