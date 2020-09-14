package iotmakerdocker

// Must be first function call
func (el *DockerSystem) Init() (err error) {

	el.ContextCreate()
	return el.ClientCreate()
}
