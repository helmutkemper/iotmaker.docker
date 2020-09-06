package iotmakerDocker

import "path/filepath"

// Get an absolute path from file
func (el *DockerSystem) FileMakeAbsolutePath(
	filePath string,
) (
	fileAbsolutePath string,
	err error,
) {

	fileAbsolutePath, err = filepath.Abs(filePath)
	return
}
