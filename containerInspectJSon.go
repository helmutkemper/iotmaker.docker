package iotmakerDocker

func (el *DockerSystem) ContainerInspectJSon(
	id string,
) (err error, inspect []byte) {

	_, inspect, err = el.cli.ContainerInspectWithRaw(el.ctx, id, true)
	if err != nil {
		return
	}

	return
}
