package iotmakerDocker

type ContainerPullProgress struct {
	Stream                     string                      `json:"stream"`
	Status                     string                      `json:"status"`
	ProgressDetail             ContainerPullProgressDetail `json:"progressDetail"`
	ID                         string                      `json:"id"`
	SysStatus                  ContainerPullStatus         `json:"-"`
	ImageName                  string
	successfullyBuildContainer bool
	successfullyBuildImage     bool
}
