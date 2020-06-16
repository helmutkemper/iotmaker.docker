package factoryDocker

import whaleAquarium "github.com/helmutkemper/iotmaker.docker"

func NewKRestartPolicyAlwaysRestart() whaleAquarium.RestartPolicy {
	return whaleAquarium.KRestartPolicyOnFailure
}
