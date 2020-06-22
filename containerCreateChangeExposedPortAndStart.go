package iotmakerDocker

import (
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
)

// en: Make a image from folder path content
//     Please note: dockerfile name must be "Dockerfile" inside root folder
//
//     Example:
//       err, dockerSys := factoryDocker.NewClient()
//       if err != nil {
//         panic(err)
//       }
//       err, id := dockerSys.ContainerCreateChangeExposedPortAndStart(
//       		"server:latest",
//       		"server",
//       		dockerSys.KRestartPolicyUnlessStopped,
//       		[]mount.Mount{},
//       		nil,
//       		[]nat.Port{
//       			"tcp/8080",
//       		},
//       		[]nat.Port{
//       			"tcp/9000",
//       		},
//       )
//       if err != nil {
//         panic(err)
//       }
//
//       ./folder
//          Dockerfile
//            FROM golang:latest
//            RUN mkdir /app
//            RUN chmod 700 /app
//            COPY . /app
//            WORKDIR /app
//            EXPOSE 8080
//            RUN go build -o ./main ./main.go
//            CMD ["./main"]
//
//          main.go
//            package main
//            import (
//            	"fmt"
//            	"net/http"
//            )
//            func hello(w http.ResponseWriter, req *http.Request) {
//            	fmt.Fprintf(w, "hello\n")
//            }
//            func main() {
//            	http.HandleFunc("/hello", hello)
//            	http.ListenAndServe(":8080", nil)
//            }
func (el *DockerSystem) ContainerCreateChangeExposedPortAndStart(
	imageName,
	containerName string,
	restart RestartPolicy,
	mountVolumes []mount.Mount,
	containerNetwork *network.NetworkingConfig,
	currentPort,
	changeToPort []nat.Port,
) (err error, containerID string) {

	imageName = el.AdjustImageName(imageName)

	err, containerID = el.ContainerCreateAndChangeExposedPort(
		imageName,
		containerName,
		restart,
		mountVolumes,
		containerNetwork,
		currentPort,
		changeToPort,
	)
	if err != nil {
		return
	}

	err = el.ContainerStart(containerID)
	return
}
