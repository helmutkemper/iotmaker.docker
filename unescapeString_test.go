package iotmakerDocker

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"html"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const (
	KAnsiReset = `\u001b[0m`

	KAnsiBold      = `\u001b[1m`
	KAnsiUnderline = `\u001b[4m`
	KAnsiReversed  = `\u001b[7m`

	KAnsiBlack  = `\u001b[30m`
	KAnsiRed    = `\u001b[31m`
	KAnsiGreen  = `\u001b[32m`
	KAnsiYellow = `\u001b[33m`
	KAnsiBlue   = `\u001b[34m`
	KAnsiPurple = `\u001b[35m`
	KAnsiCyan   = `\u001b[36m`
	KAnsiWhite  = `\u001b[37m`

	KAnsiBrightBlack  = `\u001b[90m`
	KAnsiBrightRed    = `\u001b[91m`
	KAnsiBrightGreen  = `\u001b[92m`
	KAnsiBrightYellow = `\u001b[93m`
	KAnsiBrightBlue   = `\u001b[94m`
	KAnsiBrightPurple = `\u001b[95m`
	KAnsiBrightCyan   = `\u001b[96m`
	KAnsiBrightWhite  = `\u001b[97m`

	KAnsiBgBlack  = `\u001b[40m`
	KAnsiBgRed    = `\u001b[41m`
	KAnsiBgGreen  = `\u001b[42m`
	KAnsiBgYellow = `\u001b[43m`
	KAnsiBgBlue   = `\u001b[44m`
	KAnsiBgPurple = `\u001b[45m`
	KAnsiBgCyan   = `\u001b[46m`
	KAnsiBgWhite  = `\u001b[47m`

	KAnsiBrightBgBlack  = `\u001b[100m`
	KAnsiBrightBgRed    = `\u001b[101m`
	KAnsiBrightBgGreen  = `\u001b[102m`
	KAnsiBrightBgYellow = `\u001b[103m`
	KAnsiBrightBgBlue   = `\u001b[104m`
	KAnsiBrightBgPurple = `\u001b[105m`
	KAnsiBrightBgCyan   = `\u001b[106m`
	KAnsiBrightBgWhite  = `\u001b[107m`

	KHtmlBold      = "Bold { font-weight: bold; }"
	KHtmlUnderline = "Underline { text-decoration: underline; }"
	KHtmlReversed  = "Reversed { filter: invert(100%); }"

	KHtmlBlack  = "Black { color: rgb(0,0,0); }"
	KHtmlRed    = "Red { color: rgb(255,0,0); }"
	KHtmlGreen  = "Green { color: rgb(0,255,0); }"
	KHtmlYellow = "Yellow { color: rgb(255,255,0); }"
	KHtmlBlue   = "Blue { color: rgb(0,0,255); }"
	KHtmlPurple = "Purple { color: rgb(128,0,128); }"
	KHtmlCyan   = "Cyan { color: rgb(0,255,255); }"
	KHtmlWhite  = "White { color: rgb(245,245,245); }"

	KHtmlBrightBlack  = "BrightBlack { color: rgb(105,105,105); }"
	KHtmlBrightRed    = "BrightRed { color: rgb(205,51,51); }"
	KHtmlBrightGreen  = "BrightGreen { color: rgb(127,255,0); }"
	KHtmlBrightYellow = "BrightYellow { color: rgb(255,255,127); }"
	KHtmlBrightBlue   = "BrightBlue { color: rgb(30,144,255); }"
	KHtmlBrightPurple = "BrightPurple { color: rgb(155,48,255); }"
	KHtmlBrightCyan   = "BrightCyan { color: rgb(0,238,238); }"
	KHtmlBrightWhite  = "BrightWhite { color: rgb(255,255,255); }"

	KHtmlBgBlack  = "BgBlack { background-color: rgb(0,0,0); color: rgb(255,255,255); }"
	KHtmlBgRed    = "BgRed { background-color: rgb(255,0,0); color: rgb(255,255,255); }"
	KHtmlBgGreen  = "BgGreen { background-color: rgb(0,255,0); color: rgb(255,255,255); }"
	KHtmlBgYellow = "BgYellow { background-color: rgb(255,255,0); color: rgb(255,255,255); }"
	KHtmlBgBlue   = "BgBlue { background-color: rgb(0,0,255); color: rgb(255,255,255); }"
	KHtmlBgPurple = "BgPurple { background-color: rgb(128,0,128); color: rgb(255,255,255); }"
	KHtmlBgCyan   = "BgCyan{ background-color: rgb(0,255,255); color: rgb(255,255,255); }"
	KHtmlBgWhite  = "BgWhite { background-color: rgb(245,245,245); color: rgb(255,255,255); }"

	KHtmlBgBrightBlack  = "BgBrightBlack { background-color: rgb(105,105,105); color: rgb(255,255,255); }"
	KHtmlBgBrightRed    = "BgBrightRed { background-color: rgb(205,51,51); color: rgb(255,255,255); }"
	KHtmlBgBrightGreen  = "BgBrightGreen { background-color: rgb(127,255,0); color: rgb(255,255,255); }"
	KHtmlBgBrightYellow = "BgBrightYellow { background-color: rgb(255,255,127); color: rgb(255,255,255); }"
	KHtmlBgBrightBlue   = "BgBrightBlue { background-color: rgb(30,144,255); color: rgb(255,255,255); }"
	KHtmlBgBrightPurple = "BgBrightPurple { background-color: rgb(155,48,255); color: rgb(255,255,255); }"
	KHtmlBgBrightCyan   = "BgBrightCyan { background-color: rgb(0,238,238); color: rgb(255,255,255); }"
	KHtmlBgBrightWhite  = "BgBrightWhite { background-color: rgb(255,255,255); color: rgb(255,255,255); }"
)

func ExampleUnescapeString() {
	var file []byte
	var err error
	file, err = ioutil.ReadFile("./a.txt")
	if err != nil {
		panic(err)
	}

	var resetCounter = 0

	cursorUp := regexp.MustCompile("\\\\u001b\\[\\d+A")
	file = cursorUp.ReplaceAllLiteral(file, []byte(""))
	cursorDown := regexp.MustCompile("\\\\u001b\\[\\d+B")
	file = cursorDown.ReplaceAllLiteral(file, []byte(""))
	cursorRight := regexp.MustCompile("\\\\u001b\\[\\d+C")
	file = cursorRight.ReplaceAllLiteral(file, []byte(""))
	cursorLeft := regexp.MustCompile("\\\\u001b\\[\\d+D")
	file = cursorLeft.ReplaceAllLiteral(file, []byte(""))
	cursorColor8AndCommands := regexp.MustCompile("\\\\u001b\\[\\d+m")
	file = cursorColor8AndCommands.ReplaceAllFunc(file, func(in []byte) (out []byte) {
		inData := strings.TrimSpace(string(in))

		switch inData {
		case KAnsiReset:
			for i := 0; i != resetCounter; i += 1 {
				out = append(out, []byte("</span>")...)
			}
			resetCounter = 0

		case KAnsiBold:
			resetCounter += 1
			out = []byte("<span class='Bold'>")
		case KAnsiUnderline:
			resetCounter += 1
			out = []byte("<span class='Underline'>")
		case KAnsiReversed:
			resetCounter += 1
			out = []byte("<span class='Reversed'>")

		case KAnsiBlack:
			resetCounter += 1
			out = []byte("<span class='Black'>")
		case KAnsiRed:
			resetCounter += 1
			out = []byte("<span class='Red'>")
		case KAnsiGreen:
			resetCounter += 1
			out = []byte("<span class='Green'>")
		case KAnsiYellow:
			resetCounter += 1
			out = []byte("<span class='Yellow'>")
		case KAnsiBlue:
			resetCounter += 1
			out = []byte("<span class='Blue'>")
		case KAnsiPurple:
			resetCounter += 1
			out = []byte("<span class='Purple'>")
		case KAnsiCyan:
			resetCounter += 1
			out = []byte("<span class='Cyan'>")
		case KAnsiWhite:
			resetCounter += 1
			out = []byte("<span class='White'>")

		case KAnsiBrightBlack:
			resetCounter += 1
			out = []byte("<span class='BrightBlack'>")
		case KAnsiBrightRed:
			resetCounter += 1
			out = []byte("<span class='BrightRed'>")
		case KAnsiBrightGreen:
			resetCounter += 1
			out = []byte("<span class='BrightGreen'>")
		case KAnsiBrightYellow:
			resetCounter += 1
			out = []byte("<span class='BrightYellow'>")
		case KAnsiBrightBlue:
			resetCounter += 1
			out = []byte("<span class='BrightBlue'>")
		case KAnsiBrightPurple:
			resetCounter += 1
			out = []byte("<span class='BrightPurple'>")
		case KAnsiBrightCyan:
			resetCounter += 1
			out = []byte("<span class='BrightCyan'>")
		case KAnsiBrightWhite:
			resetCounter += 1
			out = []byte("<span class='BrightWhite'>")

		case KAnsiBgBlack:
			resetCounter += 1
			out = []byte("<span class='BgBlack'>")
		case KAnsiBgRed:
			resetCounter += 1
			out = []byte("<span class='BgRed'>")
		case KAnsiBgGreen:
			resetCounter += 1
			out = []byte("<span class='BgGreen'>")
		case KAnsiBgYellow:
			resetCounter += 1
			out = []byte("<span class='BgYellow'>")
		case KAnsiBgBlue:
			resetCounter += 1
			out = []byte("<span class='BgBlue'>")
		case KAnsiBgPurple:
			resetCounter += 1
			out = []byte("<span class='BgPurple'>")
		case KAnsiBgCyan:
			resetCounter += 1
			out = []byte("<span class='BgCyan'>")
		case KAnsiBgWhite:
			resetCounter += 1
			out = []byte("<span class='BgWhite'>")

		case KAnsiBrightBgBlack:
			resetCounter += 1
			out = []byte("<span class='BrightBgBlack'>")
		case KAnsiBrightBgRed:
			resetCounter += 1
			out = []byte("<span class='BrightBgRed'>")
		case KAnsiBrightBgGreen:
			resetCounter += 1
			out = []byte("<span class='BrightBgGreen'>")
		case KAnsiBrightBgYellow:
			resetCounter += 1
			out = []byte("<span class='BrightBgYellow'>")
		case KAnsiBrightBgBlue:
			resetCounter += 1
			out = []byte("<span class='BrightBgBlue'>")
		case KAnsiBrightBgPurple:
			resetCounter += 1
			out = []byte("<span class='BrightBgPurple'>")
		case KAnsiBrightBgCyan:
			resetCounter += 1
			out = []byte("<span class='BrightBgCyan'>")
		case KAnsiBrightBgWhite:
			resetCounter += 1
			out = []byte("<span class='BrightBgWhite'>")
		}
		return
	})
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
		out = []byte(html.EscapeString(string(unicode)))
		return
	})

	file = bytes.ReplaceAll(file, []byte("\r\n"), []byte("<br>\n"))

	fmt.Printf("%s\n", file)

	//output:
	//
}
