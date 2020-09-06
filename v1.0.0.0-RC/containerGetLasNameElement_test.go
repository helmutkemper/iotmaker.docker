package iotmakerDocker

import "fmt"

func ExampleContainerGetLasNameElement() {
	name := "/container_mongo"

	fmt.Printf("%v\n", ContainerGetLasNameElement(name))

	// Output:
	// container_mongo
}
