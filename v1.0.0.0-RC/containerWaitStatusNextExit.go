package iotmakerDocker

// ContainerWait waits until the specified container is in a certain state
// indicated by the given condition, either "not-running" (default),
// "next-exit", or "removed".
func (el *DockerSystem) ContainerWaitStatusNextExit(
	id string,
) (
	err error,
) {

	wOk, wErr := el.cli.ContainerWait(el.ctx, id, "next-exit")
	select {
	case <-wOk:
	case err = <-wErr:
	}
	return
}
