package iotmaker_docker

import (
	"errors"
)

// find image id by name
func (el *DockerSystem) ImageFindIdByName(name string) (error, string) {
	err, list := el.ImageList()
	if err != nil {
		return err, ""
	}

	if len(el.imageId) == 0 {
		el.imageId = make(map[string]string)
	}

	for _, data := range list {
		for _, dataTag := range data.RepoTags {
			if dataTag == name {
				el.imageId[name] = data.ID
				return nil, data.ID
			}
		}
	}

	return errors.New("image name not found"), ""
}
