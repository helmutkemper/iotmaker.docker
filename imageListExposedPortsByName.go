package iotmaker_docker

// list image exposed ports by name
func (el *DockerSystem) ImageListExposedPortsByName(name string) (error, []string) {
	var err error
	var id string
	err, id = el.ImageFindIdByName(name)
	if err != nil {
		return err, nil
	}

	return el.ImageListExposedPorts(id)
}
