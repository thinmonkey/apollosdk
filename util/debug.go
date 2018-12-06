package util

import (
	"fmt"
	"os"
	"strings"
)

var debugLog = true

func SetDebug(debug bool) {
	debugLog = debug
}

func DebugPrintf(format string, values ...interface{}) {
	if debugLog {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		fmt.Fprintf(os.Stderr, "[apollosdk-debug] "+format, values...)
		fmt.Fprint(os.Stderr, "\n")
	}
}
