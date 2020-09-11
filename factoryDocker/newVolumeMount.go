package factorydocker

import (
	"errors"
	"github.com/docker/docker/api/types/mount"
	"github.com/helmutkemper/iotmaker.docker/util"
	iotmakerdocker "github.com/helmutkemper/iotmaker.docker/v1.0.0.0-RC"
)

func NewVolumeMount(list []iotmakerdocker.Mount) (error, []mount.Mount) {
	var err error
	var found bool
	var fileAbsolutePath string
	var ret = make([]mount.Mount, 0)

	for _, v := range list {
		found = util.VerifyFileExists(v.Source)
		if found == false {
			return errors.New("source file not found"), nil
		}

		err, fileAbsolutePath = util.FileGetAbsolutePath(v.Source)
		if err != nil {
			return err, nil
		}

		ret = append(
			ret,
			mount.Mount{
				Type:   mount.Type(v.MountType.String()),
				Source: fileAbsolutePath,
				Target: v.Destination,
			},
		)
	}

	return nil, ret
}
