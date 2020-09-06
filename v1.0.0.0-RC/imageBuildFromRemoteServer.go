package iotmakerDocker

import (
	"errors"
	"github.com/docker/docker/api/types"
	"io"
)

// en: Make a image from folder path content
//     Please note: dockerfile name must be "Dockerfile" inside root folder
//
//     For get a github token
//     settings > Developer settings > Personal access tokens > Generate new token
//     Mark [x]repo - Full control of private repositories
func (el *DockerSystem) ImageBuildFromRemoteServer(
	server,
	imageName string,
	tags []string,
	channel *chan ContainerPullStatusSendToChannel,
) (
	err error,
) {

	var imageBuildOptions types.ImageBuildOptions
	var reader io.Reader

	if len(tags) == 0 {
		tags = []string{
			imageName,
		}
	} else {
		tags = append(tags, imageName)
	}

	imageBuildOptions = types.ImageBuildOptions{
		Tags:          tags,
		Remove:        true,
		RemoteContext: server,
	}

	reader, err = el.imageBuild(nil, imageBuildOptions)
	successfully := el.processBuildAndPullReaders(&reader, channel)
	if successfully == false && err == nil {
		err = errors.New("image build error")
	}

	return
}
