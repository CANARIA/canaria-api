package config

import (
	"github.com/kelseyhightower/envconfig"
)

const (
	USER     string = "root"
	PASSWORD string = "password"
	DB       string = "canaria"
	HOST     string = "mysql"
	PORT     string = "3306"
)

type Config struct {
	Env      string `envconfig:"ENV" default:"local"`
	Port     string `envconfig:"PORT" default:":5000"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
}

var conf = &Config{}

func init() {
	envconfig.Process("canaria", conf)
}

func GetEnv() string {
	return conf.Env
}

func GetPort() string {
	return conf.Port
}

func GetLogLevel() string {
	return conf.LogLevel
}
