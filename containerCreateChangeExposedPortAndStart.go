package iotmakerDocker

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
)

func (el *DockerSystem) ContainerCreateChangeExposedPortAndStart(imageName, containerName string, restart RestartPolicy, mountVolumes []mount.Mount, net *network.NetworkingConfig, currentPort, changeToPort []nat.Port) (error, string) {
	err, id := el.ContainerCreateAndChangeExposedPort(imageName, containerName, restart, mountVolumes, net, currentPort, changeToPort)
	if err != nil {
		return err, ""
	}

	err = el.ContainerStart(id)
	return err, id
}
