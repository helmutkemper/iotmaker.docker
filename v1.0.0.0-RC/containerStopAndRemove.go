package iotmakerDocker

import (
	"time"
)

func (el *DockerSystem) ContainerStopAndRemove(
	id string,
	removeVolumes,
	removeLinks,
	force bool,
) (
	err error,
) {

	var timeout = time.Microsecond * 10000
	err = el.cli.ContainerStop(el.ctx, id, &timeout)
	if err != nil {
		return err
	}

	ok, notOk := el.cli.ContainerWait(el.ctx, id, "not-running")
	select {
	case <-ok:
		break
	case err = <-notOk:
		return err
	}

	time.Sleep(time.Second * 5)
	return el.ContainerRemove(id, removeVolumes, removeLinks, force)
}
