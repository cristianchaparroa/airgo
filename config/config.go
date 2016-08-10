package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var c Config

type Config struct {
	ClientID string `json:"clientId"`
}

// ReadConfig allow you to load files from file path
func ReadConfig(filepath string) (Config, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(fmt.Sprintf("failed to read config file: %v", err))
	}
	var c Config
	err = json.Unmarshal(file, &c)
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal config file: %v", err))
	}
	return c, err
}

func init() {
	var err error
	c, err = ReadConfig("config/config.json")
	if err != nil {
		panic(fmt.Sprintf("Error to load the config %v", err))
	}
}

func AppConfig() Config {
	return c
}
