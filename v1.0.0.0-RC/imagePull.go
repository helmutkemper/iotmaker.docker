package iotmakerDocker

import (
	"errors"
	"github.com/docker/docker/api/types"
	"io"
)

const (
	kContainerPullStatusWaitingText              = "Waiting"
	kContainerPullStatusDownloadingText          = "Downloading"
	kContainerPullStatusVerifyingChecksumText    = "Verifying Checksum"
	kContainerPullStatusDownloadCompleteText     = "Download complete"
	kContainerPullStatusExtractingText           = "Extracting"
	kContainerPullStatusPullCompleteText         = "Pull complete"
	kContainerPullStatusAuxId                    = "\"aux\":{\"ID\""
	kContainerPullStatusDigestText               = "Digest: "
	kContainerPullStatusDownloadedNewerImageText = "Status: Downloaded newer image for "
	kContainerPullStatusImageIsUpToDate          = "Status: Image is up to date for "
	kContainerBuildImageStatusSuccessContainer   = "Success fully tagged"
	kContainerBuildImageStatusSuccessImage       = "Successfully tagged"
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
	successfully := el.processBuildAndPullReaders(&reader, channel)
	if successfully == false {
		err = errors.New("image pull error")
	}

	return
}
