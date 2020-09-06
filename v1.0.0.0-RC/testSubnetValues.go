package iotmakerDocker

import (
	"errors"
	"regexp"
	"strconv"
)

func (el DockerSystem) testSubnetValues(
	subnet string,
) (
	err error,
) {

	var re *regexp.Regexp
	var subnetFieldAStr, subnetFieldBStr, subnetFieldCStr, subnetFieldDStr, subnetRangeStr string
	var subnetFieldA, subnetFieldB, subnetFieldC, subnetFieldD, subnetRange int64

	re = regexp.MustCompile(kSubnetExpressionRegular)
	subnetFieldAStr = re.ReplaceAllString(subnet, "${fieldA}")
	subnetFieldBStr = re.ReplaceAllString(subnet, "${fieldB}")
	subnetFieldCStr = re.ReplaceAllString(subnet, "${fieldC}")
	subnetFieldDStr = re.ReplaceAllString(subnet, "${fieldD}")
	subnetRangeStr = re.ReplaceAllString(subnet, "${range}")

	subnetFieldA, err = strconv.ParseInt(subnetFieldAStr, 10, 32)
	if err != nil {
		return
	}
	if subnetFieldA < 0 || subnetFieldA > 255 {
		err = errors.New("subnet fields must be between 0 and 255 and must be base 10")
		return
	}

	subnetFieldB, err = strconv.ParseInt(subnetFieldBStr, 10, 32)
	if err != nil {
		return
	}
	if subnetFieldB < 0 || subnetFieldB > 255 {
		err = errors.New("subnet fields must be between 0 and 255 and must be base 10")
		return
	}

	subnetFieldC, err = strconv.ParseInt(subnetFieldCStr, 10, 32)
	if err != nil {
		return
	}
	if subnetFieldC < 0 || subnetFieldC > 255 {
		err = errors.New("subnet fields must be between 0 and 255 and must be base 10")
		return
	}

	subnetFieldD, err = strconv.ParseInt(subnetFieldDStr, 10, 32)
	if err != nil {
		return
	}
	if subnetFieldD < 0 || subnetFieldD > 255 {
		err = errors.New("subnet fields must be between 0 and 255 and must be base 10")
		return
	}

	subnetRange, err = strconv.ParseInt(subnetRangeStr, 10, 32)
	if err != nil {
		return
	}
	if subnetRange < 0 || subnetRange > 255 {
		err = errors.New("subnet fields must be between 0 and 255 and must be base 10")
		return
	}

	return
}
