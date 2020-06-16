package factoryDocker

import (
	whaleAquarium "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker/util"
)

func NewContainerNetworkGenerator(name string, a, b, c, d byte) (err error, netGenerator util.NextNetworkConfiguration, netId string) {
	var exists bool
	var net = whaleAquarium.DockerSystem{}
	net.Init()

  err, exists = net.NetworkVerifyName(name)
	if err != nil {
	  err, netId = net.NetworkFindIdByName(name)
    netGenerator.Init(netId, name, a, b, c, d)
		return
	}

	if exists == false {
		err, netId = net.NetworkCreate(name)
    netGenerator.Init(netId, name, a, b, c, d)
		if err != nil {
			return
		}
	}

	return
}
