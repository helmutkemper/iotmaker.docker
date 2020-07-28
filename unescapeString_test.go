package iotmakerDocker

import (
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io/ioutil"
	"regexp"
	"strconv"
)

const (
	AnsiReset = "\\u001B[0m"

	AnsiBlack  = "\\u001B[30m"
	AnsiRed    = "\\u001B[31m"
	AnsiGreen  = "\\u001B[32m"
	AnsiYellow = "\\u001B[33m"
	AnsiBlue   = "\\u001B[34m"
	AnsiPurple = "\\u001B[35m"
	AnsiCyan   = "\\u001B[36m"
	AnsiWhite  = "\\u001B[37m"

	AnsiBrightBlack  = "\\u001B[90m"
	AnsiBrightRed    = "\\u001B[91m"
	AnsiBrightGreen  = "\\u001B[92m"
	AnsiBrightYellow = "\\u001B[93m"
	AnsiBrightBlue   = "\\u001B[94m"
	AnsiBrightPurple = "\\u001B[95m"
	AnsiBrightCyan   = "\\u001B[96m"
	AnsiBrightWhite  = "\\u001B[97m"

	AnsiBgBlack  = "\\u001B[40m"
	AnsiBgRed    = "\\u001B[41m"
	AnsiBgGreen  = "\\u001B[42m"
	AnsiBgYellow = "\\u001B[43m"
	AnsiBgBlue   = "\\u001B[44m"
	AnsiBgPurple = "\\u001B[45m"
	AnsiBgCyan   = "\\u001B[46m"
	AnsiBgWhite  = "\\u001B[47m"

	AnsiBrightBgBlack  = "\\u001B[100m"
	AnsiBrightBgRed    = "\\u001B[101m"
	AnsiBrightBgGreen  = "\\u001B[102m"
	AnsiBrightBgYellow = "\\u001B[103m"
	AnsiBrightBgBlue   = "\\u001B[104m"
	AnsiBrightBgPurple = "\\u001B[105m"
	AnsiBrightBgCyan   = "\\u001B[106m"
	AnsiBrightBgWhite  = "\\u001B[107m"

	HtmlBlack  = "Black { color: rgb(0,0,0); }"
	HtmlRed    = "Red { color: rgb(255,0,0); }"
	HtmlGreen  = "Green { color: rgb(0,255,0); }"
	HtmlYellow = "Yellow { color: rgb(255,255,0); }"
	HtmlBlue   = "Blue { color: rgb(0,0,255); }"
	HtmlPurple = "Purple { color: rgb(128,0,128); }"
	HtmlCyan   = "Cyan { color: rgb(0,255,255); }"
	HtmlWhite  = "White { color: rgb(245,245,245); }"

	HtmlBrightBlack  = "BrightBlack { color: rgb(105,105,105); }"
	HtmlBrightRed    = "BrightRed { color: rgb(205,51,51); }"
	HtmlBrightGreen  = "BrightGreen { color: rgb(127,255,0); }"
	HtmlBrightYellow = "BrightYellow { color: rgb(255,255,127); }"
	HtmlBrightBlue   = "BrightBlue { color: rgb(30,144,255); }"
	HtmlBrightPurple = "BrightPurple { color: rgb(155,48,255); }"
	HtmlBrightCyan   = "BrightCyan { color: rgb(0,238,238); }"
	HtmlBrightWhite  = "BrightWhite { color: rgb(255,255,255); }"

	HtmlBgBlack  = "BgBlack { background-color: rgb(0,0,0); color: rgb(255,255,255); }"
	HtmlBgRed    = "BgRed { background-color: rgb(255,0,0); color: rgb(255,255,255); }"
	HtmlBgGreen  = "BgGreen { background-color: rgb(0,255,0); color: rgb(255,255,255); }"
	HtmlBgYellow = "BgYellow { background-color: rgb(255,255,0); color: rgb(255,255,255); }"
	HtmlBgBlue   = "BgBlue { background-color: rgb(0,0,255); color: rgb(255,255,255); }"
	HtmlBgPurple = "BgPurple { background-color: rgb(128,0,128); color: rgb(255,255,255); }"
	HtmlBgCyan   = "BgCyan{ background-color: rgb(0,255,255); color: rgb(255,255,255); }"
	HtmlBgWhite  = "BgWhite { background-color: rgb(245,245,245); color: rgb(255,255,255); }"

	HtmlBgBrightBlack  = "BgBrightBlack { background-color: rgb(105,105,105); color: rgb(255,255,255); }"
	HtmlBgBrightRed    = "BgBrightRed { background-color: rgb(205,51,51); color: rgb(255,255,255); }"
	HtmlBgBrightGreen  = "BgBrightGreen { background-color: rgb(127,255,0); color: rgb(255,255,255); }"
	HtmlBgBrightYellow = "BgBrightYellow { background-color: rgb(255,255,127); color: rgb(255,255,255); }"
	HtmlBgBrightBlue   = "BgBrightBlue { background-color: rgb(30,144,255); color: rgb(255,255,255); }"
	HtmlBgBrightPurple = "BgBrightPurple { background-color: rgb(155,48,255); color: rgb(255,255,255); }"
	HtmlBgBrightCyan   = "BgBrightCyan { background-color: rgb(0,238,238); color: rgb(255,255,255); }"
	HtmlBgBrightWhite  = "BgBrightWhite { background-color: rgb(255,255,255); color: rgb(255,255,255); }"
)

func ExampleUnescapeString() {
	var file []byte
	var err error
	file, err = ioutil.ReadFile("./a.txt")
	if err != nil {
		panic(err)
	}

	cursorUp := regexp.MustCompile("\\\\u001b\\[\\d+A")
	file = cursorUp.ReplaceAllLiteral(file, []byte(""))
	cursorDown := regexp.MustCompile("\\\\u001b\\[\\d+B")
	file = cursorDown.ReplaceAllLiteral(file, []byte(""))
	cursorRight := regexp.MustCompile("\\\\u001b\\[\\d+C")
	file = cursorRight.ReplaceAllLiteral(file, []byte(""))
	cursorLeft := regexp.MustCompile("\\\\u001b\\[\\d+D")
	file = cursorLeft.ReplaceAllLiteral(file, []byte(""))
	//cursorColor8 := regexp.MustCompile("\\\\u001b\\[\\d+m")
	//file = cursorColor8.ReplaceAllFunc(file, func(in []byte) (out []byte) {
	//	in = in[7:len(in)-1]
	//	return
	//})
	//cursorColor16 := regexp.MustCompile("\\\\u001b\\[\\d+;\\d+m")
	//file = cursorColor16.ReplaceAllLiteral(file, []byte(""))
	//cursorColor256 := regexp.MustCompile("\\\\u001b\\[\\d+;\\d+;\\d+m")
	//file = cursorColor256.ReplaceAllLiteral(file, []byte(""))
	//cursorNav := regexp.MustCompile("\\\\u001b\\[\\d+D")
	//file = cursorNav.ReplaceAllLiteral(file, []byte(""))

	//carriageReturn := regexp.MustCompile("\\r")
	//file = carriageReturn.ReplaceAllLiteral(file, []byte(""))

	unicode := regexp.MustCompile("\\\\u.{4}")
	file = unicode.ReplaceAllFunc(file, func(in []byte) (out []byte) {
		//remove \u e deixa um número binário
		in = in[2:]
		number, _ := strconv.ParseInt(string(in), 16, 16)
		unicode := charmap.ISO8859_1.DecodeByte(byte(number))
		out = []byte(string(unicode))
		return
	})

	fmt.Printf("%s\n", file)

	//output:
	//
}
