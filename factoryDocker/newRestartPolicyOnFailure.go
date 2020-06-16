package factoryDocker

import (
	whaleAquarium "github.com/helmutkemper/iotmaker.docker"
)

func NewRestartPolicyOnFailureRestart() whaleAquarium.RestartPolicy {
	return whaleAquarium.KRestartPolicyOnFailure
}
