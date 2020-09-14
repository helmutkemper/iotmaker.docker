package iotmakerdocker

import (
	"errors"
	"sync"
)

// wait image pull be completed
func (el *DockerSystem) ImageWaitPull(
	name string,
) (
	err error,
) {

	var wg sync.WaitGroup

	_, found := el.imageId[name]
	if found == false {
		return errors.New("image name not found in id list")
	}

	wg.Add(1)
	go func(el *DockerSystem, wg *sync.WaitGroup, name string) {

		for {
			id, err := el.ImageFindIdByName(name)
			if err != nil {
				panic(err)
			}

			if id != "" {
				wg.Done()
				return
			}
		}

	}(el, &wg, name)

	wg.Wait()

	return nil
}
