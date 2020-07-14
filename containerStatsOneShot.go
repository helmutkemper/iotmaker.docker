package iotmakerDocker

import (
	"encoding/json"
	"github.com/docker/docker/api/types"
	"io/ioutil"
)

func (el *DockerSystem) ContainerStatisticsOneShot(
	id string,
) (
	err error,
	statsRet types.Stats,
) {

	var stats types.ContainerStats
	var body []byte

	stats, err = el.cli.ContainerStats(el.ctx, id, false)
	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(stats.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &statsRet)
	return
}
