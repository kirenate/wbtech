package utils

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type Config struct {
	Host                    string        `yaml:"host"`
	Port                    int           `yaml:"port"`
	User                    string        `yaml:"user"`
	Password                string        `yaml:"password"`
	DBName                  string        `yaml:"dbname"`
	Addr                    string        `yaml:"addr"`
	Kafka                   string        `yaml:"kafka"`
	MaxIdleConns            int           `yaml:"MaxIdleConns"`
	MaxOpenConns            int           `yaml:"maxOpenConns"`
	ConnMaxLifetimeDuration string        `yaml:"connMaxLifetimeMinutes"`
	ConnMaxLifetime         time.Duration `yaml:"-"`
	RedisAddr               string        `yaml:"redisAddr"`
	RedisDB                 string        `yaml:"redisDB"`
	MaxRetries              int           `yaml:"maxRetries"`
	DialTimeoutDuration     string        `yaml:"dial_timeout"`
	DialTimeout             time.Duration `yaml:"-"`
	TimeoutDuration         string        `yaml:"timeout"`
	Timeout                 time.Duration `yaml:"-"`
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
	MyConfig.ConnMaxLifetime, err = time.ParseDuration(MyConfig.ConnMaxLifetimeDuration)
	if err != nil {
		return errors.Wrap(err, "failed to parse ConnMaxLifetimeDuration")
	}
	MyConfig.DialTimeout, err = time.ParseDuration(MyConfig.DialTimeoutDuration)
	if err != nil {
		return errors.Wrap(err, "failed to parse DialTimeoutDuration")
	}
	MyConfig.Timeout, err = time.ParseDuration(MyConfig.TimeoutDuration)
	if err != nil {
		return errors.Wrap(err, "failed to parse TimeoutDuration")
	}
	return nil
}
