package main

import (
	"fmt"
	"os"
)

func LoadConfig() error {
	if !CheckFileExistence(CONF_FILENAME) {
		return fmt.Errorf("Configuration file: %s is not found. please run \"erk init\" first.", CONF_FILENAME)
	}
	return nil
}

func exitIfError(err error) {
	if err != nil {
		log_error("error %v", err)
		os.Exit(1)
	}
}

func main() {
	var args = os.Args
	// err := LoadConfig()
	// exitIfError(err)
	err := CommandDispatch(args)
	exitIfError(err)
}
