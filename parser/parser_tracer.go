package parser

import (
	"fmt"
	"strings"
)

var traceLevel int = 0

const traceIdentPlaceholder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	fmt.Printf("%s%s\n", identLevel(), fs)
}

func incIdent() {
	traceLevel = traceLevel + 1
}

func decIdent() {
	traceLevel = traceLevel - 1
}

func trace(msg string, vars map[string]string) string {
	incIdent()
	paramsMsg := ""

	for k, v := range vars {
		paramsMsg += k + ": " + v + ","
	}

	tracePrint("BEGIN " + msg + "," + paramsMsg)
	return msg
}

func untrace(msg string) {
	tracePrint("END " + msg)
	decIdent()
}

func print(msg string, vars map[string]string) string {
	paramsMsg := ""

	for k, v := range vars {
		paramsMsg += k + ": " + v + ","
	}

	tracePrint("BEGIN " + msg + "," + paramsMsg)
	return msg
}
