package factoryDocker

import (
	whaleAquarium "github.com/helmutkemper/iotmaker.docker"
)

func NewRestartPolicyUnlessStopped() whaleAquarium.RestartPolicy {
	return whaleAquarium.KRestartPolicyUnlessStopped
}
