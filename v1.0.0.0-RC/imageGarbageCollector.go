package iotmakerdocker

import (
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) ImageGarbageCollector() (err error) {
	var list []types.ImageSummary
	list, err = el.ImageList()
	for _, image := range list {
		if len(image.RepoTags) > 0 {
			if image.RepoTags[0] == "<none>:<none>" {
				err = el.ImageRemove(image.ID, true, true)
				if err != nil {
					return
				}
			}
		}
	}

	return
}
