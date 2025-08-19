package utils

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Host                   string `yaml:"host"`
	Port                   int    `yaml:"port"`
	User                   string `yaml:"user"`
	Password               string `yaml:"password"`
	DBName                 string `yaml:"dbname"`
	Addr                   string `yaml:"addr"`
	Kafka                  string `yaml:"kafka"`
	MaxIdleConns           int    `yaml:"MaxIdleConns"`
	MaxOpenConns           int    `yaml:"maxOpenConns"`
	ConnMaxLifetimeMinutes int    `yaml:"connMaxLifetimeMinutes"`
}

var MyConfig *Config

func NewConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	err = d.Decode(&MyConfig)
	if err != nil {
		return err
	}
	return nil
}
