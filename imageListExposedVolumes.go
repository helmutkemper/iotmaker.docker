package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
)

// list exposed volumes from image by id
func (el *DockerSystem) ImageListExposedVolumes(
	id string,
) (
	err error,
	list []string,
) {

	var imageData types.ImageInspect
	list = make([]string, 0)

	imageData, _, err = el.cli.ImageInspectWithRaw(el.ctx, id)
	if err != nil {
		return err, []string{}
	}
	for volume := range imageData.ContainerConfig.Volumes {
		list = append(list, volume)
	}

	return
}
