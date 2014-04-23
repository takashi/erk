package main

import (
	"fmt"
	"os"
)

const (
	defaultLogFormat = "erk: %s"
	errorLogFormat   = "\033[31merk: %s\033[0m"
)

func log_error(msg string, a ...interface{}) {
	msg = fmt.Sprintf(errorLogFormat, msg)
	log(msg, a...)
}

func log(msg string, a ...interface{}) {
	msg = fmt.Sprintf(msg, a...)
	fmt.Fprintf(os.Stderr, "%s\n", msg)
}
