package iotmakerDocker

import "github.com/docker/docker/api/types"

func (el *DockerSystem) ContainerNetworkInspect(id string) (err error, netDataList ContainerNetworkDataList) {
	var insp types.ContainerJSON

	netDataList = make(map[string]ContainerNetworkData)

	err, insp = el.ContainerInspect(id)
	containerNetworks := (*insp.NetworkSettings).Networks
	for k, v := range containerNetworks {
		netDataList[k] = ContainerNetworkData{
			Gateway:    v.Gateway,
			IPAddress:  v.IPAddress,
			EndpointID: v.EndpointID,
			NetworkID:  v.NetworkID,
			MacAddress: v.MacAddress,
		}
	}

	return
}
