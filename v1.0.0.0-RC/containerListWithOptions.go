package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

func (el *DockerSystem) ContainerListWithOptions(
	quiet bool,
	// English: populate types.Container.SiseRw and SizeRootFs
	size bool,
	all bool,
	latest bool,

	// English: example: "2020-09-08T00:39:53.613203298Z"
	since string,

	// English: example: "2020-09-08T00:39:53.613203298Z"
	before string,
	limit int,
	filters filters.Args,
) (
	list []types.Container,
	err error,
) {

	list, err = el.cli.ContainerList(
		el.ctx,
		types.ContainerListOptions{
			Quiet:   quiet,
			Size:    size,
			All:     all,
			Latest:  latest,
			Since:   since,
			Before:  before,
			Limit:   limit,
			Filters: filters,
		},
	)

	return
}
