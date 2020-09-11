package iotmakerdocker

import "github.com/docker/go-connections/nat"

// list image exposed ports by name
func (el *DockerSystem) ImageListExposedPortsByName(
	name string,
) (
	portList []nat.Port,
	err error,
) {

	var id string
	id, err = el.ImageFindIdByName(name)
	if err != nil {
		return nil, err
	}

	portList, err = el.ImageListExposedPorts(id)

	return
}
