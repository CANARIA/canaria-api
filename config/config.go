package config

import (
	"github.com/kelseyhightower/envconfig"
)

// TODO: yamlに移動したい
const (
	USER                   string = "root"
	PASSWORD               string = "password"
	DB                     string = "canaria"
	HOST                   string = "mysql"
	PORT                   string = "3306"
	PRE_REGISTER_MAIL_BODY string = `Canariaにご登録ありがとうございます。
  24時間以内に下記のURLからご登録下さい。
  `
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
