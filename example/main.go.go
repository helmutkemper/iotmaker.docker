// this example show how to use this project, downloading a mongodb container,
// installing it, testing and removing container
package main

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/helmutkemper/iotmaker.db.mongodb.config/factoryMongoDBConfig"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	factoryDocker "github.com/helmutkemper/iotmaker.docker/factoryDocker"
	"github.com/helmutkemper/iotmaker.docker/util"
	"io/ioutil"
	"log"
	"os"
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
	var volumesInitialList []types.Volume
	var networkInitialList []types.NetworkResource
	var imagesInitialList []types.ImageSummary
	var imagesFinalList []types.ImageSummary
	var nextNetworkConfig *network.NetworkingConfig
	var currentPort nat.Port
	var newPort nat.Port
	var currentPortList []nat.Port
	var newPortList []nat.Port
	var networkUtil util.NextNetworkConfiguration
	var imageListBeforeTest []types.ImageSummary
	var listOfExposedPortsByName, listOfExposedPortsById []string
	var imageID, imageNAME, imageIdFindByName string
	//var client *mongo.Client

	//var mongoDbURL = "mongodb://localhost:27017"
	var networkName = "network_test"
	var relativeMongoDBConfigFilePathToGenerateAndSave = "./config.conf"
	var imageName = "mongo:latest"
	var containerName = "containerTest"
	var mongoDBDefaultPort = "27017"
	var mongoDBOutputPort = "27017"
	var mongoDbConnectionProtocol = "tcp"

	// init docker
	var dockerSys = iotmakerDocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		panic(nil)
	}

	err, networkInitialList = dockerSys.NetworkList()
	if err != nil {
		panic(nil)
	}
	_ = networkInitialList

	err, volumesInitialList = dockerSys.VolumeList()
	if err != nil {
		panic(nil)
	}
	_ = volumesInitialList

	err, imagesInitialList = dockerSys.ImageList()
	if err != nil {
		panic(nil)
	}
	_ = imagesInitialList

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
	if err != nil {
		panic(err)
	}
	/*
		  // wait MongoDB start time inside container
			time.Sleep(time.Second * 60)
			client, err = mongo.NewClient(options.Client().ApplyURI(mongoDbURL))
			if err != nil {
				log.Fatal(err)
			}
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			err = client.Connect(ctx)
			if err != nil {
				log.Fatal(err)
			}

			err = client.Ping(context.Background(), nil)
			if err != nil {
				log.Fatal(err)
			}

		  err = client.Disconnect(ctx)
		  if err != nil {
		    log.Fatal(err)
		  }
	*/
	err = dockerSys.ContainerStopAndRemove(id, true, true, true)
	if err != nil {
		log.Fatal(err)
	}

	err = dockerSys.NetworkRemoveByName(networkName)
	if err != nil {
		panic(nil)
	}

	// Remove todas as imagens que não estiverem na lista inicial
	err, imagesFinalList = dockerSys.ImageList()
	if err != nil {
		panic(nil)
	}
	for _, imgInicial := range imagesInitialList {
		var pass = false
		for _, imgFinal := range imagesFinalList {
			if imgInicial.ID == imgFinal.ID {
				pass = true
				break
			}
		}

		if pass == true {
			err = dockerSys.ImageRemove(imgInicial.ID)
			if err != nil {
				panic(nil)
			}
		}
	}

	fmt.Println("test pass!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

	// Output:
	//
}
