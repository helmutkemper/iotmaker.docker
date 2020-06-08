package iotmaker_docker

import (
	"github.com/docker/docker/api/types"
)

// create network
func (el *DockerSystem) NetworkCreate(name string) error {
	resp, err := el.cli.NetworkCreate(el.ctx, name, types.NetworkCreate{
		Labels: map[string]string{
			"name": name,
		},
	})

	if err != nil {
		return err
	}

	if len(el.networkId) == 0 {
		el.networkId = make(map[string]string)
	}

	el.networkId[name] = resp.ID

	return err
}
