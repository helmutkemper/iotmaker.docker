package iotmakerDocker

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"strings"
	"testing"
)

var pullStatusChannel chan ContainerPullStatusSendToChannel

func RunBeforeTestToMakeAChannel() {
	pullStatusChannel = make(chan ContainerPullStatusSendToChannel, 1)

	go func(c chan ContainerPullStatusSendToChannel) {

		for {
			select {
			case status := <-c:
				fmt.Printf("image pull status: %+v\n", status)

				if status.Closed == true {
					fmt.Println("image pull complete!")
					return
				}
			}
		}

	}(pullStatusChannel)
}

func ImageBuildFromRemoteServer(t *testing.T, d *DockerSystem, server, imageName string, tags []string, channel *chan ContainerPullStatusSendToChannel) {
	var err = d.ImageBuildFromRemoteServer(server, imageName, tags, channel)
	if err != nil {
		t.Fatalf("error: %v", err.Error())
		return
	}
}

func ContainerCreateChangeExposedPortAndStart(t *testing.T, d *DockerSystem, imageName, containerName string, restartPolicy RestartPolicy, mountVolumes []mount.Mount, containerNetwork *network.NetworkingConfig, currentPort, changeToPort []nat.Port) {
	var err error
	var containerID string

	err, containerID = d.ContainerCreateChangeExposedPortAndStart(imageName, containerName, restartPolicy, mountVolumes, containerNetwork, currentPort, changeToPort)
	if err != nil {
		t.Fatalf("error: %v", err.Error())
		return
	}

	if containerID == "" {
		t.Fatal("we have a bug! ContainerCreateChangeExposedPortAndStart()")
	}
}

func ImageGarbageCollector(t *testing.T, d *DockerSystem) {
	var err error
	var imageList []types.ImageSummary
	err, imageList = d.ImageList()
	if err != nil {
		t.Fatalf("error: %v", err.Error())
		return
	}

	var pass = false
	for _, image := range imageList {
		if image.RepoTags[0] == "<none>:<none>" {
			pass = true
			break
		}
	}
	if pass == false {
		t.Fatal("garbage images not found")
		return
	}

	err = d.ImageGarbageCollector()
	if err != nil {
		t.Fatalf("error: %v", err.Error())
		return
	}

	err, imageList = d.ImageList()
	if err != nil {
		t.Fatalf("error: %v", err.Error())
		return
	}

	pass = true
	for _, image := range imageList {
		if image.RepoTags[0] == "<none>:<none>" {
			pass = false
			break
		}
	}
	if pass == false {
		t.Fatal("garbage images found after garbage collector runs")
		return
	}
}

func ContainerStopAndRemove(t *testing.T, d *DockerSystem) {
	var err error
	var list []types.Container
	err, list = d.ContainerListAll()
	if err != nil {
		t.Fatalf("error: %v", err.Error())
		return
	}

	for _, container := range list {
		if len(container.Names) > 0 && strings.Contains(container.Names[0], "delete") {
			err = d.ContainerStopAndRemove(container.ID, true, true, true)
			if err != nil {
				t.Fatalf("error: %v", err.Error())
				return
			}
		}
	}

	err, list = d.ContainerListAll()
	if err != nil {
		t.Fatalf("error: %v", err.Error())
		return
	}

	for _, container := range list {
		if len(container.Names) > 0 && strings.Contains(container.Names[0], "delete") {
			t.Fatal("we have a bug! ContainerStopAndRemove(ID)")
			return
		}
	}
}

func ImageRemoveByName(t *testing.T, d *DockerSystem) {
	var err error
	err = d.ImageRemoveByName("delete:latest")
	if err != nil {
		t.Fatalf("error: %v", err.Error())
		return
	}
}

func TestDockerSystem_ImageBuildFromRemoteServer(t *testing.T) {
	var serverPath = "https://github.com/helmutkemper/iotmaker.docker.util.install.container.git"
	var newImageName = "delete:latest"
	var newContainerName = "delete"
	var imageTags = make([]string, 0)
	var restartPolicy RestartPolicy = KRestartPolicyUnlessStopped
	var mountVolumes = []mount.Mount{
		{
			Type:   "bind",
			Source: "//var/run/docker.sock",
			Target: "/var/run/docker.sock",
		},
	}
	var containerNetwork *network.NetworkingConfig = nil
	var currentPort = []nat.Port{
		"3000/tcp",
	}
	var changeToPort = []nat.Port{
		"3000/tcp",
	}

	var err error
	var dockerSys = DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		t.Fatalf("error: %v", err.Error())
		return
	}

	RunBeforeTestToMakeAChannel()
	ImageBuildFromRemoteServer(t, &dockerSys, serverPath, newImageName, imageTags, &pullStatusChannel)
	ContainerCreateChangeExposedPortAndStart(t, &dockerSys, newImageName, newContainerName, restartPolicy, mountVolumes, containerNetwork, currentPort, changeToPort)
	ImageGarbageCollector(t, &dockerSys)
	ContainerStopAndRemove(t, &dockerSys)
	ImageRemoveByName(t, &dockerSys)
}
