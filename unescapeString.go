package iotmakerDocker

import (
	"unicode/utf8"
)

func UtftoAscii(s string) []byte {
	t := make([]byte, utf8.RuneCountInString(s))
	i := 0
	for _, r := range s {
		t[i] = byte(r)
		i++
	}
	return t
}
