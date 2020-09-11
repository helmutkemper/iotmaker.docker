package iotmakerdocker

import (
	"github.com/docker/docker/api/types"
)

// list images
func (el *DockerSystem) ImageList() (
	list []types.ImageSummary,
	err error,
) {

	list, err = el.cli.ImageList(el.ctx, types.ImageListOptions{})
	return
}
