package iotmakerDocker

func (el *DockerSystem) ContainerWait(id string) (err error) {
	wOk, wErr := el.cli.ContainerWait(el.ctx, id, "not-running")
	select {
	case <-wOk:

	case err = <-wErr:

	}

	return
}
