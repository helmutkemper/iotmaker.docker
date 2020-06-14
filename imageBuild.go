package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
	"io"
)

// en: image build from reader
//     please, see ImageBuildFromFolder(folderPath string, tags []string) and
//     ImageBuildFromRemoteServer(server string, tags []string)
func (el *DockerSystem) imageBuild(dockerFileTarReader io.Reader, imageBuildOptions types.ImageBuildOptions) (err error, reader io.ReadCloser) {
	var response types.ImageBuildResponse

	response, err = el.cli.ImageBuild(el.ctx, dockerFileTarReader, imageBuildOptions)
	if err != nil {
		return
	}

	reader = response.Body

	return
}
