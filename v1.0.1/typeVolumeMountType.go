package iotmakerdocker

type VolumeMountType int

const (
	// TypeBind is the type for mounting host dir (real folder inside computer where this code work)
	KVolumeMountTypeBindString = "bind"

	// TypeVolume is the type for remote storage volumes
	KVolumeMountTypeVolumeString = "volume"

	// TypeTmpfs is the type for mounting tmpfs
	KVolumeMountTypeTmpfsString = "tmpfs"

	// TypeNamedPipe is the type for mounting Windows named pipes
	KVolumeMountTypeNpipeString = "npipe"
)

const (
	// TypeBind is the type for mounting host dir (real folder inside computer where this code work)
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
