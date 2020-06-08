package util

import (
	"errors"
	"fmt"
)

type IPv4Generator struct {
	a byte
	b byte
	c byte
	d byte
}

func (el *IPv4Generator) Init(a, b, c, d byte) {
	el.a = a
	el.b = b
	el.c = c
	el.d = d
}

func (el IPv4Generator) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", el.a, el.b, el.c, el.d)
}

func (el *IPv4Generator) Inc() error {
	var err error

	if el.a == 0 && el.b == 0 && el.c == 0 && el.d == 0 {
		el.a = 10
		el.b = 0
		el.c = 0
		el.d = 0
	}

	if el.a == 255 && el.b == 255 && el.c == 255 && el.d == 255 {
		err = errors.New("overflow")
	}

	if el.d < 255 {
		el.d += 1
		return err
	}
	el.d = 0

	if el.c < 255 {
		el.c += 1
		return err
	}
	el.c = 0

	if el.b < 255 {
		el.b += 1
		return err
	}
	el.b = 0

	if el.a < 255 {
		el.a += 1
		return err
	}
	el.a = 0

	return err
}
