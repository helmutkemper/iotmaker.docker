package iotmakerDocker

import (
	"errors"
	"github.com/docker/docker/api/types"
	"strings"
)

func (el *DockerSystem) ImageFindIdByNameContains(
	name string,
) (
	err error,
	list []NameAndId,
) {

	var listTmp []types.ImageSummary
	list = make([]NameAndId, 0)

	err, listTmp = el.ImageList()
	if err != nil {
		return
	}

	if len(el.imageId) == 0 {
		el.imageId = make(map[string]string)
	}

	for _, data := range listTmp {
		if len(data.RepoTags) == 0 {
			continue
		}

		var tag = data.RepoTags[0]
		el.imageId[tag] = data.ID
		if strings.Contains(tag, name) == true {
			list = append(list, NameAndId{
				ID:   data.ID,
				Name: tag,
			})
		}
	}

	if len(list) == 0 {
		err = errors.New("image name not found")
	}

	return
}
