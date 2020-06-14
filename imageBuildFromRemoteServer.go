package iotmakerDocker

import (
	"github.com/docker/docker/api/types"
	"io"
)

// en: Make a image from folder path content
//     Please note: dockerfile name must be "Dockerfile" inside root folder
//
//     For get a github token
//     settings > Developer settings > Personal access tokens > Generate new token
//     Mark [x]repo - Full control of private repositories
//
//     Example:
//       err, dockerSys := factoryDocker.NewClient()
//       if err != nil {
//         panic(err)
//       }
//       server := "https://github.com/__USER__/__PROJECT__.git"
//       server := "https://x-access-token:__TOKEN__@github.com/__USER__/__PROJECT__.git"
//       err = dockerSys.ImageBuildFromRemoteServer(server, []string{"server:latest"})
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
//       git server content
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
func (el *DockerSystem) ImageBuildFromRemoteServer(server string, tags []string, channel *chan ContainerPullStatusSendToChannel) (err error) {
	var imageBuildOptions types.ImageBuildOptions
	var reader io.Reader

	imageBuildOptions = types.ImageBuildOptions{
		Tags:          tags,
		Remove:        true,
		RemoteContext: server,
	}

	err, reader = el.imageBuild(nil, imageBuildOptions)
	el.processBuildAndPullReaders(&reader, channel)

	return
}
