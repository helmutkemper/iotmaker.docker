package iotmakerDocker

func (el *DockerSystem) ImageRemoveByName(name string, force, pruneChildren bool) error {
	var err error
	var id string

	err, id = el.ImageFindIdByName(name)

	if err != nil {
		return err
	}

	err = el.ImageRemove(id, force, pruneChildren)

	return err
}
