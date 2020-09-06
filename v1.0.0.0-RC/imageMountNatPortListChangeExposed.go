package iotmakerDocker

import (
	"github.com/docker/go-connections/nat"
)

// Mount nat por list by image config
func (el *DockerSystem) ImageMountNatPortListChangeExposed(
	imageId string,
	currentPortList,
	changeToPortList []nat.Port,
) (
	nat.PortMap,
	error,
) {

	return el.ImageMountNatPortListChangeExposedWithIpAddress(imageId, "", currentPortList, changeToPortList)
}
