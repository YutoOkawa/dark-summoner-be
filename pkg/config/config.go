package config

import (
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Port             string `yaml:"port"`
	SummonerFilePath string `yaml:"summoner_file_path"`
	MonsterFilePath  string `yaml:"monster_file_path"`
}

func LoadConfigFile(filePath string) (*Config, error) {
	configBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(configBytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
