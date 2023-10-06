package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	ConfigDirectoryPath = "data"
	ConfigFile          = "data.json"
)

type Config interface {
	Install() error
	Update() error
	Clean() error
	Diff() error
}

type ConfigApps struct {
	Apps map[string]App `json:"apps"`
}

type App struct {
	Paths []Path `json:"paths"`
}

type Path struct {
	Internal string `json:"internal"`
	External string `json:"External"`
}

func LoadConfig(path string) (*ConfigReal, *ConfigDemo, error) {
	configPath := filepath.Join(path, ConfigDirectoryPath, ConfigFile)
	bytes, err := os.ReadFile(configPath)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to read file %s: %w", configPath, err)
	}

	var cfgApps ConfigApps
	if err = json.Unmarshal(bytes, &cfgApps); err != nil {
		return nil, nil, fmt.Errorf("Failed to unmarshal: %w", err)
	}

	cfgReal := ConfigReal{ConfigApps: cfgApps}
	cfgDemo := ConfigDemo{ConfigApps: cfgApps}

	return &cfgReal, &cfgDemo, nil

}
