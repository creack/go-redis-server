// +build debug

package redis

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Debugf retreives the stack info and output the given message to stdout
func Debugf(format string, a ...interface{}) {
	// Retrieve the stack infos
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<unknown>"
		line = -1
	} else {
		file = file[strings.LastIndex(file, "/")+1:]
	}

	fmt.Fprintf(os.Stderr, fmt.Sprintf("[%d] [debug] %s:%d %s\n", os.Getpid(), file, line, format), a...)
}
