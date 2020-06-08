package iotmaker_docker

import (
	"github.com/docker/docker/api/types"
)

// verify if network name exists
func (el *DockerSystem) NetworkVerifyName(name string) (error, bool) {
	resp, err := el.cli.NetworkList(el.ctx, types.NetworkListOptions{})
	if err != nil {
		return err, false
	}

	for _, v := range resp {
		if v.Name == name {
			return nil, true
		}
	}

	return nil, false
}
