package iotmaker_docker

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
)

// pt-br: cria um novo struct para rede do container
// en: create a new container network struct
func NewContainerNetwork() networkStt {
	return networkStt{}
}

// pt-br: container network struct
// en: container network struct
type networkStt struct {
	data *types.ContainerJSON
}

// pt-br: lista de ligações para o container
// en: list of volume bindings for this container
func (el *networkStt) GetBinds() []string {
	return el.data.HostConfig.Binds
}

// pt-br: arquivo (caminho) onde o container id está escrito
// en: file (path) where the containerId is written
func (el *networkStt) GetContainerIDFile() string {
	return el.data.HostConfig.ContainerIDFile
}

// pt-br: configuração de log do container
// en: configuration of the logs for this containerStt
func (el *networkStt) GetLogConfig() container.LogConfig {
	return el.data.HostConfig.LogConfig
}

// pt-br: modo de rede do container
// en: network mode to use for the containerStt
func (el *networkStt) GetNetworkMode() container.NetworkMode {
	return el.data.HostConfig.NetworkMode
}

// pt-br: mapa de portas entre o container e o host
// en: port mapping between the exposed port (containerStt) and the host
func (el *networkStt) GetPortBindings() nat.PortMap {
	return el.data.HostConfig.PortBindings
}

// pt-br: política de reinício do container
// en: restart policy to be used for the containerStt
func (el *networkStt) GetRestartPolicy() container.RestartPolicy {
	return el.data.HostConfig.RestartPolicy
}

// pt-br: remoção automática do container quando finalizado
// en: automatically remove containerStt when it exits
func (el *networkStt) GetAutoRemove() bool {
	return el.data.HostConfig.AutoRemove
}

// pt-br: nome do driver usado para montar os discos
// en: name of the volume driver used to mount volumes
func (el *networkStt) GetVolumeDriver() string {
	return el.data.HostConfig.VolumeDriver
}

func (el *networkStt) GetVolumesFrom() []string {
	return el.data.HostConfig.VolumesFrom
}

// pt-br: devolve o nome da rede
// en: Bridge is the Bridge name the network uses(e.g. `docker0`)
func (el *networkStt) GetBridge() string {
	return el.data.NetworkSettings.Bridge
}

// pt-br: SandboxID representa a pilha de rede do container
// en: SandboxID uniquely represents a containerStt's network stack
func (el *networkStt) GetSandboxID() string {
	return el.data.NetworkSettings.SandboxID
}

// en: HairpinMode specifies if hairpin NAT should be enabled on the virtual interface
func (el *networkStt) GetHairpinMode() bool {
	return el.data.NetworkSettings.HairpinMode
}

// en: LinkLocalIPv6Address is an IPv6 unicast address using the link-local prefix
func (el *networkStt) GetLinkLocalIPv6Address() string {
	return el.data.NetworkSettings.LinkLocalIPv6Address
}

// en: LinkLocalIPv6PrefixLen is the prefix length of an IPv6 unicast address
func (el *networkStt) GetLinkLocalIPv6PrefixLen() int {
	return el.data.NetworkSettings.LinkLocalIPv6PrefixLen
}

// en: Ports is a collection of PortBinding indexed by Port
func (el *networkStt) GetPorts() nat.PortMap {
	return el.data.NetworkSettings.Ports
}

// en: SandboxKey identifies the sandbox
func (el *networkStt) GetSandboxKey() string {
	return el.data.NetworkSettings.SandboxKey
}

// en: secondary container ip address
func (el *networkStt) GetSecondaryIPAddresses() []network.Address {
	return el.data.NetworkSettings.SecondaryIPAddresses
}

// en: secondary container ipv6 address
func (el *networkStt) GetSecondaryIPv6Addresses() []network.Address {
	return el.data.NetworkSettings.SecondaryIPv6Addresses
}
