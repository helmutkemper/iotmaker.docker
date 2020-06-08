package iotmaker_docker

type VolumeMountType int

const (
	// TypeBind is the type for mounting host dir
	KVolumeMountTypeBind VolumeMountType = iota

	// TypeVolume is the type for remote storage volumes
	KVolumeMountTypeVolume

	// TypeTmpfs is the type for mounting tmpfs
	KVolumeMountTypeTmpfs

	// TypeNamedPipe is the type for mounting Windows named pipes
	KVolumeMountTypeNpipe
)

func (el VolumeMountType) String() string {
	return volumeMountTypes[el]
}

var volumeMountTypes = [...]string{
	"bind",
	"volume",
	"tmpfs",
	"npipe",
}
