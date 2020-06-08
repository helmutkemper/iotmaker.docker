package iotmaker_docker

import (
	"github.com/docker/docker/api/types"
)

// list images
func (el *DockerSystem) ImageList() (error, []types.ImageSummary) {
	ret, err := el.cli.ImageList(el.ctx, types.ImageListOptions{})
	return err, ret
}
