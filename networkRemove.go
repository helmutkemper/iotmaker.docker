package iotmaker_docker

func (el *DockerSystem) NetworkRemove(id string) error {
	return el.cli.NetworkRemove(el.ctx, id)
}
