package config

import (
	"github.com/kelseyhightower/envconfig"
)

// TODO: yamlに移動したい
const (
	USER     string = "root"
	PASSWORD string = "password"
	DB       string = "canaria"
	// HOST     string = "mysql"
	HOST string = "127.0.0.1"
	PORT string = "3306"

	SALT string = "@#$%"

	UserInfo string = "UserInfo"
	IsDebug  string = "IS_DEBUG"
	Env      string = "API_ENV"
)

type Config struct {
	Env         string `envconfig:"ENV" default:"local"`
	MailAddress string `envconfig:"MAILADDRESS" default:"canaria-dev@gmail.com"`
	LogLevel    string `envconfig:"LOG_LEVEL" default:"info"`
	Host        string `envconfig:"HOST" default:"localhost"`
	Port        string `envconfig:"PORT" default:":5000"`
}

var conf = &Config{}

func init() {
	envconfig.Process("canaria", conf)
}

func GetEnv() string {
	return conf.Env
}

func GetMailAddress() string {
	return conf.MailAddress
}

func GetLogLevel() string {
	return conf.LogLevel
}

func GetHost() string {
	return conf.Host
}

func GetPort() string {
	return conf.Port
}
