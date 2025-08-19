package utils

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type Config struct {
	Host                   string        `yaml:"host"`
	Port                   int           `yaml:"port"`
	User                   string        `yaml:"user"`
	Password               string        `yaml:"password"`
	DBName                 string        `yaml:"dbname"`
	Addr                   string        `yaml:"addr"`
	Kafka                  string        `yaml:"kafka"`
	MaxIdleConns           int           `yaml:"MaxIdleConns"`
	MaxOpenConns           int           `yaml:"maxOpenConns"`
	ConnMaxLifetimeMinutes string        `yaml:"connMaxLifetimeMinutes"`
	ConnMaxLifetime        time.Duration `yaml:"-"`
}

var MyConfig *Config

func NewConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "failed to open config file")
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	err = d.Decode(&MyConfig)
	if err != nil {
		return errors.Wrap(err, "failed to decode config")
	}
	MyConfig.ConnMaxLifetime, err = time.ParseDuration(MyConfig.ConnMaxLifetimeMinutes)
	if err != nil {
		return errors.Wrap(err, "failed to parse ConnMaxLifetimeMinutes")
	}
	return nil
}
