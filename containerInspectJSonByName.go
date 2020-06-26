package iotmakerDocker

func (el *DockerSystem) ContainerInspectJSonByName(name string) (err error, inspect []byte) {
	var id string

	err, id = el.ContainerFindIdByName(name)
	if err != nil {
		return
	}

	err, inspect = el.ContainerInspectJSon(id)

	return err, inspect
}
