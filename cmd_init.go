package main

import (
	"fmt"
	"io/ioutil"
)

var CmdInit = &Cmd{
	Name: "init",
	Desc: "Creates erkconf.json file to configurate.",
	Fn: func(args []string) error {
		// check if configuration file already exists.
		if CheckFileExistence(CONF_FILENAME) {
			return fmt.Errorf("%s is already exists... ;)", CONF_FILENAME)
		}
		// creates new configuration file.
		err := ioutil.WriteFile(CONF_FILENAME, []byte(INIT_FILE_MESSAGE), 0666)
		if err != nil {
			return err
		}
		fmt.Println("erk initialized!")
		return nil
	},
}
