package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
)

// list images
func (el *DockerSystem) ImageList() (
	err error,
	list []types.ImageSummary,
) {

	list, err = el.cli.ImageList(el.ctx, types.ImageListOptions{})
	return
}
