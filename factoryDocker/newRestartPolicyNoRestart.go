package factoryDocker

import whaleAquarium "github.com/helmutkemper/iotmaker.docker"

func NewRestartPolicyNoRestart() whaleAquarium.RestartPolicy {
	return whaleAquarium.KRestartPolicyNo
}
