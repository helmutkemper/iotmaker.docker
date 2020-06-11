package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	volumeTypes "github.com/docker/docker/api/types/volume"
)

func (el *DockerSystem) VolumeList() (err error, volList []types.Volume) {
	var list volumeTypes.VolumeListOKBody
	list, err = el.cli.VolumeList(el.ctx, filters.Args{})

	for _, data := range list.Volumes {
		volList = append(volList, *data)
	}

	return
}
