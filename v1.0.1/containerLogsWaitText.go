package iotmakerdocker

import (
	"github.com/docker/docker/api/types"
	"io"
	"io/ioutil"
	"strings"
)

// ContainerLogs (English): Returns container std out
//
// ContainerLogs (Português): Retorna a saída padrão do container
func (el *DockerSystem) ContainerLogsWaitText(
	id string,
	text string,
) (
	log []byte,
	err error,
) {

	var reader io.ReadCloser

	for {
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
		if err != nil {
			return
		}

		if strings.Contains(string(log), text) == true {
			return
		}
	}
}
