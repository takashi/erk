package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type RemoteConfig struct {
	Host     string
	Repo     string
	Api      string
	ApiToken string `json: "api_token"`
}

type Config struct {
	Dirs         []string
	Files        []string
	Remote       bool
	RemoteConfig RemoteConfig `json: "remote_config"`
	Ignore       []string
}

func LoadConfig() (Config, error) {
	var c Config
	// raise error if erkconfig.json is nor found in current directory.
	if !CheckFileExistence(CONF_FILENAME) {
		return c, fmt.Errorf("Configuration file: %s is not found. please run \"erk init\" first.", CONF_FILENAME)
	}

	// read json config file.
	body, err := ioutil.ReadFile(CONF_FILENAME)
	if err != nil {
		return c, err
	}

	// parse json and put them into Config struct.
	err = json.Unmarshal(body, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}
