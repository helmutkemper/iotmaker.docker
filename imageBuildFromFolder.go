package iotmakerDocker

import (
	"bytes"
	"github.com/docker/docker/api/types"
)

// en: Make a image from folder path content
//     Please note: dockerfile name must be "Dockerfile" inside root folder
//
//     Example:
//       err, dockerSys := factoryDocker.NewClient()
//       if err != nil {
//         panic(err)
//       }
//       err = dockerSys.ImageBuildFromFolder("./folder", []string{"server:latest"})
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
func (el *DockerSystem) ImageBuildFromFolder(folderPath string, tags []string) (err error) {
	var tarFileReader *bytes.Reader
	var imageBuildOptions types.ImageBuildOptions

	err, tarFileReader = el.imageBuildPrepareFolderContext(folderPath)
	if err != err {
		return
	}

	imageBuildOptions = types.ImageBuildOptions{
		Tags:   tags,
		Remove: true,
	}

	err = el.imageBuild(tarFileReader, imageBuildOptions)

	return
}
