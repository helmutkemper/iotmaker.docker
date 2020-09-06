package iotmakerDocker

import (
	"bytes"
	"errors"
	"github.com/docker/docker/api/types"
	"io"
)

// English: Make a image from folder path content
//     Please note: dockerfile name must be "Dockerfile" inside root folder
func (el *DockerSystem) ImageBuildFromFolder(
	folderPath string,
	tags []string,
	channel *chan ContainerPullStatusSendToChannel,
) (
	err error,
) {

	var tarFileReader *bytes.Reader
	var imageBuildOptions types.ImageBuildOptions
	var reader io.Reader

	tarFileReader, err = el.imageBuildPrepareFolderContext(folderPath)
	if err != err {
		return
	}

	imageBuildOptions = types.ImageBuildOptions{
		Tags:   tags,
		Remove: true,
	}

	reader, err = el.imageBuild(tarFileReader, imageBuildOptions)
	successfully := el.processBuildAndPullReaders(&reader, channel)
	if successfully == false && err == nil {
		err = errors.New("image build error")
	}

	return
}
