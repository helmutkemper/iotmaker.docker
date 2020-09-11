package iotmakerdocker

// ContainerWait waits until the specified container is in a certain state
// indicated by the given condition, either "not-running" (default),
// "next-exit", or "removed".
func (el *DockerSystem) ContainerWaitStatusNotRunning(
	id string,
) (
	err error,
) {

	wOk, wErr := el.cli.ContainerWait(el.ctx, id, "not-running")
	select {
	case <-wOk:
	case err = <-wErr:
	}
	return
}
