package iotmaker_docker

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
)

// pt-br: cria um novo struct para rede do container
// en: create a new container network struct
func NewContainerNetwork() ContainerNetwork {
	return ContainerNetwork{}
}

// pt-br: container network struct
// en: container network struct
type ContainerNetwork struct {
	data *types.ContainerJSON
}

// pt-br: lista de ligações para o container
// en: list of volume bindings for this container
func (el *ContainerNetwork) GetBinds() []string {
	return el.data.HostConfig.Binds
}

// pt-br: arquivo (caminho) onde o container id está escrito
// en: file (path) where the containerId is written
func (el *ContainerNetwork) GetContainerIDFile() string {
	return el.data.HostConfig.ContainerIDFile
}

// pt-br: configuração de log do container
// en: configuration of the logs for this Container
func (el *ContainerNetwork) GetLogConfig() container.LogConfig {
	return el.data.HostConfig.LogConfig
}

// pt-br: modo de rede do container
// en: network mode to use for the Container
func (el *ContainerNetwork) GetNetworkMode() container.NetworkMode {
	return el.data.HostConfig.NetworkMode
}

// pt-br: mapa de portas entre o container e o host
// en: port mapping between the exposed port (Container) and the host
func (el *ContainerNetwork) GetPortBindings() nat.PortMap {
	return el.data.HostConfig.PortBindings
}

// pt-br: política de reinício do container
// en: restart policy to be used for the Container
func (el *ContainerNetwork) GetRestartPolicy() container.RestartPolicy {
	return el.data.HostConfig.RestartPolicy
}

// pt-br: remoção automática do container quando finalizado
// en: automatically remove Container when it exits
func (el *ContainerNetwork) GetAutoRemove() bool {
	return el.data.HostConfig.AutoRemove
}

// pt-br: nome do driver usado para montar os discos
// en: name of the volume driver used to mount volumes
func (el *ContainerNetwork) GetVolumeDriver() string {
	return el.data.HostConfig.VolumeDriver
}

func (el *ContainerNetwork) GetVolumesFrom() []string {
	return el.data.HostConfig.VolumesFrom
}

// pt-br: devolve o nome da rede
// en: Bridge is the Bridge name the network uses(e.g. `docker0`)
func (el *ContainerNetwork) GetBridge() string {
	return el.data.NetworkSettings.Bridge
}

// pt-br: SandboxID representa a pilha de rede do container
// en: SandboxID uniquely represents a Container's network stack
func (el *ContainerNetwork) GetSandboxID() string {
	return el.data.NetworkSettings.SandboxID
}

// en: HairpinMode specifies if hairpin NAT should be enabled on the virtual interface
func (el *ContainerNetwork) GetHairpinMode() bool {
	return el.data.NetworkSettings.HairpinMode
}

// en: LinkLocalIPv6Address is an IPv6 unicast address using the link-local prefix
func (el *ContainerNetwork) GetLinkLocalIPv6Address() string {
	return el.data.NetworkSettings.LinkLocalIPv6Address
}

// en: LinkLocalIPv6PrefixLen is the prefix length of an IPv6 unicast address
func (el *ContainerNetwork) GetLinkLocalIPv6PrefixLen() int {
	return el.data.NetworkSettings.LinkLocalIPv6PrefixLen
}

// en: Ports is a collection of PortBinding indexed by Port
func (el *ContainerNetwork) GetPorts() nat.PortMap {
	return el.data.NetworkSettings.Ports
}

// en: SandboxKey identifies the sandbox
func (el *ContainerNetwork) GetSandboxKey() string {
	return el.data.NetworkSettings.SandboxKey
}

// en: secondary container ip address
func (el *ContainerNetwork) GetSecondaryIPAddresses() []network.Address {
	return el.data.NetworkSettings.SecondaryIPAddresses
}

// en: secondary container ipv6 address
func (el *ContainerNetwork) GetSecondaryIPv6Addresses() []network.Address {
	return el.data.NetworkSettings.SecondaryIPv6Addresses
}
