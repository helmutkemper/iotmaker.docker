package iotmakerDocker

import (
	"archive/tar"
	"bytes"
	"io/ioutil"
	"os"
	"strings"
)

func (el DockerSystem) ImageMountContext(dirPath string) (err error, file []byte) {
	var buf bytes.Buffer
	var tarWriter *tar.Writer
	tarWriter = tar.NewWriter(&buf)

	err = el.imageMountContextSupport(dirPath, &buf, tarWriter)

	err = tarWriter.Close()
	if err != nil {
		return
	}

	file = buf.Bytes()

	return
}

func (el DockerSystem) imageMountContextSupport(dirPath string, buf *bytes.Buffer, tarWriter *tar.Writer) (err error) {
	var dirContent []os.FileInfo
	var tarHeader *tar.Header
	var fileData []byte
	var filePath string

	if strings.HasSuffix(dirPath, "/") == false {
		dirPath += "/"
	}

	dirContent, err = ioutil.ReadDir(dirPath)
	if err != nil {
		return
	}

	for _, folderItem := range dirContent {
		filePath = dirPath + folderItem.Name()

		if folderItem.IsDir() == true {
			err = el.imageMountContextSupport(filePath, buf, tarWriter)
			if err != nil {
				return
			}
		} else {
			fileData, err = ioutil.ReadFile(filePath)
			if err != nil {
				return
			}

			tarHeader = &tar.Header{
				Name: filePath,
				Mode: 0600,
				Size: folderItem.Size(),
			}

			if err = tarWriter.WriteHeader(tarHeader); err != nil {
				return
			}
			if _, err = tarWriter.Write(fileData); err != nil {
				return
			}
		}
	}

	return
}
