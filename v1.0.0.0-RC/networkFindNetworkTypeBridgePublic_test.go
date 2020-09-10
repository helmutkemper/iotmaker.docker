package iotmakerDocker

import (
	"errors"
	"github.com/docker/docker/api/types"
	"reflect"
	"testing"
)

func TestDockerSystem_NetworkFindNetworkTypeBridgePublic(t *testing.T) {
	var err error
	var inspect types.NetworkResource

	d := DockerSystem{}
	err = d.Init()
	if err != nil {
		t.Fail()
		panic(err)
	}

	_, _, err = d.NetworkCreate(
		"delete_before_test",
		KNetworkDriveBridge,
		"local",
		"10.0.0.0/16",
		"10.0.0.1",
	)
	if err != nil {
		t.Fail()
		panic(err)
	}

	inspect, err = d.NetworkFindNetworkTypeBridgePublic()
	if err != nil {
		t.Fail()
		panic(err)
	}

	if reflect.DeepEqual(inspect, types.NetworkResource{}) == true {
		t.Fail()
		panic(errors.New("pubic network not found"))
	}
}
