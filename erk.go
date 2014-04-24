package main

import "os"

func main() {
	var args = os.Args
	err := CommandDispatch(args)
	if err != nil {
		log_error("error %v", err)
		os.Exit(1)
	}
}
