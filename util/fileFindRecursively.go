package util

import (
	"errors"
	"os"
	"path/filepath"
)

func FileFindRecursively(fileName string) (filePath string, err error) {
	if _, err = os.Stat(fileName); os.IsNotExist(err) == false {
		filePath = fileName
		return
	}

	fileName = filepath.Base(fileName)
	err = filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == fileName {
				filePath = path
				return nil
			}

			return nil
		},
	)

	if filePath == "" {
		err = errors.New("file not found")
	}

	return
}
