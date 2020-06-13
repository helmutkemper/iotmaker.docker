package factoryDocker

import iotmakerDocker "github.com/helmutkemper/iotmaker.docker"

// en: Prepare docker system
//
//     Example:
//       err, dockerSys := factoryDocker.NewClient()
//       if err != nil {
//         panic(err)
//       }
//       dockerSys.ContainerCreateChangeExposedPortAndStart(...)
func NewClient() (err error, dockerSystem *iotmakerDocker.DockerSystem) {
	dockerSystem = &iotmakerDocker.DockerSystem{}
	dockerSystem.ContextCreate()
	err = dockerSystem.ClientCreate()

	return
}
