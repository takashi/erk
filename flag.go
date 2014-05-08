package main

import (
	"flag"
	"os"
)

var runWithRemote bool

func SetArgs() {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.BoolVar(&runWithRemote, "remote", false, "will updates remote repo issue(need remote_config option in conf file)")
	f.Parse(os.Args[1:])
	for 0 < f.NArg() {
		f.Parse(f.Args()[1:])
	}
}

func LoadArgs() {
	SetArgs()
	flag.Parse()
}
