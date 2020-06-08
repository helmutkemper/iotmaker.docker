package iotmaker_docker

import "strings"

func ContainerGetLasNameElement(name string) string {
	names := strings.Split(name, "/")

	l := len(names) - 1

	if l > -1 {
		return names[l]
	}

	return name
}
