package iotmakerDocker

func (el *DockerSystem) ContainerInspectJSonByName(
	name string,
) (
	inspect []byte,
	err error,
) {

	var id string

	id, err = el.ContainerFindIdByName(name)
	if err != nil {
		return
	}

	inspect, err = el.ContainerInspectJSon(id)

	return inspect, err
}
