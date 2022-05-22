package iotmakerdocker

import (
	"bytes"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/helmutkemper/iotmaker.docker/util"
	"io"
	"io/ioutil"
	"path/filepath"
	"time"
)

// FindDockerFile (English): Find dockerfile in folder tree.
//   Priority order: './Dockerfile-iotmaker', './Dockerfile', './dockerfile', 'Dockerfile.*', 'dockerfile.*',
//   '.*Dockerfile.*', '.*dockerfile.*'
//
// FindDockerFile (Português): Procura pelo arquivo dockerfile na árvore de diretórios.
//   Ordem de prioridade: './Dockerfile-iotmaker', './Dockerfile', './dockerfile', 'Dockerfile.*', 'dockerfile.*',
//   '.*Dockerfile.*', '.*dockerfile.*'
func (el *DockerSystem) FindDockerFile(folderPath string) (fullPathInsideTarFile string, err error) {
	var fileExists bool

	folderPath, err = filepath.Abs(folderPath)
	if err != nil {
		return
	}

	_, err = ioutil.ReadDir(folderPath)
	if err != nil {
		return
	}

	fileExists = util.VerifyFileExists(folderPath + "/Dockerfile-iotmaker")
	if fileExists == true {
		fullPathInsideTarFile = "/Dockerfile-iotmaker"
		return
	}

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

	fullPathInsideTarFile, err = util.FileFindHasPrefixRecursively("Dockerfile-iotmaker")
	if err == nil {
		return
	}

	fullPathInsideTarFile, err = util.FileFindHasPrefixRecursively("dockerfile-iotmaker")
	if err == nil {
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
//
//     Note: dockerfile priority order: './Dockerfile-iotmaker', './Dockerfile', './dockerfile', 'Dockerfile.*',
//     'dockerfile.*', '.*Dockerfile.*', '.*dockerfile.*'
//
// ImageBuildFromFolder (Português): Monta uma imagem a partir de um diretório
//   folderPath: string caminho absoluto do diretório
//   tags: []string tags da imagem
//   channel: *chan channel com dados do pull/build da imagem
//
//     Nota: ordem de prioridade do dockerfile: './Dockerfile-iotmaker', './Dockerfile', './dockerfile', 'Dockerfile.*',
//     'dockerfile.*', '.*Dockerfile.*', '.*dockerfile.*'
//
func (el *DockerSystem) ImageBuildFromFolder(
	folderPath string,
	imageName string,
	tags []string,
	imageBuildOptions types.ImageBuildOptions,
	channel *chan ContainerPullStatusSendToChannel,
) (
	imageID string,
	err error,
) {

	var tarFileReader *bytes.Reader
	var reader io.Reader
	var dockerFilePath string
	var dockerFileName string

	if len(tags) == 0 {
		tags = []string{
			imageName,
		}
	} else {
		tags = append(tags, imageName)
	}

	tarFileReader, err = el.ImageBuildPrepareFolderContext(folderPath)
	if err != err {
		return
	}

	dockerFilePath, err = el.FindDockerFile(folderPath)
	if err != nil {
		return
	}

	dockerFileName = filepath.Base(dockerFilePath)

	if len(imageBuildOptions.Tags) == 0 {
		imageBuildOptions.Tags = tags
	} else {
		imageBuildOptions.Tags = append(imageBuildOptions.Tags, tags...)
	}

	imageBuildOptions.Remove = true
	imageBuildOptions.Dockerfile = dockerFileName

	reader, err = el.ImageBuild(tarFileReader, imageBuildOptions)
	if err != nil {
		return
	}

	var successfully, processEnd bool
	var abort = make(chan struct{}, 1)
	var tk = time.NewTicker(1 * time.Second)
	go func(processEnd *bool, err *error) {
		successfully, *err = el.processBuildAndPullReaders(&reader, channel, abort)
		if successfully == false || *err != nil {
			if *err != nil {
				*processEnd = true
				return
			}

			*err = errors.New("image build error")
			*processEnd = true
			return
		}

		*processEnd = true
	}(&processEnd, &err)

	for {
		select {
		case <-tk.C:

			if err != nil {
				return
			}

			imageID, err = el.ImageFindIdByName(imageBuildOptions.Tags[0])
			if err != nil && err.Error() != "image name not found" {
				abort <- struct{}{}
				return
			}

			if imageID != "" {
				abort <- struct{}{}
				return
			}

			if processEnd == true && imageID == "" {
				imageID, err = el.ImageFindIdByName(imageBuildOptions.Tags[0])
				return
			}
		}
	}
}
