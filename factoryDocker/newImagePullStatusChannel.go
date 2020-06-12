package factoryDocker

import iotmakerDocker "github.com/helmutkemper/iotmaker.docker"

func NewImagePullStatusChannel() chan iotmakerDocker.ContainerPullStatusSendToChannel {
	return make(chan iotmakerDocker.ContainerPullStatusSendToChannel, 1)
}
