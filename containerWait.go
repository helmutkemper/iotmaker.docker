package iotmakerDocker

import (
	"fmt"
)

//todo: fazer
func (el *DockerSystem) ContainerWait(id string) {
	wOk, wErr := el.cli.ContainerWait(el.ctx, id, "not-running")
	select {
	case <-wOk:
		fmt.Println()
	case <-wErr:
		fmt.Println()
	}
}
