package iotmakerDocker

import "strings"

func (el DockerSystem) AdjustImageName(
	imageName string,
) string {

	if strings.Contains(imageName, ":") == false {
		imageName = imageName + ":latest"
	}

	return imageName
}
