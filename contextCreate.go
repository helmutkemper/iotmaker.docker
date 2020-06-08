package iotmaker_docker

import "context"

func (el *DockerSystem) contextCreate() {
	el.ctx = context.Background()
}
