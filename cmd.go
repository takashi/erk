package main

import "fmt"

type Cmd struct {
	Name    string
	Desc    string
	Args    []string
	Aliases []string
	Fn      func(args []string) error
}

var config Config

var CmdList []*Cmd

func init() {
	CmdList = []*Cmd{
		CmdVersion,
		CmdInit,
		CmdHelp,
		CmdRun,
	}
}

func CommandDispatch(args []string) error {
	var command *Cmd
	var name string
	var commandArgs = make([]string, 0)
	var err error

	if len(args) < 2 {
		name = "help"
	} else {
		name = args[1]
	}

	for _, cmd := range CmdList {
		if cmd.Name == name {
			command = cmd
			break
		}
	}

	if command == nil {
		return fmt.Errorf("Command \"%s\" not found", name)
	} else if name != "init" &&
		name != "help" &&
		name != "version" {
		config, err = LoadConfig()
		if err != nil {
			return err
		}
		err = command.Fn(commandArgs)
		return err
	}
	err = command.Fn(commandArgs)
	return err
}
