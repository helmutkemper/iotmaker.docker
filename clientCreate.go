package iotmakerDocker

import (
	"github.com/docker/docker/client"
)

// English: Create a new client from docker
//
//     Example:
//       dockerSys = &iotmakerDocker.DockerSystem{}
//	     dockerSys.ContextCreate()
//	     err := dockerSys.ClientCreate()
//       if err != nil {
//         panic(err)
//       }
//
//     Please, use:
//       err, dockerSys := factoryDocker.NewClient()
//       if err != nil {
//         panic(err)
//       }
//       dockerSys.ContainerCreateChangeExposedPortAndStart(...)
//
// Português: Cria um novo cliente para o docker
//
//     Exemplo:
//       dockerSys = &iotmakerDocker.DockerSystem{}
//	     dockerSys.ContextCreate()
//	     err := dockerSys.ClientCreate()
//       if err != nil {
//         panic(err)
//       }
//
//     Por favor, use:
//       err, dockerSys := factoryDocker.NewClient()
//       if err != nil {
//         panic(err)
//       }
//       dockerSys.ContainerCreateChangeExposedPortAndStart(...)
//
func (el *DockerSystem) ClientCreate() error {
	var err error

	el.cli, err = client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)

	return err
}
