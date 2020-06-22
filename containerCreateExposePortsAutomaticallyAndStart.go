package iotmakerDocker

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
)

func (el *DockerSystem) ContainerCreateExposePortsAutomaticallyAndStart(
	imageName,
	containerName string,
	restartPolicy RestartPolicy,
	mountVolumes []mount.Mount,
	containerNetwork *network.NetworkingConfig,
) (err error, containerID string) {

	err, containerID = el.ContainerCreateAndExposePortsAutomatically(
		imageName,
		containerName,
		restartPolicy,
		mountVolumes,
		containerNetwork,
	)
	if err != nil {
		return
	}

	err = el.ContainerStart(containerID)
	return
}
