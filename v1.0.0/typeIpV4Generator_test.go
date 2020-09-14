package iotmakerdocker

import (
	"testing"
)

func TestIPv4Generator_InitWithStringAndAllowMaxValue(t *testing.T) {
	var err error

	ipGenerator := IPv4Generator{}
	err = ipGenerator.InitWithStringAndAllowMaxValue("10.0.0.2/4")
	if err != nil {
		t.Fail()
		return
	}
	ip := ipGenerator.String()
	if ip != "10.0.0.2/4" {
		t.Fail()
		return
	}

	err = ipGenerator.InitWithStringAndAllowMaxValue("10.0.0.1/4")
	if err != nil {
		t.Fail()
		return
	}
	ip = ipGenerator.String()
	if ip != "10.0.0.2/4" {
		t.Fail()
		return
	}

	err = ipGenerator.InitWithStringAndAllowMaxValue("10.0.0.3/4")
	if err != nil {
		t.Fail()
		return
	}
	ip = ipGenerator.String()
	if ip != "10.0.0.3/4" {
		t.Fail()
		return
	}

	err = ipGenerator.InitWithStringAndAllowMaxValue("10.0.1.0/16")
	if err != nil {
		t.Fail()
		return
	}
	ip = ipGenerator.String()
	if ip != "10.0.1.0/16" {
		t.Fail()
		return
	}

	err = ipGenerator.InitWithStringAndAllowMaxValue("10.1.0.0/16")
	if err != nil && err.Error() != "max allowed ip is 010.000.255.255" {
		t.Fail()
		return
	}

	err = ipGenerator.InitWithStringAndAllowMaxValue("10.1.0.0/32")
	if err != nil {
		t.Fail()
		return
	}
	ip = ipGenerator.String()
	if ip != "10.1.0.0/32" {
		t.Fail()
		return
	}
}

func TestIPv4Generator_Inc(t *testing.T) {
	var err error

	ipGenerator := IPv4Generator{}
	err = ipGenerator.InitWithString("10.0.0.0/4")
	if err != nil {
		t.Fail()
		return
	}

	err = ipGenerator.Inc()
	if err != nil {
		t.Fail()
		return
	}
	ip := ipGenerator.String()
	if ip != "10.0.0.1/4" {
		t.Fail()
		return
	}

	for i := 0; i != 14; i += 1 {
		err = ipGenerator.Inc()
		if err != nil {
			t.Fail()
			return
		}
	}

	ip = ipGenerator.String()
	if ip != "10.0.0.15/4" {
		t.Fail()
		return
	}

	err = ipGenerator.Inc()
	if err != nil && err.Error() != "max allowed ip is 010.000.000.015" {
		t.Fail()
		return
	}
}

func TestCalcMaxValue(t *testing.T) {
	ipGenerator := IPv4Generator{}

	expectedResponse := []string{
		"000.000.001",
		"000.000.003",
		"000.000.007",
		"000.000.015",
		"000.000.031",
		"000.000.063",
		"000.000.127",
		"000.000.255",
		"000.001.255",
		"000.003.255",
		"000.007.255",
		"000.015.255",
		"000.031.255",
		"000.063.255",
		"000.127.255",
		"000.255.255",
		"001.255.255",
		"003.255.255",
		"007.255.255",
		"015.255.255",
		"031.255.255",
		"063.255.255",
		"127.255.255",
		"255.255.255",
	}
	calculeResponse := make([]string, 0)
	for i := 1; i != 25; i += 1 {
		calculeResponse = append(calculeResponse, ipGenerator.cidrToHumanNotation(i))
	}

	for i := 0; i != len(expectedResponse); i += 1 {
		if expectedResponse[i] != calculeResponse[i] {
			t.Fail()
		}
	}
}

func TestIPv4Generator_ParserIP(t *testing.T) {
	var err error
	ipGenerator := IPv4Generator{}

	err = ipGenerator.InitWithString("10.0.0.1")
	if err != nil {
		t.Fail()
		return
	}

	ip := ipGenerator.String()
	if ip != "10.0.0.1" {
		t.Fail()
		return
	}

	err = ipGenerator.InitWithString("010.000.000.001")
	if err != nil {
		t.Fail()
		return
	}
	ip = ipGenerator.String()
	if ip != "10.0.0.1" {
		t.Fail()
		return
	}

	err = ipGenerator.InitWithString("10.0.0.1/16")
	if err != nil {
		t.Fail()
		return
	}

	ip = ipGenerator.String()
	if ip != "10.0.0.1/16" {
		t.Fail()
		return
	}

	err = ipGenerator.InitWithString("010.000.000.001/16")
	if err != nil {
		t.Fail()
		return
	}
	ip = ipGenerator.String()
	if ip != "10.0.0.1/16" {
		t.Fail()
		return
	}

	err = ipGenerator.InitWithString("010.000.000.256/16")
	if err != nil && err.Error() != "max theoretical allowed value is 255.255.255.255/128" {
		t.Fail()
		return
	}

	err = ipGenerator.InitWithString("10.0.0.128/4")
	if err != nil && err.Error() != "max allowed ip is 010.000.000.015" {
		t.Fail()
		return
	}

}

func TestVerify(t *testing.T) {
	ipGenerator := IPv4Generator{}

	expectedResponse := []string{
		"max allowed ip is 010.000.000.001",
		"max allowed ip is 010.000.000.003",
		"max allowed ip is 010.000.000.007",
		"max allowed ip is 010.000.000.015",
		"max allowed ip is 010.000.000.031",
		"max allowed ip is 010.000.000.063",
		"max allowed ip is 010.000.000.127",
		"max allowed ip is 010.000.000.255",
		"max allowed ip is 010.000.001.255",
		"max allowed ip is 010.000.003.255",
		"max allowed ip is 010.000.007.255",
		"max allowed ip is 010.000.015.255",
		"max allowed ip is 010.000.031.255",
		"max allowed ip is 010.000.063.255",
		"max allowed ip is 010.000.127.255",
		"max allowed ip is 010.000.255.255",
		"max allowed ip is 010.001.255.255",
		"max allowed ip is 010.003.255.255",
		"max allowed ip is 010.007.255.255",
		"max allowed ip is 010.015.255.255",
		"max allowed ip is 010.031.255.255",
		"max allowed ip is 010.063.255.255",
		"max allowed ip is 010.127.255.255",
	}
	calculeResponse := make([]string, 0)
	for i := 1; i != 24; i += 1 {
		err := ipGenerator.verify(10, 255, 255, 255, i)
		if err == nil {
			t.Fail()
			return
		}
		calculeResponse = append(calculeResponse, err.Error())
	}

	for i := 0; i != len(expectedResponse); i += 1 {
		if expectedResponse[i] != calculeResponse[i] {
			t.Fail()
		}
	}
}
