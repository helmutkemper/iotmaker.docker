package factoryDocker

import (
	"fmt"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func NewPullStatusMonitor() (pullStatusChannel *chan iotmakerDocker.ContainerPullStatusSendToChannel) {
	pullStatusChannel = NewImagePullStatusChannel()

	go func(c chan iotmakerDocker.ContainerPullStatusSendToChannel) {

		for {
			select {
			case status := <-c:
				fmt.Printf("image pull status: %+v\n", status)

				if status.Closed == true {
					fmt.Println("image pull complete!")
				}
			}
		}

	}(*pullStatusChannel)

	return
}
