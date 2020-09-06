package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
)

// list exposed volumes from image by id
func (el *DockerSystem) ImageListExposedVolumes(
	id string,
) (
	list []string,
	err error,
) {

	var imageData types.ImageInspect
	list = make([]string, 0)

	imageData, _, err = el.cli.ImageInspectWithRaw(el.ctx, id)
	if err != nil {
		return []string{}, err
	}
	for volume := range imageData.ContainerConfig.Volumes {
		list = append(list, volume)
	}

	return
}
