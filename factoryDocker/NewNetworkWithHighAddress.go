package factoryDocker

import (
	"fmt"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

func NewNetworkWithHighAddress(networkName string, address, subnetMask byte) (err error, networkId string, networkAutoConfiguration *iotmakerDocker.NextNetworkAutoConfiguration) {
	var dockerSys = iotmakerDocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		return
	}

	err, networkId, networkAutoConfiguration = dockerSys.NetworkCreate(
		networkName,
		iotmakerDocker.KNetworkDriveBridge,
		"local",
		fmt.Sprintf("%d.0.0.0/%d", address, subnetMask),
		fmt.Sprintf("%d.0.0.1", address),
	)

	return
}
