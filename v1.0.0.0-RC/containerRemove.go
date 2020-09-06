package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
)

// English: Remove a container by id
//   id - container id
//   removeVolumes - remove container and volumes
//   removeLinks - remove container and links
//   force - force remove
//
// Português: Remove container por id
//   id - container id
//   removeVolumes - remove o container e os volumes
//   removeLinks - remove o container e os links
//   force - força a emoção
func (el *DockerSystem) ContainerRemove(
	id string,
	removeVolumes,
	removeLinks,
	force bool,
) (
	err error,
) {

	return el.cli.ContainerRemove(
		el.ctx,
		id,
		types.ContainerRemoveOptions{
			RemoveVolumes: removeVolumes,
			RemoveLinks:   removeLinks,
			Force:         force,
		},
	)
}
