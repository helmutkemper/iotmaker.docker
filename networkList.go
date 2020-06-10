package iotmaker_docker

import (
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) NetworkList() (err error, netList []types.NetworkResource) {
	netList, err = el.cli.NetworkList(el.ctx, types.NetworkListOptions{})

	return
}
