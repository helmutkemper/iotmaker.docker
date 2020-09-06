package iotmakerDocker

// list exposed volumes from image by name
func (el *DockerSystem) ImageListExposedVolumesByName(
	name string,
) (
	list []string,
	err error,
) {

	var id string
	id, err = el.ImageFindIdByName(name)
	if err != nil {
		return nil, err
	}

	return el.ImageListExposedVolumes(id)
}
