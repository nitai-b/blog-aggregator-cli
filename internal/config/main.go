// Package config
package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func gatorConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".gatorconfig.json"), nil
}

func write(cfg Config) error {
	configPath, err := gatorConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(configPath, data, 0o644); err != nil {
		return err
	}

	return nil
}

func Read() (Config, error) {
	var config Config

	path, err := gatorConfigPath()
	if err != nil {
		return config, err
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	if err := json.Unmarshal(bytes, &config); err != nil {
		return config, err
	}

	return config, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	if err := write(*c); err != nil {
		return err
	}
	return nil
}
