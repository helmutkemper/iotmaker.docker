package iotmakerDocker

import "github.com/docker/docker/api/types"

// Remove unreferenced volumes
func (el DockerSystem) VolumesUnreferencedRemove() (err error) {
	var volumes []types.Volume

	err, volumes = el.VolumeList()
	if err != nil {
		return
	}

	for _, volumeData := range volumes {
		if volumeData.UsageData == nil || volumeData.UsageData.RefCount == -1 {
			// bug: volume do portainer tem volumeData.UsageData = nil
			_ = el.VolumeRemove(volumeData.Name)
		}
	}

	return
}
