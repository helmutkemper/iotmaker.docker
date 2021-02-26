package iotmakerdocker

import (
	"bytes"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/iotmaker.docker/util"
	"io"
)

// FindDockerFile (English): Find dockerfile in folder tree.
//   Priority order: './Dockerfile', './dockerfile', 'Dockerfile.*', 'dockerfile.*',
//   '.*Dockerfile.*', '.*dockerfile.*'
//
// FindDockerFile (Português): Procura pelo arquivo dockerfile na árvore de diretórios.
//   Ordem de prioridade: './Dockerfile', './dockerfile', 'Dockerfile.*', 'dockerfile.*',
//   '.*Dockerfile.*', '.*dockerfile.*'
func (el *DockerSystem) FindDockerFile(folderPath string) (fullPathInsideTarFile string, err error) {
	var fileExists bool

	fileExists = util.VerifyFileExists(folderPath + "/Dockerfile")
	if fileExists == true {
		fullPathInsideTarFile = "/Dockerfile"
		return
	}

	fileExists = util.VerifyFileExists(folderPath + "/dockerfile")
	if fileExists == true {
		fullPathInsideTarFile = "/dockerfile"
		return
	}

	fullPathInsideTarFile, err = util.FileFindHasPrefixRecursively("Dockerfile")
	if err == nil {
		return
	}

	fullPathInsideTarFile, err = util.FileFindHasPrefixRecursively("dockerfile")
	if err == nil {
		return
	}

	fullPathInsideTarFile, err = util.FileFindContainsRecursively("Dockerfile")
	if err == nil {
		return
	}

	fullPathInsideTarFile, err = util.FileFindContainsRecursively("dockerfile")

	return
}

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
	var dockerFilePath string

	tarFileReader, err = el.ImageBuildPrepareFolderContext(folderPath)
	if err != err {
		return
	}

	dockerFilePath, err = el.FindDockerFile(folderPath)
	if err != nil {
		return
	}

	_ = dockerFilePath
	imageBuildOptions = types.ImageBuildOptions{
		Tags:       tags,
		Remove:     true,
		Dockerfile: "dockerfile",
	}

	reader, err = el.ImageBuild(tarFileReader, imageBuildOptions)
	if err != nil {
		return
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
