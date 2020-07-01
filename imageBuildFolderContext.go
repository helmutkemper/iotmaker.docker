package iotmakerDocker

import (
	"archive/tar"
	"bytes"
	"strings"
)

func (el DockerSystem) imageBuildPrepareFolderContext(
	dirPath string,
) (
	err error,
	file *bytes.Reader,
) {

	var buf bytes.Buffer
	var tarWriter *tar.Writer
	tarWriter = tar.NewWriter(&buf)

	if strings.HasSuffix(dirPath, "/") == false {
		dirPath += "/"
	}

	err = el.imageBuildPrepareFolderContextSupport(dirPath, dirPath, &buf, tarWriter)

	err = tarWriter.Close()
	if err != nil {
		return
	}

	file = bytes.NewReader(buf.Bytes())

	return
}
