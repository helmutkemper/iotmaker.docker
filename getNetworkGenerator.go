package iotmakerDocker

func (el *DockerSystem) GetNetworkGenerator(name string) *NextNetworkAutoConfiguration {
	return el.networkGenerator[name]
}
