package iotmaker_docker

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"time"
)

func (el *DockerSystem) ContainerStop(id string) error {
	var timeout = time.Microsecond * 1000
	return el.cli.ContainerStop(el.ctx, id, &timeout)
}

func (el *DockerSystem) ContainerStopAndRemove(id string, removeVolumes, removeLinks, force bool) error {
	var err error
	var timeout = time.Microsecond * 10000
	err = el.cli.ContainerStop(el.ctx, id, &timeout)
	if err != nil {
		return err
	}

	ok, notOk := el.cli.ContainerWait(el.ctx, id, "not-running")
	select {
	case o := <-ok:
		fmt.Printf("%v\n\n\n", o)
	case err = <-notOk:
		return err
	}

	err = el.cli.NetworkDisconnect(el.ctx, "network_test", id, force)

	time.Sleep(time.Second * 5)
	return el.cli.ContainerRemove(el.ctx, id, types.ContainerRemoveOptions{
		RemoveVolumes: removeVolumes,
		RemoveLinks:   removeLinks,
		Force:         force,
	})
}
