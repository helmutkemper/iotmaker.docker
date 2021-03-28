package iotmakerdocker

import (
	"bytes"
	"github.com/docker/docker/api/types"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

// ContainerLogs (English):
//
// ContainerLogs (Português):
func (el *DockerSystem) ContainerLogsWaitText(
	id string,
	text string,
	out io.Writer,
) (
	logContainer []byte,
	err error,
) {

	var reader io.ReadCloser
	var previousLog = make([]byte, 0)
	var cleanLog = make([]byte, 0)

	if out != nil {
		log.New(out, "", 0)
	}

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

		logContainer, err = ioutil.ReadAll(reader)
		if err != nil {
			return
		}

		cleanLog = bytes.Replace(logContainer, previousLog, []byte(""), -1)
		previousLog = make([]byte, len(logContainer))
		copy(previousLog, logContainer)

		if out != nil {
			log.Printf("%s", cleanLog)
		}

		if strings.Contains(string(logContainer), text) == true {
			return
		}
	}
}
