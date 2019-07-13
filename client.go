package iotmaker_docker

import "github.com/docker/docker/client"

func GetClient() (*client.Client, error) {
	return client.NewClientWithOpts(client.WithVersion(kDockerVersion))
}
