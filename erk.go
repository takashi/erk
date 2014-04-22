package main

import "os"

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
