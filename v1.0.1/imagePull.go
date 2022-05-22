package iotmakerdocker

import (
	"errors"
	"github.com/docker/docker/api/types"
	"io"
	"time"
)

func (el *DockerSystem) ImagePull(
	name string,
	channel *chan ContainerPullStatusSendToChannel,
) (
	imageId string,
	imageName string,
	err error,
) {

	var reader io.Reader

	//esse valor é trocado no final do download
	imageName = name

	reader, err = el.cli.ImagePull(el.ctx, name, types.ImagePullOptions{})
	if err != nil {
		return
	}

	if len(el.imageId) == 0 {
		el.imageId = make(map[string]string)
	}

	el.imageId[name] = ""
	var successfully, processEnd bool
	var abort = make(chan struct{}, 1)
	var tk = time.NewTicker(1 * time.Second)
	go func(processEnd *bool, err *error) {
		successfully, *err = el.processBuildAndPullReaders(&reader, channel, abort)
		if successfully == false || *err != nil {
			if *err != nil {
				*processEnd = true
				return
			}

			*err = errors.New("image pull error")
		}

		*processEnd = true
	}(&processEnd, &err)

	for {
		select {
		case <-tk.C:
			if err != nil {
				abort <- struct{}{}
				return
			}

			imageId, err = el.ImageFindIdByName(name)
			if err != nil && err.Error() != "image name not found" {
				abort <- struct{}{}
				return
			}

			if imageId != "" {
				abort <- struct{}{}
				return
			}

			if processEnd == true && imageId == "" {
				imageId, err = el.ImageFindIdByName(name)
				return
			}
		}
	}
}
