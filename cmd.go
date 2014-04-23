package main

import "fmt"

type Cmd struct {
	Name    string
	Desc    string
	Args    []string
	Aliases []string
	Fn      func(args []string, config Config) error
}

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
	var commandName string
	var commandArgs = make([]string, 0)
	var config Config

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
	} else if command.Name != "init" &&
		command.Name != "help" &&
		commandName != "version" {
		config, err := LoadConfig()
		if err != nil {
			return err
		}
		err = command.Fn(commandArgs, config)
		return err
	}
	err := command.Fn(commandArgs, config)
	return err
}
