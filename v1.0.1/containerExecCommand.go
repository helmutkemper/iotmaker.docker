package iotmakerdocker

import (
	"bytes"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/stdcopy"
	"io/ioutil"
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
	stdout *bytes.Buffer,
	err error,
) {

	stdout = new(bytes.Buffer)

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

	var written int64

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

	stderr := new(bytes.Buffer)
	if err != nil {
		return
	}

	var i types.ContainerExecInspect
	i, err = el.cli.ContainerExecInspect(el.ctx, idResponse.ID)
	if err != nil {
		return
	}

	written, err = stdcopy.StdCopy(stdout, stderr, resp.Reader)
	log.Printf("------------------------written: %v", written)
	var out []byte
	out, err = ioutil.ReadAll(stdout)
	log.Printf("%s", out)

	exitCode = i.ExitCode
	runing = i.Running

	return
}
