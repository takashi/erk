package main

var CmdRun = &Cmd{
	Name: "run",
	Desc: "run erk to search inline issue.",
	Fn: func(args []string) error {
		basePath := "/Users/tak0303/src/beatrobo/PlugAir-Manage-API/"
		err := ParseFiles(basePath)
		if err != nil {
			return err
		}
		return nil
	},
}
