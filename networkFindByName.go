package iotmakerDocker

func (el *DockerSystem) NetworkFindByName(name string) (err error, id string) {
	err, list := el.NetworkList()
	if err != nil {
		return
	}

	for _, data := range list {
		if data.Name == name {
			id = data.ID
			return
		}
	}

	return
}
