package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) VolumeFindByName(name string) (err error, volume types.Volume) {
	err, list := el.VolumeList()
	for _, data := range list {
		if data.Name == name {
			volume = data
			return
		}
	}
	return
}
