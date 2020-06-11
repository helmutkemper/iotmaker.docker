package iotmakerDocker

// verify if exposed volume (folder only) defined by user is exposed
// in image
func (el *DockerSystem) ImageVerifyVolume(id, path string) (error, bool) {
	err, list := el.ImageListExposedVolumes(id)
	if err != nil {
		return err, false
	}

	for _, volume := range list {
		if volume == path {
			return nil, true
		}
	}

	return nil, false
}
