package iotmakerdocker

import (
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
)

// create network
//    name:    string       Ex.: "containerNetwork"
//    drive:   NetworkDrive Ex.: KNetworkDriveBridge
//    scope:   string       Ex.: local
//    subnet:  string       Ex.: 10.0.0.0/16 (note: use base 10)
//    gateway: string       Ex.: 10.0.0.1    (note: use base 10)
func (el *DockerSystem) NetworkCreate(
	name string,
	drive NetworkDrive,
	scope,
	subnet,
	gateway string,
) (
	id string,
	networkGenerator *NextNetworkAutoConfiguration,
	err error,
) {

	//todo: se já tem uma rede, ajustar o ip automático para o próximo endereço
	var resp types.NetworkCreateResponse
	var insp types.NetworkResource
	var gatewayFieldA, gatewayFieldB, gatewayFieldC, gatewayFieldD int

	networkGenerator = &NextNetworkAutoConfiguration{}

	id, _ = el.NetworkFindIdByName(name)
	if id != "" {

		insp, err = el.cli.NetworkInspect(
			el.ctx,
			id,
			types.NetworkInspectOptions{
				Scope:   scope,
				Verbose: false,
			},
		)
		if err != nil {
			return
		}
		pass := false
		for _, v := range insp.IPAM.Config {
			if v.Gateway == gateway && v.Subnet == subnet {
				pass = true
				break
			}
		}

		if pass == true {
			return
		}

		err = errors.New("there is a network with this name")
		return
	}

	gatewayFieldA, gatewayFieldB, gatewayFieldC, gatewayFieldD, err = el.testGatewayValues(gateway)
	if err != nil {
		return
	}

	err = el.testSubnetValues(subnet)
	if err != nil {
		return
	}

	resp, err = el.cli.NetworkCreate(el.ctx, name, types.NetworkCreate{
		//CheckDuplicate: false,
		Driver: drive.String(),
		Scope:  scope,
		IPAM: &network.IPAM{
			Driver: "default",
			Config: []network.IPAMConfig{
				{
					Subnet:  subnet,
					Gateway: gateway,
				},
			},
		},
		Attachable: true,
		Labels: map[string]string{
			"name": name,
		},
	})
	if err != nil {
		return
	}

	networkGenerator.Init(
		resp.ID,
		name,
		gateway,
		byte(gatewayFieldA),
		byte(gatewayFieldB),
		byte(gatewayFieldC),
		byte(gatewayFieldD),
	)

	if len(el.networkId) == 0 {
		el.networkId = make(map[string]string)
	}

	el.networkId[name] = resp.ID
	id = resp.ID

	if len(el.networkGenerator) == 0 {
		el.networkGenerator = make(map[string]*NextNetworkAutoConfiguration)
	}

	el.networkGenerator[name] = networkGenerator

	return
}
