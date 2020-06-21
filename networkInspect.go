package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) NetworkInspect(
	id string,
) (err error, inspect types.NetworkResource) {

	inspect, err = el.cli.NetworkInspect(el.ctx, id, types.NetworkInspectOptions{})
	return err, inspect
}
