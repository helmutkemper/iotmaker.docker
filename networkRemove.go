package iotmaker_docker

import (
	"errors"
)

// remove network by name
func (el *DockerSystem) NetworkRemove(name string) error {
	_, found := el.networkId[name]
	if found != false {
		return errors.New("network name not found in network created list")
	}

	return el.cli.NetworkRemove(el.ctx, el.networkId[name])
}
