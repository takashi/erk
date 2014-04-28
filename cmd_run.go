package main

import "fmt"

var CmdRun = &Cmd{
	Name: "run",
	Desc: "run erk to search inline issue.",
	Fn: func(args []string) error {
		// basePath, err := os.Getwd()
		// if err != nil {
		// 	return err
		// }
		err := ParseFiles( /*basePath*/)
		if err != nil {
			return err
		}
		for _, issue := range IssueList {
			fmt.Printf(`[%s] at %s
  line %d - %s %s

`, issue.Label, issue.FileName, issue.Line, issue.Title, issue.FilePath)
		}
		return nil
	},
}
