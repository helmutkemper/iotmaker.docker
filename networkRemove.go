package iotmakerDocker

func (el *DockerSystem) NetworkRemove(id string) error {
	return el.cli.NetworkRemove(el.ctx, id)
}
