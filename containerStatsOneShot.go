package iotmaker_docker

import (
	"encoding/json"
	"github.com/docker/docker/api/types"
	"io/ioutil"
)

func (el *DockerSystem) ContainerStatsOneShot(id string) (error, types.Stats) {
	var err error
	var stats types.ContainerStats
	var body []byte
	var ret types.Stats

	stats, err = el.cli.ContainerStats(el.ctx, id, true)
	if err != nil {
		return err, ret
	}

	body, err = ioutil.ReadAll(stats.Body)
	if err != nil {
		return err, ret
	}

	err = json.Unmarshal(body, &ret)
	return err, ret
}
