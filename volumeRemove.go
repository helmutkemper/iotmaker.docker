package iotmaker_docker

func (el *DockerSystem) VolumeRemove(id string) (err error) {
	err = el.cli.VolumeRemove(el.ctx, id, false)

	return
}
