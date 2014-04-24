package main

var CmdVersion = &Cmd{
	Name: "version",
	Desc: "Show current erk version.",
	Fn: func(args []string) error {
		println(VERSION)
		return nil
	},
}
