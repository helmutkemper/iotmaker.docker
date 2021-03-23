package iotmakerdocker

import (
	"errors"
	"regexp"
	"strconv"
)

func (el DockerSystem) testGatewayValues(
	gateway string,
) (
	a,
	b,
	c,
	d int,
	err error,
) {

	var re *regexp.Regexp
	var gatewayFieldAStr, gatewayFieldBStr, gatewayFieldCStr, gatewayFieldDStr string
	var gatewayFieldA, gatewayFieldB, gatewayFieldC, gatewayFieldD int64

	re = regexp.MustCompile(kGatewayExpressionRegular)
	gatewayFieldAStr = re.ReplaceAllString(gateway, "${fieldA}")
	gatewayFieldBStr = re.ReplaceAllString(gateway, "${fieldB}")
	gatewayFieldCStr = re.ReplaceAllString(gateway, "${fieldC}")
	gatewayFieldDStr = re.ReplaceAllString(gateway, "${fieldD}")

	gatewayFieldA, err = strconv.ParseInt(gatewayFieldAStr, 10, 32)
	if err != nil {
		return
	}
	if gatewayFieldA < 0 || gatewayFieldA > 255 {
		err = errors.New("gateway fields must be between 0 and 255 and must be base 10")
		return
	}

	gatewayFieldB, err = strconv.ParseInt(gatewayFieldBStr, 10, 32)
	if err != nil {
		return
	}
	if gatewayFieldB < 0 || gatewayFieldB > 255 {
		err = errors.New("gateway fields must be between 0 and 255 and must be base 10")
		return
	}

	gatewayFieldC, err = strconv.ParseInt(gatewayFieldCStr, 10, 32)
	if err != nil {
		return
	}
	if gatewayFieldC < 0 || gatewayFieldC > 255 {
		err = errors.New("gateway fields must be between 0 and 255 and must be base 10")
		return
	}

	gatewayFieldD, err = strconv.ParseInt(gatewayFieldDStr, 10, 32)
	if err != nil {
		return
	}
	if gatewayFieldD < 0 || gatewayFieldD > 255 {
		err = errors.New("gateway fields must be between 0 and 255 and must be base 10")
		return
	}

	a = int(gatewayFieldA)
	b = int(gatewayFieldB)
	c = int(gatewayFieldC)
	d = int(gatewayFieldD)

	return
}
