package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
	"io"
	"os"
)

// en: image build from reader
//     please, see ImageBuildFromFolder(folderPath string, tags []string) and
//     ImageBuildFromRemoteServer(server string, tags []string)
func (el *DockerSystem) imageBuild(dockerFileTarReader io.Reader, imageBuildOptions types.ImageBuildOptions) (err error) {
	var response types.ImageBuildResponse

	response, err = el.cli.ImageBuild(el.ctx, dockerFileTarReader, imageBuildOptions)

	defer response.Body.Close()
	_, err = io.Copy(os.Stdout, response.Body)

	return
}
