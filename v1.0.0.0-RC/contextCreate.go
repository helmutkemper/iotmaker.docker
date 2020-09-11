package iotmakerdocker

import "context"

func (el *DockerSystem) ContextCreate() {
	el.ctx = context.Background()
}
