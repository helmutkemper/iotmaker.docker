//https://stackoverflow.com/questions/38804313/build-docker-image-from-go-code

package main

import (
	"archive/tar"
	"bytes"
	"github.com/docker/docker/api/types"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker/factoryDocker"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	err, dockerSys := factoryDocker.NewClient()
	if err != nil {
		panic(err)
	}

	dockerSys.ImageMountContext("./lixo")
}

func _NotOk() {

	err, dockerSys := factoryDocker.NewClient()
	if err != nil {
		panic(err)
	}

	dockerBuildContext, err := os.Open("./lixo.tar")
	if err != nil {
		panic(err)
	}
	defer dockerBuildContext.Close()

	t := types.ImageBuildOptions{
		Tags: []string{
			"kemper:latest",
		},
		Remove:     true,
		PullParent: true,
		//Context:        dockerBuildContext,
	}

	err, imageBuildResponse := dockerSys.ImageBuild(dockerBuildContext, t)
	if err != nil {
		panic(err)
	}

	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		log.Fatal(err, " :unable to read image build response")
	}

	err, c := dockerSys.ContainerBuild()
	if err != nil {
		panic(err)
	}

	_ = c.ID
}

func ok() {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	dockerFile := "./Dockerfile"
	dockerFileReader, err := os.Open("./Dockerfile")
	if err != nil {
		log.Fatal(err, " :unable to open Dockerfile")
	}
	readDockerFile, err := ioutil.ReadAll(dockerFileReader)
	if err != nil {
		log.Fatal(err, " :unable to read dockerfile")
	}

	tarHeader := &tar.Header{
		Name: dockerFile,
		Size: int64(len(readDockerFile)),
	}
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		log.Fatal(err, " :unable to write tar header")
	}
	_, err = tw.Write(readDockerFile)
	if err != nil {
		log.Fatal(err, " :unable to write tar body")
	}
	dockerFileTarReader := bytes.NewReader(buf.Bytes())
	_ = dockerFileTarReader

	dockerSys := iotmakerDocker.DockerSystem{}
	err = dockerSys.Init()
	if err != nil {
		panic(err)
	}

	t := types.ImageBuildOptions{
		Tags: []string{
			"kemper:latest",
		},
		Remove:     true,
		PullParent: true,
		//Context:        dockerFileTarReader,
		//RemoteContext: "git@github.com:helmutkemper/lixo.git",
		//RemoteContext: "https://github.com/helmutkemper/lixo.git",
		RemoteContext: "https://x-access-token:bf98ec92704f3ee461ec274eeda8e82cb7cee32b@github.com/helmutkemper/lixo.git",
	} //a5596205290ccac09003b3b8b3c2378eac135c53

	err, imageBuildResponse := dockerSys.ImageBuild(nil, t)
	if err != nil {
		panic(err)
	}

	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		log.Fatal(err, " :unable to read image build response")
	}

	err, c := dockerSys.ContainerBuild()
	if err != nil {
		panic(err)
	}

	_ = c.ID
}
