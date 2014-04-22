package main

var CmdVersion = &Cmd{
	Name: "version",
	Fn: func(args []string) error {
		println(VERSION)
		return nil
	},
}
