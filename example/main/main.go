package main

import (
	"fmt"
	"github.com/helmutkemper/iotmaker.docker.util.whaleAquarium/workInProgress/util"
)

func main() {
	err, ins := util.NetworkFindTypeBridge()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", ins)
}
