package iotmakerdocker

import (
	"github.com/docker/docker/api/types"
)

func (el *DockerSystem) ContainerExecCommand(
	id string,
	commands []string,
) (
	exitCode int,
	runing bool,
	err error,
) {

	var idResponse types.IDResponse
	idResponse, err = el.cli.ContainerExecCreate(
		el.ctx,
		id,
		types.ExecConfig{
			Cmd: commands,
		},
	)
	if err != nil {
		return
	}

	err = el.cli.ContainerExecStart(el.ctx, idResponse.ID, types.ExecStartCheck{})
	if err != nil {
		return
	}

	var inspect types.ContainerExecInspect
	inspect, err = el.cli.ContainerExecInspect(el.ctx, idResponse.ID)
	if err != nil {
		return
	}

	exitCode = inspect.ExitCode
	runing = inspect.Running

	return
}
