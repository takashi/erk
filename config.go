package main

import "fmt"

func LoadConfig() error {
	if !CheckFileExistence(CONF_FILENAME) {
		return fmt.Errorf("Configuration file: %s is not found. please run \"erk init\" first.", CONF_FILENAME)
	}
	return nil
}
