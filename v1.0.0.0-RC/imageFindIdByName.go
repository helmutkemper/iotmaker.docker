package iotmakerDocker

import (
	"errors"
	"github.com/docker/docker/api/types"
)

// find image id by name
func (el *DockerSystem) ImageFindIdByName(
	name string,
) (
	ID string,
	err error,
) {

	var list []types.ImageSummary

	list, err = el.ImageList()
	if err != nil {
		return "", err
	}

	if len(el.imageId) == 0 {
		el.imageId = make(map[string]string)
	}

	for _, data := range list {
		for _, dataTag := range data.RepoTags {
			if dataTag == name {
				el.imageId[name] = data.ID
				return data.ID, nil
			}
		}
	}

	return "", errors.New("image name not found")
}
