package log

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
}

func Infof(format string, args ...interface{}) {
	logger.Output(2, fmt.Sprintf("\033[1;34m[I]\033[0m "+format, args...))
}

func Warnf(format string, args ...interface{}) {
	logger.Output(2, fmt.Sprintf("\033[1;33m[W]\033[0m "+format, args...))
}

func Errorf(format string, args ...interface{}) {
	logger.Output(2, fmt.Sprintf("\033[1;31m[E]\033[0m "+format, args...))
}

func Fatalf(format string, args ...interface{}) {
	logger.Output(2, fmt.Sprintf("\033[1;35m[C]\033[0m "+format, args...))
	os.Exit(1)
}
