package utils

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Addr     string `yaml:"addr"`
}

var MyConfig *Config

func NewConfig() error {
	file, err := os.Open(".env/.yaml")
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
