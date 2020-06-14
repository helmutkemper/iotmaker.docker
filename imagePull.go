package iotmakerDocker

import (
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
	kContainerPullStatusDigestText               = "Digest: "
	kContainerPullStatusDownloadedNewerImageText = "Status: Downloaded newer image for "
	kContainerPullStatusImageIsUpToDate          = "Status: Image is up to date for "
	kContainerBuildImageStatusSuccess            = "Success fully tagged"
)

func (el *DockerSystem) ImagePull(name string, channel *chan ContainerPullStatusSendToChannel) (err error, imageId string, imageName string) {
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
	el.processBuildAndPullReaders(&reader, channel)

	return
}
