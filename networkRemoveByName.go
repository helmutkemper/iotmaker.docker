package iotmakerDocker

func (el *DockerSystem) NetworkRemoveByName(name string) error {
	err, id := el.NetworkFindByName(name)
	if err != nil {
		return err
	}

	return el.cli.NetworkRemove(el.ctx, id)
}
