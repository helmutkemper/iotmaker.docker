package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) VolumeFindByName(
	name string,
) (
	volume types.Volume,
	err error,
) {

	var list []types.Volume
	list, err = el.VolumeList()
	for _, data := range list {
		if data.Name == name {
			volume = data
			return
		}
	}
	return
}
