package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
)

// list exposed volumes from image by id
func (el *DockerSystem) ImageListExposedVolumes(id string) (error, []string) {
	var err error
	var imageData types.ImageInspect
	var ret = make([]string, 0)

	imageData, _, err = el.cli.ImageInspectWithRaw(el.ctx, id)
	if err != nil {
		return err, []string{}
	}
	for volume := range imageData.ContainerConfig.Volumes {
		ret = append(ret, volume)
	}

	return nil, ret
}
