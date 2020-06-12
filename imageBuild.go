package iotmakerDocker

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/go-connections/nat"
	"io"
)

func (el *DockerSystem) ImageBuild(dockerFileTarReader io.Reader, t types.ImageBuildOptions) (err error, response types.ImageBuildResponse) {
	response, err = el.cli.ImageBuild(el.ctx, dockerFileTarReader, t)

	return
}

func (el *DockerSystem) ContainerBuild() (err error, response container.ContainerCreateCreatedBody) {
	var id string
	err, id = el.ContainerCreateChangeExposedPortAndStart(
		"kemper:latest",
		"server",
		KRestartPolicyUnlessStopped,
		[]mount.Mount{},
		nil,
		[]nat.Port{
			"tcp/8080",
		},
		[]nat.Port{
			"tcp/8080",
		},
	)
	fmt.Printf("container id: %v\n", id)

	return
}

//git@github.com:helmutkemper/lixo.git
