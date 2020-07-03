package iotmakerDocker

func (el *DockerSystem) getNetworkGenerator(name string) *NextNetworkAutoConfiguration {
	return el.networkGenerator[name]
}
