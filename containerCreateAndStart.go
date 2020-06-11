package iotmakerDocker

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
)

func (el *DockerSystem) ContainerCreateAndStart(imageName, containerName string, restart RestartPolicy, mountVolumes []mount.Mount, net *network.NetworkingConfig) (error, string) {
	err, id := el.ContainerCreate(imageName, containerName, restart, mountVolumes, net)
	if err != nil {
		return err, ""
	}

	err = el.ContainerStart(id)
	return err, id
}
