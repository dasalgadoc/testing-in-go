package configs

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	DB  db  `yaml:"db"`
	Api Api `yaml:"Api"`
}

type db struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `host:"host"`
	Port     string `port:"port"`
	Database string `database:"database"`
}

type Api struct {
	Port string `yaml:"port"`
}

func LoadConfig(file string) (Config, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
