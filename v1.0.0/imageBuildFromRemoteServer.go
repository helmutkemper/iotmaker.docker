package iotmakerdocker

import (
	"errors"
	"github.com/docker/docker/api/types"
	"io"
)

// ImageBuildFromRemoteServer (English): Make a image from server content
//   server: Server path.
//     Example: https://[<token>@]github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git
//   imageName: image name. Example: server:lasted
//   tags: image tags
//   channel: channel of pull/build data
//
//   Note: For get a github token
//   settings > Developer settings > Personal access tokens > Generate new token
//   Mark [x]repo - Full control of private repositories
//
// ImageBuildFromRemoteServer (Português): Prepara uma imagem a partir do conteúdo de um servidor
//   server: Caminho do arquivo.
//     Exemplo: https://[<token>@]github.com/helmutkemper/iotmaker.docker.util.whaleAquarium.sample.git
//   imageName: nome da imagem. Exemplo: server:lasted
//   tags: tags da imagem
//   channel: canal com dados de pull/build da imagem
//
//   Note: Para usar um token do github
//   settings > Developer settings > Personal access tokens > Generate new token
//   Mark [x]repo - Full control of private repositories
func (el *DockerSystem) ImageBuildFromRemoteServer(
	server,
	imageName string,
	tags []string,
	channel *chan ContainerPullStatusSendToChannel,
) (
	imageID string,
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

	reader, err = el.ImageBuild(nil, imageBuildOptions)
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
