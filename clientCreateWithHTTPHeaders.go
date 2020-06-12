package iotmakerDocker

import (
	"github.com/docker/docker/client"
)

func (el *DockerSystem) ClientCreateWithHTTPHeaders(headers map[string]string) error {
	var err error

	el.cli, err = client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
		client.WithHTTPHeaders(headers),
	)

	return err
}
