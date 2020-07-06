package iotmakerDocker

import "github.com/docker/docker/api/types"

func (el *DockerSystem) ImageGarbageCollector() (err error) {
	var list []types.ImageSummary
	err, list = el.ImageList()
	for _, image := range list {
		if len(image.RepoTags) > 0 {
			if image.RepoTags[0] == "<none>:<none>" {
				err = el.ImageRemove(image.ID)
				if err != nil {
					return
				}
			}
		}
	}

	return
}
