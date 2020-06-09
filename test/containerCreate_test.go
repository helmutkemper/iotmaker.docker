package test

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/helmutkemper/iotmaker.db.mongodb.config/factoryMongoDBConfig"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	factoryDocker "github.com/helmutkemper/iotmaker.docker/factoryDocker"
	"github.com/helmutkemper/iotmaker.docker/util"
	"io/ioutil"
	"os"
)

func ExampleDockerSystem_ContainerCreate() {
	var err error
	var id string
	var file []byte
	var mountList []mount.Mount
	var nextNetworkConfig *network.NetworkingConfig
	var currentPort nat.Port
	var newPort nat.Port
	var currentPortList []nat.Port
	var newPortList []nat.Port
	var networkUtil util.NetworkGenerator
	var imageListBeforeTest []types.ImageSummary

	var networkName = "network_test"
	var relativeMongoDBConfigFilePathToGenerateAndSave = "./config.conf"
	var imageName = "mongo:latest"
	var containerName = "containerTest"
	var mongoDBDefaultPort = "27017"
	var mongoDBOutputPort = "27017"
	var mongoDbConnectionProtocol = "tcp"

	err, networkUtil = factoryDocker.NewContainerNetworkGenerator(networkName, 10, 0, 0, 1)
	if err != nil {
		panic(nil)
	}

	currentPort, err = nat.NewPort(mongoDbConnectionProtocol, mongoDBDefaultPort)
	if err != nil {
		panic(nil)
	}

	currentPortList = []nat.Port{
		currentPort,
	}

	newPort, err = nat.NewPort(mongoDbConnectionProtocol, mongoDBOutputPort)
	if err != nil {
		panic(nil)
	}

	newPortList = []nat.Port{
		newPort,
	}

	// basic MongoDB configuration
	var conf = factoryMongoDBConfig.NewBasicConfigWithEphemeralData()
	err, file = conf.ToYaml(0)
	if err != nil {
		panic(nil)
	}

	// save MongoDB configuration into disk
	err = ioutil.WriteFile(relativeMongoDBConfigFilePathToGenerateAndSave, file, os.ModePerm)
	if err != nil {
		panic(nil)
	}

	// init docker
	var dockerSys = iotmakerDocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		panic(nil)
	}

	// example: dockerSys.ImageList()
	err, imageListBeforeTest = dockerSys.ImageList()
	if err != nil {
		panic(nil)
	}

	_ = imageListBeforeTest

	// image pull and wait (true)
	err = dockerSys.ImagePull(imageName, false)
	if err != nil {
		panic(nil)
	}

	//dockerSys.ImageFindIdByName("mongo:latest")

	// define an external MongoDB config file path
	err, mountList = factoryDocker.NewVolumeMount(
		[]iotmakerDocker.Mount{
			{
				MountType:   iotmakerDocker.KVolumeMountTypeBind,
				Source:      relativeMongoDBConfigFilePathToGenerateAndSave,
				Destination: "/etc/mongo.conf",
			},
		},
	)
	if err != nil {
		panic(err)
	}

	err, nextNetworkConfig = networkUtil.GetNext()
	if err != nil {
		panic(err)
	}

	err, id = dockerSys.ContainerCreateChangeExposedPortAndStart(
		imageName,
		containerName,
		iotmakerDocker.KRestartPolicyUnlessStopped,
		mountList,
		nextNetworkConfig,
		currentPortList,
		newPortList,
	)

	err = dockerSys.ContainerStop(id)

	// Output:
	//
}
