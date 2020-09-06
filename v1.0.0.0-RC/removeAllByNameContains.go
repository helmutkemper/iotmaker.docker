package iotmakerDocker

import "github.com/docker/docker/api/types"

// Use this function to remove trash after test.
// This function removes container, image and network by name, and unlinked volumes and
// imagens
func (el DockerSystem) RemoveAllByNameContains(name string) (err error) {
	var nameAndId []NameAndId
	var container types.ContainerJSON

	nameAndId, err = el.ContainerFindIdByNameContains(name)
	if err != nil && err.Error() != "container name not found" {
		return err
	}

	for _, data := range nameAndId {
		container, err = el.ContainerInspect(data.ID)
		if err != nil {
			return
		}

		if container.State != nil && container.State.Running == true {
			err = el.ContainerStopAndRemove(data.ID, true, false, false)
			if err != nil {
				return
			}
		}

		if container.State != nil && container.State.Running == false {
			err = el.ContainerRemove(data.ID, true, false, false)
			if err != nil {
				return
			}
		}
	}

	nameAndId, err = el.ImageFindIdByNameContains(name)
	if err != nil && err.Error() != "image name not found" {
		return err
	}
	for _, data := range nameAndId {
		err = el.ImageRemove(data.ID, false, false)
		if err != nil {
			return
		}
	}

	nameAndId, err = el.NetworkFindIdByNameContains(name)
	if err != nil && err.Error() != "network name not found" {
		return err
	}
	for _, data := range nameAndId {
		err = el.NetworkRemove(data.ID)
		if err != nil {
			return
		}
	}

	err = el.VolumesUnreferencedRemove()
	if err != nil {
		return
	}

	err = el.ImageGarbageCollector()
	if err != nil {
		return
	}

	return
}
