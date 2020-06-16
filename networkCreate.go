package iotmakerDocker

import (
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
)

const (
	kGatewayExpressionRegular = `^(?P<fieldA>[0-9]{0,3})\.(?P<fieldB>[0-9]{0,3})\.(?P<fieldC>[0-9]{0,3})\.(?P<fieldD>[0-9]{0,3})$`
	kSubnetExpressionRegular  = `^(?P<fieldA>[0-9]{0,3})\.(?P<fieldB>[0-9]{0,3})\.(?P<fieldC>[0-9]{0,3})\.(?P<fieldD>[0-9]{0,3})/(?<range>[0-9]{0,3})$`
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
) (err error, id string, networkGenerator NextNetworkConfiguration) {

	var resp types.NetworkCreateResponse
	var gatewayFieldA, gatewayFieldB, gatewayFieldC, gatewayFieldD int

	_, id = el.NetworkFindIdByName(name)
	if id != "" {
		err = errors.New("there is a network with this name")
		return
	}

	err, gatewayFieldA, gatewayFieldB, gatewayFieldC, gatewayFieldD = el.testGatewayValues(gateway)
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
		Attachable: false,
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

	return
}
