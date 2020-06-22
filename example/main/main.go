package main

import (
	"fmt"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/v1.0.0/factoryContainerFromRemoteServer"
	"github.com/helmutkemper/iotmaker.docker/factoryDocker"
)

func main() {
	var err error
	var pullStatusChannel = factoryDocker.NewImagePullStatusChannel()
	var networkAutoConfiguration *iotmakerDocker.NextNetworkAutoConfiguration

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

	err, _, networkAutoConfiguration = factoryDocker.NewNetwork(
		"network_delete_before_test",
	)
	if err != nil {
		panic(err)
	}

	// address: 10.0.0.5:8080
	err, _, _, _ = factoryContainerFromRemoteServer.NewContainerFromRemoteServerWithNetworkConfiguration(
		"image_server_delete_before_test:latest",
		"cont_server_delete_before_test:2",
		iotmakerDocker.KRestartPolicyUnlessStopped,
		networkAutoConfiguration,
		"https://github.com/helmutkemper/lixo.git",
		[]string{},
		pullStatusChannel,
	)
	if err != nil {
		panic(err)
	}
}
