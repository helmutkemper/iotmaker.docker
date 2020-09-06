package iotmakerDocker

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
)

//
// English: Creates a container and automatically exposes the ports listed in the image
//   imageName: image name for download and pull
//   containerName: unique container name
//   RestartPolicy:
//      KRestartPolicyNo - Do not automatically restart the container. (the default)
//      KRestartPolicyOnFailure - Restart the container if it exits due to an error, which manifests as a non-zero exit code.
//      KRestartPolicyAlways - Always restart the container if it stops. If it is manually stopped, it is restarted only when Docker daemon restarts or the container itself is manually restarted. (See the second bullet listed in restart policy details)
//      KRestartPolicyUnlessStopped - Similar to always, except that when the container is stopped (manually or otherwise), it is not restarted even after Docker daemon restarts.
//   mountVolumes: please use a factoryWhaleAquarium.NewVolumeMount()
//      for a complete list of volumes exposed by image, use ImageListExposedVolumes(id) and ImageListExposedVolumesByName(name)
func (el *DockerSystem) ContainerCreateAndExposePortsAutomatically(
	imageName,
	containerName string,
	restartPolicy RestartPolicy,
	mountVolumes []mount.Mount,
	containerNetwork *network.NetworkingConfig,
) (
	containerID string,
	err error,
) {

	var imageId string
	var portExposedList nat.PortMap

	imageName = el.AdjustImageName(imageName)

	imageId, err = el.ImageFindIdByName(imageName)
	if err != nil {
		return "", err
	}

	portExposedList, err = el.ImageMountNatPortList(imageId)
	if err != nil {
		return "", err
	}

	if len(el.container) == 0 {
		el.container = make(map[string]container.ContainerCreateCreatedBody)
	}

	return el.ContainerCreate(
		imageName,
		containerName,
		restartPolicy,
		portExposedList,
		mountVolumes,
		containerNetwork,
	)
}