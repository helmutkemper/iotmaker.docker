package iotmakerDocker

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
)

func (el *DockerSystem) ContainerCreateChangeExposedPortAndStart(
	imageName,
	containerName string,
	restartPolicy RestartPolicy,
	mountVolumes []mount.Mount,
	containerNetwork *network.NetworkingConfig,
	currentPort,
	changeToPort []nat.Port,
) (err error, containerID string) {

	err, containerID = el.ContainerCreateAndChangeExposedPort(
		imageName,
		containerName,
		restartPolicy,
		mountVolumes,
		containerNetwork,
		currentPort,
		changeToPort,
	)
	if err != nil {
		return
	}

	err = el.ContainerStart(containerID)
	return
}
