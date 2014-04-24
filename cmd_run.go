package main

import "os"

var CmdRun = &Cmd{
	Name: "run",
	Desc: "run erk to search inline issue.",
	Fn: func(args []string) error {
		basePath, err := os.Getwd()
		if err != nil {
			return err
		}
		err = ParseFiles(basePath)
		if err != nil {
			return err
		}
		return nil
	},
}
