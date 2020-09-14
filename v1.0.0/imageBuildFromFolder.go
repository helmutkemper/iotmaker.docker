package iotmakerdocker

import (
	"bytes"
	"errors"
	"github.com/docker/docker/api/types"
	"io"
)

// ImageBuildFromFolder (English): Make a image from folder path content
//   folderPath: string absolute folder path
//   tags: []string image tags
//   channel: *chan channel of pull/build data
// Note: dockerfile name must be "Dockerfile" inside root folder
//
// ImageBuildFromFolder (Português): Monta uma imagem a partir de um diretório
//   folderPath: string caminho absoluto do diretório
//   tags: []string tags da imagem
//   channel: *chan channel com dados do pull/build da imagem
// Nota: O nome do arquivo dockerfile dentro da raiz do diretório deve ser "Dockerfile"
func (el *DockerSystem) ImageBuildFromFolder(
	folderPath string,
	tags []string,
	channel *chan ContainerPullStatusSendToChannel,
) (
	imageID string,
	err error,
) {

	var tarFileReader *bytes.Reader
	var imageBuildOptions types.ImageBuildOptions
	var reader io.Reader

	tarFileReader, err = el.ImageBuildPrepareFolderContext(folderPath)
	if err != err {
		return
	}

	imageBuildOptions = types.ImageBuildOptions{
		Tags:   tags,
		Remove: true,
	}

	reader, err = el.ImageBuild(tarFileReader, imageBuildOptions)
	if err != nil {
		panic(err)
	}

	successfully := el.processBuildAndPullReaders(&reader, channel)
	if successfully == false {
		err = errors.New("image build error")
		return
	}

	imageID, err = el.ImageFindIdByName(imageBuildOptions.Tags[0])
	if err != nil {
		return
	}

	return
}
