package main

import (
	"fmt"
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
		color.Printf(INIT_MESSAGE, basePath)
		for _, issue := range IssueList {
			color.Printf(`[@{!b}%s@{|}] at @{!c}%s@{|}
  line %d - @{!w}%s

`, issue.Label, issue.FilePath, issue.Line, issue.Title)
		}
		if runWithRemote {
			color.Printf("updating remote issues...\n")
			if HasRemoteConfiguration() {
				// [todo] - supports other adapters(for bitbucket?)
				var adapter Adapter = &AdapterGithub{}
				adapter.Update()
			} else {
				return fmt.Errorf("Configuration file: can't load remote_config option successfully. please check your %s.", CONF_FILENAME)
			}
		}
		return nil
	},
}
