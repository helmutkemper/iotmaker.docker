package iotmaker_docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"time"
)

// pt-br: cria um novo struct para containers
// en: create a new container struct
func NewContainer() (error, Container) {
	var c = Container{}
	var err = c.Init()

	return err, c
}

// pt-br: construct do container
// en: construct of container
type Container struct {
	context    context.Context
	client     *client.Client
	awaysList  bool
	hasStarted bool
	Error      error
	list       []types.Container
}

// pt-br: inicializa todas as funções críticas
// en: start all critical functions
func (el *Container) Init() error {
	el.context = GetContext()
	el.client, el.Error = GetClient()
	if el.Error != nil {
		return el.Error
	}

	el.hasStarted = true

	return nil
}

// pt-br: força a atualização da lista de containers
// en: force of the updating of the list of containers
func (el *Container) SetAwaysList(value bool) {
	el.awaysList = value
}

// pt-br: lista todos os containers
// en: list all containers
func (el *Container) GetList() (error, []types.Container) {
	if len(el.list) == 0 || el.awaysList {
		_, _ = el.ListAll()
	}

	return el.Error, el.list
}

// pr-br: para o container por id
// en: stop the container by id
func (el *Container) StopById(id string, timeOut time.Duration) error {
	if timeOut == 0 {
		timeOut = kTimeOut
	}

	if !el.hasStarted {
		_ = el.Init()
		if el.Error != nil {
			return el.Error
		}
	}

	el.Error = el.client.ContainerStop(el.context, id, &timeOut)

	return el.Error
}

// pt-br: para todos os containers rodando
// en: stop all running containers
func (el *Container) StopAll(timeOut time.Duration) error {
	if timeOut == 0 {
		timeOut = kTimeOut
	}

	if !el.hasStarted {
		_ = el.Init()
		if el.Error != nil {
			return el.Error
		}
	}

	_, err := el.ListAll()
	if err != nil {
		return el.Error
	}

	for _, containerData := range el.list {
		if containerData.State == kContainerStateRunning {
			el.Error = el.client.ContainerStop(el.context, containerData.ID, &timeOut)
			if el.Error != nil {
				return el.Error
			}
		}
	}

	return nil
}

// pt-br: lista todos os containers
// en: list all containers
func (el *Container) ListAll() (error, []types.Container) {
	el.list = make([]types.Container, 0)

	if !el.hasStarted {
		_ = el.Init()
		if el.Error != nil {
			return el.Error, make([]types.Container, 0)
		}
	}

	el.list, el.Error = el.client.ContainerList(el.context, types.ContainerListOptions{All: true})
	if el.Error != nil {
		return el.Error, make([]types.Container, 0)
	}

	return el.Error, el.list
}

// pt-br: Inspeciona o container por by
// en: inspect a container by id
func (el *Container) InspectById(id string) (error, types.ContainerJSON) {
	var inspect = types.ContainerJSON{}

	if !el.hasStarted {
		_ = el.Init()
		if el.Error != nil {
			return el.Error, inspect
		}
	}

	inspect, _, el.Error = el.client.ContainerInspectWithRaw(el.context, id, true)

	return el.Error, inspect
}

// pt-br: retorna uma lista de ids em função do nome da imagem passada
// en: returns a list of ids by the name of the image
func (el *Container) ContainerGetIdByImageName(name string) (error, []string) {
	var ret = make([]string, 0)
	_, err := el.ListAll()
	if err != nil {
		return el.Error, ret
	}

	for _, containerData := range el.list {
		if name == containerData.Image {
			ret = append(ret, containerData.ID)
		}
	}

	return nil, ret
}

// pt-br: inicializa um container por id
// en: start a container by id
func (el *Container) ContainerStartById(id string, checkpointDir, checkpointID string) error {
	if !el.hasStarted {
		_ = el.Init()
		if el.Error != nil {
			return el.Error
		}
	}

	el.Error = el.client.ContainerStart(el.context, id, types.ContainerStartOptions{CheckpointDir: checkpointDir, CheckpointID: checkpointID})

	return el.Error
}

// pt-br: cria um container
// en: create a container
func (el *Container) ContainerCreate(config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, name string) (container.ContainerCreateCreatedBody, error) {
	createdBody := container.ContainerCreateCreatedBody{}
	if !el.hasStarted {
		_ = el.Init()
		if el.Error != nil {
			return createdBody, el.Error
		}
	}

	createdBody, el.Error = el.client.ContainerCreate(el.context, config, hostConfig, networkingConfig, name)

	return createdBody, el.Error
}

// pt-br: remove um container por id
// en: removes a container by id
func (el *Container) ContainerRemoveById(id string) error {
	if !el.hasStarted {
		_ = el.Init()
		if el.Error != nil {
			return el.Error
		}
	}

	el.Error = el.client.ContainerRemove(el.context, id, types.ContainerRemoveOptions{})

	return el.Error
}
