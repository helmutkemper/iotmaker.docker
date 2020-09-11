package iotmakerdocker

import (
	"github.com/docker/docker/api/types"
	"io"
	"io/ioutil"
)

func (el *DockerSystem) ContainerLogs(
	id string,
) (
	log []byte,
	err error,
) {

	var reader io.ReadCloser

	reader, err = el.cli.ContainerLogs(el.ctx, id, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Timestamps: true,
		Follow:     false,
		Details:    false,
	})
	if err != nil {
		return
	}

	log, err = ioutil.ReadAll(reader)

	return
}
