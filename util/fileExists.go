package util

import "os"

func VerifyFileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true //!info.IsDir()
}
