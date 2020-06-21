package factoryDocker

import (
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

// create a network with gateway 10.0.0.1 and subnet 10.0.0.0/16
func NewNetwork(networkName string) (
	err error,
	networkId string,
	networkAutoConfiguration *iotmakerDocker.NextNetworkAutoConfiguration,
) {

	var dockerSys = iotmakerDocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	err, networkId, networkAutoConfiguration = dockerSys.NetworkCreate(
		networkName,
		iotmakerDocker.KNetworkDriveBridge,
		"local",
		"10.0.0.0/16",
		"10.0.0.1",
	)

	return
}
