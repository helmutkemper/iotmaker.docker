// this example show how to use this project, downloading a mongodb container,
// installing it, testing and removing container
package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/helmutkemper/iotmaker.db.mongodb.config/factoryMongoDBConfig"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	factoryDocker "github.com/helmutkemper/iotmaker.docker/factoryDocker"
	"github.com/helmutkemper/iotmaker.docker/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var pullStatusChannel = factoryDocker.NewImagePullStatusChannel()

func prepareThreadToPrintImagePullStatus() {
	go func(c chan iotmakerDocker.ContainerPullStatusSendToChannel) {

		for {
			select {
			case status := <-c:
				fmt.Printf("image pull status: %+v\n", status)

				if status.Closed == true {
					fmt.Println("image pull complete!")
				}
			}
		}

	}(pullStatusChannel)
}

func main() {
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
	var listOfExposedPortsByName, listOfExposedPortsById []string
	var imageID, imageNAME, imageIdFindByName string
	var client *mongo.Client

	var mongoDbURL = "mongo://localhost:27017"
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

	prepareThreadToPrintImagePullStatus()
	err, imageID, imageNAME = dockerSys.ImagePull(imageName, pullStatusChannel)
	if err != nil {
		panic(nil)
	}

	_ = imageID
	_ = imageNAME

	//sha256:be8d903a68997dd63f64479004a7eeb4f0674dde7ab3cbd1145e5658da3a817b
	//sha256:66c68b650ad44f7a95c256ad2df5c40fbc3b13001f36ac7b7cd25f5f9a09be7d

	err, imageIdFindByName = dockerSys.ImageFindIdByName(imageName)

	err, listOfExposedPortsByName = dockerSys.ImageListExposedPortsByName(imageName)
	err, listOfExposedPortsById = dockerSys.ImageListExposedPorts(imageIdFindByName)

	_ = listOfExposedPortsByName
	_ = listOfExposedPortsById

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

	time.Sleep(time.Second * 5)
	client, err = mongo.NewClient(options.Client().ApplyURI(mongoDbURL))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = dockerSys.ContainerStop(id)

	// Output:
	//
}
