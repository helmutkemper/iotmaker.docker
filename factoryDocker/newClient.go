package factoryDocker

import iotmakerDocker "github.com/helmutkemper/iotmaker.docker"

func NewClient() (err error, dockerSystem *iotmakerDocker.DockerSystem) {
	dockerSystem = &iotmakerDocker.DockerSystem{}
	dockerSystem.ContextCreate()
	err = dockerSystem.ClientCreate()

	return
}
