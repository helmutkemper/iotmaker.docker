package iotmaker_docker

//  mountType:
//     KVolumeMountTypeBind - TypeBind is the type for mounting host dir
//     KVolumeMountTypeVolume - TypeVolume is the type for remote storage
//     volumes
//     KVolumeMountTypeTmpfs - TypeTmpfs is the type for mounting tmpfs
//     KVolumeMountTypeNpipe - TypeNamedPipe is the type for mounting
//     Windows named pipes
//  source: relative file/dir path in computer
//  destination: full path inside container
type Mount struct {
	MountType   VolumeMountType
	Source      string
	Destination string
}
