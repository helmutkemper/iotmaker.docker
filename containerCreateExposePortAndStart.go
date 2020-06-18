package iotmakerDocker

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
)

func (el *DockerSystem) ContainerCreateExposePortAndStart(
	imageName,
	containerName string,
	restart RestartPolicy,
	mountVolumes []mount.Mount,
	net *network.NetworkingConfig,
	portExposedList nat.PortMap,
) (error, string) {

	var err error
	var resp container.ContainerCreateCreatedBody

	if len(el.container) == 0 {
		el.container = make(map[string]container.ContainerCreateCreatedBody)
	}

	resp, err = el.cli.ContainerCreate(
		el.ctx,
		&container.Config{
			Image: imageName,
		},
		&container.HostConfig{
			PortBindings: portExposedList,
			RestartPolicy: container.RestartPolicy{
				Name: restart.String(),
			},
			Resources: container.Resources{},
			Mounts:    mountVolumes,
		},
		net,
		containerName,
	)
	if err != nil {
		return err, ""
	}

	el.container[resp.ID] = resp

	return nil, resp.ID
}
