package iotmakerdocker

import (
	"bytes"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/stdcopy"
	"log"
	"time"
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
	stdout *bytes.Buffer,
	err error,
) {

	log.Print("entrou 0")
	stdout = &bytes.Buffer{}

	var idResponse types.IDResponse
	time.Sleep(1 * time.Second)
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
	log.Print("entrou 1")
	var resp types.HijackedResponse
	resp, err = el.cli.ContainerExecAttach(el.ctx, idResponse.ID, types.ExecStartCheck{})
	if err != nil {
		return
	}
	defer resp.Close()
	log.Print("entrou 2")
	//select {
	//case <-el.ctx.Done():
	//}
	log.Print("entrou 3")
	stderr := new(bytes.Buffer)
	_, err = stdcopy.StdCopy(stdout, stderr, resp.Reader)
	log.Print("entrou 4")
	if err != nil {
		return
	}

	var i types.ContainerExecInspect
	i, err = el.cli.ContainerExecInspect(el.ctx, idResponse.ID)
	if err != nil {
		return
	}
	log.Print("entrou 5")
	exitCode = i.ExitCode
	runing = i.Running

	return
}
