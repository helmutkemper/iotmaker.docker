package factoryDocker

import iotmakerDocker "github.com/helmutkemper/iotmaker.docker"

func NewClientWithHTTPHeaders(headers map[string]string) (err error, dockerSystem *iotmakerDocker.DockerSystem) {
	dockerSystem = &iotmakerDocker.DockerSystem{}
	dockerSystem.ContextCreate()
	err = dockerSystem.ClientCreateWithHTTPHeaders(headers)

	return
}
