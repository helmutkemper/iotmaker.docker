package iotmakerDocker

// list exposed volumes from image by name
func (el *DockerSystem) ImageListExposedVolumesByName(name string) (error, []string) {
	var err error
	var id string
	err, id = el.ImageFindIdByName(name)
	if err != nil {
		return err, nil
	}

	return el.ImageListExposedVolumes(id)
}
