package iotmaker_docker

import (
	"github.com/docker/docker/api/types"
	"io"
	"os"
)

// image pull
func (el *DockerSystem) ImagePull(name string, attachStdOut bool) error {
	reader, err := el.cli.ImagePull(el.ctx, name, types.ImagePullOptions{})
	if err != nil {
		return err
	}

	if len(el.imageId) == 0 {
		el.imageId = make(map[string]string)
	}

	el.imageId[name] = ""

	if attachStdOut == true {
		io.Copy(os.Stdout, reader)
	}

	return err
}
