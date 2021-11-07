package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type AppConfig struct {
	Token     string `json:"token"`
	BotPrefix string `json:"botprefix"`
}

func LoadConfig() (*AppConfig, error) {
	fmt.Println("Load config file...")

	var cfg AppConfig
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
