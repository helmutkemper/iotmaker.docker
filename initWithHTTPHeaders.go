package iotmakerDocker

func (el *DockerSystem) InitWithHTTPHeaders(headers map[string]string) error {
	el.ContextCreate()
	return el.ClientCreateWithHTTPHeaders(headers)
}
