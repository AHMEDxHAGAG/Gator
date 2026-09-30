// Package config
package config

import (
	"encoding/json"
	"os"
)

const configPath = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	path, err := getConfPath()
	if err != nil {
		return Config{}, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	var conf Config
	if err := json.Unmarshal(data, &conf); err != nil {
		return Config{}, err
	}
	return conf, nil
}

func (conf *Config) SetUser(userName string) error {
	conf.CurrentUserName = userName
	data, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	path, err := getConfPath()
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, data, 0o777); err != nil {
		return err
	}
	return nil
}

func getConfPath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homePath + "/" + configPath, nil
}
