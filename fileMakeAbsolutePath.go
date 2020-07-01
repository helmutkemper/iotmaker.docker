package iotmakerDocker

import "path/filepath"

// Get an absolute path from file
func (el *DockerSystem) FileMakeAbsolutePath(
	filePath string,
) (
	err error,
	fileAbsolutePath string,
) {

	fileAbsolutePath, err = filepath.Abs(filePath)
	return
}
