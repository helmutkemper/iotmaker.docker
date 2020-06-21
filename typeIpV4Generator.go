package iotmakerDocker

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const (
	kGatewayExpressionRegular = `^(?P<fieldA>[0-9]{0,3})\.(?P<fieldB>[0-9]{0,3})\.(?P<fieldC>[0-9]{0,3})\.(?P<fieldD>[0-9]{0,3})$`
	kSubnetExpressionRegular  = `^(?P<fieldA>[0-9]{0,3})\.(?P<fieldB>[0-9]{0,3})\.(?P<fieldC>[0-9]{0,3})\.(?P<fieldD>[0-9]{0,3})/(?P<range>[0-9]{0,3})$`
)

type calcBase256 struct {
	M int
	D int
}

type IPv4Generator struct {
	a          byte
	b          byte
	c          byte
	d          byte
	cidrPrefix byte
}

func (el *IPv4Generator) InitWithString(IP string) (err error) {
	var re1 = regexp.MustCompile(`^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$`)
	var re2 = regexp.MustCompile(`^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}/[0-9]{1,3}$`)
	var a, b, c, d, cidrPrefix int

	if re1.MatchString(IP) == true || re2.MatchString(IP) == true {
		err, a, b, c, d, cidrPrefix = el.split(IP)
		err = el.verify(a, b, c, d, cidrPrefix)
		if err != nil {
			return
		}

		el.a = byte(a)
		el.b = byte(b)
		el.c = byte(c)
		el.d = byte(d)
		el.cidrPrefix = byte(cidrPrefix)

		return
	}

	err = errors.New("IP parser error")
	return
}

// english: simplifies the initialization of the network function when you have an ip
// list, continuing from the highest value found
//
// português: simplifica a inicialização da função de rede quando você tem uma lista de
// ip, continuando a partir do valor mais alto encontrado
//
//   Example:
//     ipGen := IPv4Generator{}
//     for _, v := rage []string{"10.0.0.1/16", "10.0.0.2/16", "10.0.0.3/16"} {
//       ipGen.InitWithStringAndAllowMaxValue( v )
//     }
func (el *IPv4Generator) InitWithStringAndAllowMaxValue(IP string) (err error) {
	var re1 = regexp.MustCompile(`^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$`)
	var re2 = regexp.MustCompile(`^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}/[0-9]{1,3}$`)
	var a, b, c, d, cidrPrefix int

	if re1.MatchString(IP) == true || re2.MatchString(IP) == true {
		err, a, b, c, d, cidrPrefix = el.split(IP)
		err = el.verify(a, b, c, d, cidrPrefix)
		if err != nil {
			return
		}

		ruleA := byte(a) > el.a
		ruleB := byte(a) == el.a && byte(b) > el.b
		ruleC := byte(a) == el.a && byte(b) == el.b && byte(c) > el.c
		ruleD := byte(a) == el.a && byte(b) == el.b && byte(c) == el.c && byte(d) > el.d

		if ruleA == true || ruleB == true || ruleC == true || ruleD == true {
			el.a = byte(a)
			el.b = byte(b)
			el.c = byte(c)
			el.d = byte(d)
			el.cidrPrefix = byte(cidrPrefix)
		}

		err = el.Inc()
		return
	}

	err = errors.New("IP parser error")
	return
}

func (el IPv4Generator) split(ip string) (err error, a, b, c, d, CIDRPrefix int) {
	var tmpA, tmpB, tmpC, tmpD, tmpE string
	var aInt64, bInt64, cInt64, dInt64, CIDRPrefixInt64 int64

	list := strings.Split(ip, ".")
	tmpA = list[0]
	tmpB = list[1]
	tmpC = list[2]
	tmpD = list[3]

	tmpEArr := strings.Split(tmpD, "/")
	tmpD = tmpEArr[0]

	if len(tmpEArr) == 2 {
		tmpE = tmpEArr[1]
	} else {
		tmpE = ""
	}

	aInt64, err = strconv.ParseInt(tmpA, 10, 32)
	if err != nil {
		return
	}

	bInt64, err = strconv.ParseInt(tmpB, 10, 32)
	if err != nil {
		return
	}

	cInt64, err = strconv.ParseInt(tmpC, 10, 32)
	if err != nil {
		return
	}

	dInt64, err = strconv.ParseInt(tmpD, 10, 32)
	if err != nil {
		return
	}

	if tmpE == "" {
		CIDRPrefixInt64 = 0
	} else {
		CIDRPrefixInt64, err = strconv.ParseInt(tmpE, 10, 32)
	}

	a = int(aInt64)
	b = int(bInt64)
	c = int(cInt64)
	d = int(dInt64)
	CIDRPrefix = int(CIDRPrefixInt64)

	return
}

func (el IPv4Generator) toHumanNotation(resp *[]calcBase256) string {
	var toOut = make([]string, 0)
	toOut = append(toOut, fmt.Sprintf("%03d", (*resp)[len(*resp)-1].D))
	for i := len(*resp) - 1; i != -1; i -= 1 {
		toOut = append(toOut, fmt.Sprintf("%03d", (*resp)[i].M))
	}

	if len(toOut) > 3 {
		toOut = toOut[len(toOut)-3:]
	}

	for i := len(toOut); i <= 2; i += 1 {
		toOut = append([]string{fmt.Sprintf("%03d", 0)}, toOut...)
	}

	return strings.Join(toOut, ".")
}

func (el IPv4Generator) cidrToHumanNotation(cidr int) (humanNotation string) {
	resp := make([]calcBase256, 0)
	value := el.cidrPrefixToDecimal(cidr)
	el.calcMaxValueSupport(value, &resp)

	return el.toHumanNotation(&resp)
}

func (el IPv4Generator) calcMaxValueSupport(value int, resp *[]calcBase256) {
	if len(*resp) == 0 {
		*resp = make([]calcBase256, 0)
	}

	m := value % 256
	d := value / 256

	*resp = append(*resp, calcBase256{M: m, D: d})

	if m != 0 {
		el.calcMaxValueSupport(d, resp)
	}
}

func (el IPv4Generator) cidrPrefixToDecimal(cidr int) (cidrDecimal int) {
	return int(math.Pow(2.0, float64(cidr)) - 1)
}

func (el IPv4Generator) verify(a, b, c, d, CIDRPrefix int) (err error) {
	return nil
	if a > 255 || b > 255 || c > 255 || d > 255 || CIDRPrefix > 128 {
		err = errors.New("max theoretical allowed value is 255.255.255.255/128")
	}

	if CIDRPrefix != 0 {
		tmpB := float64(b) * math.Pow(256.0, 2.0)
		tmpC := float64(c) * math.Pow(256.0, 1.0)
		tmpD := float64(d) * math.Pow(256.0, 0.0)

		if int(tmpB+tmpC+tmpD) > el.cidrPrefixToDecimal(CIDRPrefix) {
			max := el.cidrToHumanNotation(CIDRPrefix)
			err = errors.New(fmt.Sprintf("max allowed ip is %03d.%v", a, max))
		}
	}

	return
}

func (el *IPv4Generator) InitWithCIDRPrefix(a, b, c, d, cidr byte) {
	el.a = a
	el.b = b
	el.c = c
	el.d = d
	el.cidrPrefix = cidr
}

func (el *IPv4Generator) Inc() (err error) {

	if el.a == 0 && el.b == 0 && el.c == 0 && el.d == 0 {
		el.a = 10
		el.b = 0
		el.c = 0
		el.d = 0
		el.cidrPrefix = 16
	}

	if el.d < 255 {
		el.d += 1
		return el.verify(int(el.a), int(el.b), int(el.c), int(el.d), int(el.cidrPrefix))
	}
	el.d = 0

	if el.c < 255 {
		el.c += 1
		return el.verify(int(el.a), int(el.b), int(el.c), int(el.d), int(el.cidrPrefix))
	}
	el.c = 0

	if el.b < 255 {
		el.b += 1
		return el.verify(int(el.a), int(el.b), int(el.c), int(el.d), int(el.cidrPrefix))
	}
	el.b = 0

	if el.a < 255 {
		el.a += 1
		return el.verify(int(el.a), int(el.b), int(el.c), int(el.d), int(el.cidrPrefix))
	}
	el.a = 0

	return el.verify(int(el.a), int(el.b), int(el.c), int(el.d), int(el.cidrPrefix))
}

func (el *IPv4Generator) Init(a, b, c, d byte) {
	el.a = a
	el.b = b
	el.c = c
	el.d = d
}

func (el IPv4Generator) String() string {
	if el.cidrPrefix == 0 {
		return fmt.Sprintf("%v.%v.%v.%v", el.a, el.b, el.c, el.d)
	}

	return fmt.Sprintf("%v.%v.%v.%v/%v", el.a, el.b, el.c, el.d, el.cidrPrefix)
}
