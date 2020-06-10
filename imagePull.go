package iotmaker_docker

import (
	"encoding/json"
	"github.com/docker/docker/api/types"
	"io"
	"strings"
)

func (el *DockerSystem) imagePullWriteChannel(progressChannel *chan ContainerPullStatusSendToChannel, data ContainerPullStatusSendToChannel) {
	if progressChannel == nil {
		return
	}

	l := len(*progressChannel)
	if l != 0 {
		return
	}

	*progressChannel <- data
}

func (el *DockerSystem) ImagePull(name string, channel chan ContainerPullStatusSendToChannel) (err error, imageId string, imageName string) {
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

	var bufferReader = make([]byte, 1)
	var bufferDataInput = make([]byte, 0)

	var channelOut ContainerPullProgress
	var toChannel ContainerPullStatusSendToChannel
	var toProcess = make(map[string]ContainerPullProgress)

	for {
		_, err = reader.Read(bufferReader)
		if err != nil {
			if err.Error() == "EOF" {
				err = nil

				//>>>>> send to channel
				toChannel.ImageName = imageName
				toChannel.ImageID = imageId
				toChannel.Closed = true
				el.imagePullWriteChannel(&channel, toChannel)

				return
			}
			return
		}

		bufferDataInput = append(bufferDataInput, bufferReader[0])

		if bufferReader[0] == byte(0x0A) {
			err = json.Unmarshal(bufferDataInput, &channelOut)
			bufferDataInput = make([]byte, 0)

			if strings.Contains(channelOut.Status, KContainerPullStatusDownloadedNewerImageText) {
				imageName = strings.Replace(channelOut.Status, KContainerPullStatusDownloadedNewerImageText, "", 1)
			}

			if strings.Contains(channelOut.Status, KContainerPullStatusDigestText) {
				imageId = strings.Replace(channelOut.Status, KContainerPullStatusDigestText, "", 1)
			}

			if strings.Contains(channelOut.Status, KContainerPullStatusPullCompleteText) {
				channelOut.SysStatus = KContainerPullStatusPullComplete
			}

			if strings.Contains(channelOut.Status, KContainerPullStatusExtractingText) {
				channelOut.SysStatus = KContainerPullStatusExtracting
			}

			if strings.Contains(channelOut.Status, KContainerPullStatusWaitingText) {
				channelOut.SysStatus = KContainerPullStatusWaiting
			}

			if strings.Contains(channelOut.Status, KContainerPullStatusDownloadingText) {
				channelOut.SysStatus = KContainerPullStatusDownloading
			}

			if strings.Contains(channelOut.Status, KContainerPullStatusVerifyingChecksumText) {
				channelOut.SysStatus = KContainerPullStatusVerifyingChecksum
			}

			if strings.Contains(channelOut.Status, KContainerPullStatusDownloadCompleteText) {
				channelOut.SysStatus = KContainerPullStatusDownloadComplete
			}

			if strings.Contains(channelOut.Status, KContainerPullStatusImageIsUpToDate) {
				imageName = strings.Replace(channelOut.Status, KContainerPullStatusImageIsUpToDate, "", 1)
			}

			toProcess[channelOut.ID] = channelOut

			for _, v := range toProcess {
				if v.SysStatus == KContainerPullStatusPullComplete {
					toChannel.PullComplete += 1
				} else if v.SysStatus == KContainerPullStatusExtracting {
					toChannel.Extracting.Count += 1
					toChannel.Extracting.Total += v.ProgressDetail.Total
					toChannel.Extracting.Current += v.ProgressDetail.Current
				} else if v.SysStatus == KContainerPullStatusWaiting {
					toChannel.Waiting += 1
				} else if v.SysStatus == KContainerPullStatusDownloading {
					toChannel.Downloading.Count += 1
					toChannel.Downloading.Total += v.ProgressDetail.Total
					toChannel.Downloading.Current += v.ProgressDetail.Current
				} else if v.SysStatus == KContainerPullStatusVerifyingChecksum {
					toChannel.VerifyingChecksum += 1
				} else if v.SysStatus == KContainerPullStatusDownloadComplete {
					toChannel.DownloadComplete += 1
				}
			}

			toChannel.ImageName = imageName
			toChannel.ImageID = imageId

			//>>>>> send to channel
			el.imagePullWriteChannel(&channel, toChannel)

			toChannel = ContainerPullStatusSendToChannel{}
		}
	}
}

type ContainerPullStatusSendToChannelCount struct {
	Count   int
	Current int
	Total   int
}

type ContainerPullStatusSendToChannel struct {
	Waiting           int
	Downloading       ContainerPullStatusSendToChannelCount
	VerifyingChecksum int
	DownloadComplete  int
	Extracting        ContainerPullStatusSendToChannelCount
	PullComplete      int
	ImageName         string
	ImageID           string
	Closed            bool
}

type ContainerPullProgressDetail struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}

// Complete list of states:
//  {"status":"Pulling from library/mongo","id":"latest"}
//  {"status":"Pulling fs layer","progressDetail":{},"id":"23884877105a"}
//  {"status":"Waiting","progressDetail":{},"id":"358ed78d3204"}
//  {"status":"Downloading","progressDetail":{"current":422,"total":35367},"progress":"[\u003e          ]     422B/35.37kB","id":"bc38caa0f5b9"}
//  {"status":"Verifying Checksum","progressDetail":{},"id":"bc38caa0f5b9"}
//  {"status":"Download complete","progressDetail":{},"id":"bc38caa0f5b9"}
//  {"status":"Extracting","progressDetail":{"current":294912,"total":26689802},"progress":"[\u003e     ]  294.9kB/26.69MB","id":"23884877105a"}
//  {"status":"Pull complete","progressDetail":{},"id":"23884877105a"}
//  {"status":"Digest: sha256:be8d903a68997dd63f64479004a7eeb4f0674dde7ab3cbd1145e5658da3a817b"}
//  {"status":"Status: Downloaded newer image for mongo:latest"}
//  {"status":"Status: Status: Image is up to date for mongo:latest"}
type ContainerPullProgress struct {
	Status         string                      `json:"status"`
	ProgressDetail ContainerPullProgressDetail `json:"progressDetail"`
	ID             string                      `json:"id"`
	SysStatus      ContainerPullStatus         `json:"-"`
	ImageName      string
}

type ContainerPullStatus int

const (
	KContainerPullStatusWaiting ContainerPullStatus = iota + 1
	KContainerPullStatusDownloading
	KContainerPullStatusVerifyingChecksum
	KContainerPullStatusDownloadComplete
	KContainerPullStatusExtracting
	KContainerPullStatusPullComplete
)

const (
	KContainerPullStatusWaitingText              = "Waiting"
	KContainerPullStatusDownloadingText          = "Downloading"
	KContainerPullStatusVerifyingChecksumText    = "Verifying Checksum"
	KContainerPullStatusDownloadCompleteText     = "Download complete"
	KContainerPullStatusExtractingText           = "Extracting"
	KContainerPullStatusPullCompleteText         = "Pull complete"
	KContainerPullStatusDigestText               = "Digest: "
	KContainerPullStatusDownloadedNewerImageText = "Status: Downloaded newer image for "
	KContainerPullStatusImageIsUpToDate          = "Status: Image is up to date for "
)
