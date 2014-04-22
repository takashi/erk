package main

import "fmt"

type Cmd struct {
	Name    string
	Desc    string
	Args    []string
	Aliases []string
	Fn      func(args []string) error
}

var CmdList []*Cmd

func init() {
	CmdList = []*Cmd{
		CmdVersion,
		CmdInit,
		CmdHelp,
	}
}

func CommandDispatch(args []string) error {
	var command *Cmd
	var commandName string
	var commandArgs = make([]string, 0)

	if len(args) < 2 {
		commandName = "help"
	} else {
		commandName = args[1]
	}

	for _, cmd := range CmdList {
		if cmd.Name == commandName {
			command = cmd
			break
		}
	}

	if command == nil {
		return fmt.Errorf("Command \"%s\" not found", commandName)
	}

	err := command.Fn(commandArgs)
	return err
}
