package iotmakerDocker

import (
	"encoding/json"
	"io"
	"strings"
)

func (el *DockerSystem) processBuildAndPullReaders(reader *io.Reader, channel *chan ContainerPullStatusSendToChannel) {
	var err error
	var imageName string
	var imageId string
	var bufferReader = make([]byte, 1)
	var bufferDataInput = make([]byte, 0)
	var channelOut ContainerPullProgress
	var toChannel ContainerPullStatusSendToChannel
	var toProcess = make(map[string]ContainerPullProgress)

	if reader == nil {
		return
	}

	if *reader == nil {
		return
	}

	for {
		_, err = (*reader).Read(bufferReader)
		if err != nil {
			if err.Error() == "EOF" {
				err = nil

				//>>>>> send to channel
				toChannel.calcPercentage()
				toChannel.ImageName = imageName
				toChannel.ImageID = imageId
				toChannel.Closed = true
				el.imagePullWriteChannel(channel, toChannel)

				return
			}
			return
		}

		bufferDataInput = append(bufferDataInput, bufferReader[0])

		if bufferReader[0] == byte(0x0A) {
			err = json.Unmarshal(bufferDataInput, &channelOut)
			bufferDataInput = make([]byte, 0)

			if strings.Contains(channelOut.Stream, kContainerBuildImageStatusSuccess) {
				channelOut.SysStatus = KContainerPullStatusComplete
			} else if channelOut.Stream != "" {
				channelOut.SysStatus = KContainerPullStatusBuilding
			}

			if strings.Contains(channelOut.Status, kContainerPullStatusDownloadedNewerImageText) {
				imageName = strings.Replace(channelOut.Status, kContainerPullStatusDownloadedNewerImageText, "", 1)
			}

			if strings.Contains(channelOut.Status, kContainerPullStatusDigestText) {
				imageId = strings.Replace(channelOut.Status, kContainerPullStatusDigestText, "", 1)
			}

			if strings.Contains(channelOut.Status, kContainerPullStatusPullCompleteText) {
				channelOut.SysStatus = KContainerPullStatusPullComplete
			}

			if strings.Contains(channelOut.Status, kContainerPullStatusExtractingText) {
				channelOut.SysStatus = KContainerPullStatusExtracting
			}

			if strings.Contains(channelOut.Status, kContainerPullStatusWaitingText) {
				channelOut.SysStatus = KContainerPullStatusWaiting
			}

			if strings.Contains(channelOut.Status, kContainerPullStatusDownloadingText) {
				channelOut.SysStatus = KContainerPullStatusDownloading
			}

			if strings.Contains(channelOut.Status, kContainerPullStatusVerifyingChecksumText) {
				channelOut.SysStatus = KContainerPullStatusVerifyingChecksum
			}

			if strings.Contains(channelOut.Status, kContainerPullStatusDownloadCompleteText) {
				channelOut.SysStatus = KContainerPullStatusDownloadComplete
			}

			if strings.Contains(channelOut.Status, kContainerPullStatusImageIsUpToDate) {
				imageName = strings.Replace(channelOut.Status, kContainerPullStatusImageIsUpToDate, "", 1)
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

			toChannel.calcPercentage()
			toChannel.ImageName = imageName
			toChannel.ImageID = imageId
			toChannel.Stream = channelOut.Stream

			//>>>>> send to channel
			el.imagePullWriteChannel(channel, toChannel)

			toChannel = ContainerPullStatusSendToChannel{}
		}
	}
}
