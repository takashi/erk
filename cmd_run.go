package main

import (
	"github.com/wsxiaoys/terminal/color"
	"os"
)

var CmdRun = &Cmd{
	Name: "run",
	Desc: "run erk to search inline issue.",
	Fn: func(args []string) error {
		err := ParseFiles( /*basePath*/)
		if err != nil {
			return err
		}

		basePath, err := os.Getwd()
		if err != nil {
			return err
		}
		color.Printf(`--------------------------
@{g}erk@{|} - inline issue manager
Run in: %s
--------------------------

`, basePath)
		for _, issue := range IssueList {
			color.Printf(`[@{!b}%s@{|}] at @{!c}%s@{|}
  line %d - @{!w}%s

`, issue.Label, issue.FilePath, issue.Line, issue.Title)
		}
		if HasRemoteConfiguration() {
			var adapter Adapter = &AdapterGithub{}
			adapter.Update()
		}
		return nil
	},
}
