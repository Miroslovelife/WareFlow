package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Уникальные типы для конфигурации
type MongoURI string
type DBName string

type Config struct {
	MongoURI MongoURI `yaml:"mongo_uri"`
	DBName   DBName   `yaml:"db_name"`
}

// LoadConfig загружает конфигурацию из YAML-файла
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return &cfg, nil
}
