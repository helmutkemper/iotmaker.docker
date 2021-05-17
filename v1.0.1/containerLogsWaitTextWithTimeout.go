package iotmakerdocker

import (
	"bytes"
	"errors"
	"github.com/docker/docker/api/types"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"
)

// ContainerLogsWaitTextWithTimeout (English):
//
// ContainerLogsWaitTextWithTimeout (Português):
func (el *DockerSystem) ContainerLogsWaitTextWithTimeout(
	id string,
	text string,
	timeout time.Duration,
	out io.Writer,
) (
	logContainer []byte,
	err error,
) {

	var wg sync.WaitGroup
	var reader io.ReadCloser
	var previousLog = make([]byte, 0)
	var cleanLog = make([]byte, 0)
	var ticker *time.Ticker
	var done = make(chan bool)

	if out != nil {
		log.New(out, "", 0)
	}

	wg.Add(1)
	go func(err *error, ticker *time.Ticker) {
		select {
		case <-done:
			ticker.Stop()
			return

		case <-ticker.C:
			ticker.Stop()
			*err = errors.New("timeout")
			wg.Done()
		}
	}(&err, ticker)

	go func(el *DockerSystem, err *error, reader *io.ReadCloser, previousLog, cleanLog, logContainer *[]byte, text *string, id string) {
		defer func() {
			done <- true
		}()

		for {
			*reader, *err = el.cli.ContainerLogs(el.ctx, id, types.ContainerLogsOptions{
				ShowStdout: true,
				ShowStderr: true,
				Timestamps: true,
				Follow:     false,
				Details:    false,
			})
			if *err != nil {
				return
			}

			*logContainer, *err = ioutil.ReadAll(*reader)
			if *err != nil {
				return
			}

			*cleanLog = bytes.Replace(*logContainer, *previousLog, []byte(""), -1)
			*previousLog = make([]byte, len(*logContainer))
			copy(*previousLog, *logContainer)

			//
			if out != nil && len(*cleanLog) != 0 {
				log.Printf("%s", *cleanLog)
			}

			if strings.Contains(string(*logContainer), *text) == true {
				wg.Done()
				return
			}

			time.Sleep(kWaitTextLoopSleep)
		}
	}(el, &err, &reader, &previousLog, &cleanLog, &logContainer, &text, id)

	ticker = time.NewTicker(timeout)
	wg.Wait()

	return
}
