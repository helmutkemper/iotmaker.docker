package iotmakerDocker

import "github.com/docker/go-connections/nat"

// list image exposed ports by name
func (el *DockerSystem) ImageListExposedPortsByName(
	name string,
) (
	err error,
	portList []nat.Port,
) {

	var id string
	err, id = el.ImageFindIdByName(name)
	if err != nil {
		return err, nil
	}

	err, portList = el.ImageListExposedPorts(id)

	return
}
