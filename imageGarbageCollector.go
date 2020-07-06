package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) ImageGarbageCollector() (err error) {
	var list []types.ImageSummary
	err, list = el.ImageList()
	for _, image := range list {
		if len(image.RepoTags) > 0 {
			if image.RepoTags[0] == "<none>:<none>" {
				//id := strings.Split(image.ID, ":")
				//err = el.ImageRemove(id[len(id)-1], true, true)
				err = el.ImageRemove(image.ID, true, true)
				if err != nil {
					return
				}
			}
		}
	}

	return
}
