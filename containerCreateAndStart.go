package iotmakerDocker

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
)

// en: Create a container and automatically exposes the ports listed in the image before
//     start
//   imageName: image name for download and pull
//   containerName: unique container name
//   RestartPolicy:
//      KRestartPolicyNo - Do not automatically restart the container. (the
//          default)
//      KRestartPolicyOnFailure - Restart the container if it exits due to an
//          error, which manifests as a non-zero exit code.
//      KRestartPolicyAlways - Always restart the container if it stops. If it is
//          manually stopped, it is restarted only when Docker daemon restarts or
//          the container itself is manually restarted. (See the second bullet
//          listed in restart policy details)
//      KRestartPolicyUnlessStopped - Similar to always, except that when the
//          container is stopped (manually or otherwise), it is not restarted
//          even after Docker daemon restarts.
//   mountVolumes: please use a factoryWhaleAquarium.NewVolumeMount()
//      for a complete list of volumes exposed by image, use
//      ImageListExposedVolumes(id) and ImageListExposedVolumesByName(name)
//
//     Example:
//       err, dockerSys := factoryDocker.NewClient()
//       if err != nil {
//         panic(err)
//       }
//       err, id := dockerSys.ContainerCreateAndStart(
//       		"image:latest",
//       		"server",
//       		dockerSys.KRestartPolicyUnlessStopped,
//       		[]mount.Mount{},
//       		nil,
//       )
//       if err != nil {
//         panic(err)
//       }
func (el *DockerSystem) ContainerCreateAndStart(
	imageName,
	containerName string,
	restart RestartPolicy,
	portExposedList nat.PortMap,
	mountVolumes []mount.Mount,
	containerNetwork *network.NetworkingConfig,
) (err error, containerID string) {

	imageName = el.AdjustImageName(imageName)

	err, containerID = el.ContainerCreate(
		imageName,
		containerName,
		restart,
		portExposedList,
		mountVolumes,
		containerNetwork,
	)
	if err != nil {
		return
	}

	err = el.ContainerStart(containerID)
	return
}
