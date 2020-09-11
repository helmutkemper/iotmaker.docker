package factorydocker

import (
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0.0-RC"
)

// en: Prepare docker system
//
//     Example:
//       err, dockerSys := factoryDocker.NewClient()
//       if err != nil {
//         panic(err)
//       }
//       dockerSys.ContainerCreateChangeExposedPortAndStart(...)
func NewClient() (err error, dockerSystem *iotmakerdocker.DockerSystem) {
	dockerSystem = &iotmakerdocker.DockerSystem{}
	dockerSystem.ContextCreate()
	err = dockerSystem.ClientCreate()

	return
}

//
