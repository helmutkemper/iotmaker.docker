package factoryDocker

import iotmakerDocker "github.com/helmutkemper/iotmaker.docker"

func NewImagePullStatusChannel() (chanPointer *chan iotmakerDocker.ContainerPullStatusSendToChannel) {
	channel := make(chan iotmakerDocker.ContainerPullStatusSendToChannel, 1)

	chanPointer = &channel
	return
}
