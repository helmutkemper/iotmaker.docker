package iotmakerdocker

func (el *DockerSystem) ContainerInspectJSon(
	id string,
) (
	inspect []byte,
	err error,
) {

	_, inspect, err = el.cli.ContainerInspectWithRaw(el.ctx, id, true)
	if err != nil {
		return
	}

	return
}
