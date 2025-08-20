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
}

type RedisCfg struct {
	RedisAddr           string        `yaml:"redisAddr"`
	RedisDB             int           `yaml:"redisDB"`
	MaxRetries          int           `yaml:"maxRetries"`
	DialTimeoutDuration string        `yaml:"dialTimeoutDuration"`
	DialTimeout         time.Duration `yaml:"-"`
	TimeoutDuration     string        `yaml:"timeoutDuration"`
	Timeout             time.Duration `yaml:"-"`
	TTLDuration         string        `yaml:"ttl"`
	TTL                 time.Duration `yaml:"-"`
}

var RedisCFG *RedisCfg

func NewRedisConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "failed to open redis config file")
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	err = d.Decode(&RedisCFG)
	if err != nil {
		return errors.Wrap(err, "failed to decode redis config")
	}
	RedisCFG.DialTimeout, err = time.ParseDuration(RedisCFG.DialTimeoutDuration)
	if err != nil {
		return errors.Wrap(err, "failed to parse DialTimeoutDuration")
	}
	RedisCFG.Timeout, err = time.ParseDuration(RedisCFG.TimeoutDuration)
	if err != nil {
		return errors.Wrap(err, "failed to parse TimeoutDuration")
	}
	RedisCFG.TTL, err = time.ParseDuration(RedisCFG.TTLDuration)
	if err != nil {
		return errors.Wrap(err, "failed to parse TTLDuration")
	}
	return nil
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
	return nil
}
