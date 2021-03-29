package iotmakerdocker

import (
	"bytes"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/stdcopy"
	"log"
)

type ExecResult struct {
	StdOut   string
	StdErr   string
	ExitCode int
}

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
			Cmd:          commands,
			Privileged:   true,
			AttachStderr: true,
			AttachStdin:  true,
			AttachStdout: true,
		},
	)
	if err != nil {
		return
	}

	//var e types.ExecStartCheck
	//err = el.cli.ContainerExecStart(el.ctx, idResponse.ID, e)
	//if err != nil {
	//	return
	//}

	var resp types.HijackedResponse
	resp, err = el.cli.ContainerExecAttach(el.ctx, idResponse.ID, types.ExecStartCheck{})
	if err != nil {
		return
	}
	defer resp.Close()

	select {
	case <-el.ctx.Done():
	}

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	_, err = stdcopy.StdCopy(stdout, stderr, resp.Reader)

	if err != nil {
		return
	}

	log.Println(stdout.String())

	var i types.ContainerExecInspect
	i, err = el.cli.ContainerExecInspect(el.ctx, idResponse.ID)

	exitCode = i.ExitCode
	runing = i.Running

	return
}
