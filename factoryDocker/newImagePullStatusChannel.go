package factorydocker

import iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0.0-RC"

func NewImagePullStatusChannel() (chanPointer *chan iotmakerdocker.ContainerPullStatusSendToChannel) {
	channel := make(chan iotmakerdocker.ContainerPullStatusSendToChannel, 1)

	chanPointer = &channel
	return
}
