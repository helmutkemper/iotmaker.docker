package iotmakerDocker

import "github.com/docker/docker/api/types"

func (el *DockerSystem) ContainerRemoveWithLinksAndVolumes(id string) error {
	return el.cli.ContainerRemove(el.ctx, id, types.ContainerRemoveOptions{RemoveLinks: true, RemoveVolumes: true})
}

func (el *DockerSystem) ContainerBuild(dockerFile string) (err error, response types.ImageBuildResponse) {
	response, err = el.cli.ImageBuild(el.ctx, nil, types.ImageBuildOptions{
		Tags: []string{"container:test"},
		Labels: map[string]string{
			"version": "0",
		},
		Version:    "0",
		Dockerfile: dockerFile,
	})

	return
}

//git@github.com:helmutkemper/lixo.git
