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
	var successfully bool
	var abort = make(chan struct{}, 1)
	var tk = time.NewTicker(1 * time.Second)
	var errReaders error
	go func() {
		successfully, errReaders = el.processBuildAndPullReaders(&reader, channel, abort)
		if successfully == false || err != nil {
			if err != nil {
				return
			}

			err = errors.New("image pull error")
		}
	}()

	for {
		select {
		case <-tk.C:
			if errReaders != nil {
				abort <- struct{}{}
				err = errReaders
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
		}
	}

}
