package iotmakerDocker

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
)

func (el *DockerSystem) ContainerCreateWithoutExposePortsAndStart(
	imageName,
	containerName string,
	restart RestartPolicy,
	mountVolumes []mount.Mount,
	net *network.NetworkingConfig,
) (err error, containerID string) {

	err, containerID = el.ContainerCreateWithoutExposePorts(imageName, containerName, restart, mountVolumes, net)
	if err != nil {
		return
	}

	err = el.ContainerStart(containerID)

	return
}
