package env

import (
	"os"

	"fmt"

	"github.com/CANARIA/canaria-api/config"
	"github.com/Sirupsen/logrus"
	"github.com/dogenzaka/ruslog"
	"github.com/CANARIA/canaria-api/logger"
)

const (
	Local      string = "LOCAL"
	Staging    string = "STG"
	Production string = "PRD"
)

type Environment interface {
	SetUp() error
	GetEnvName() string
	GetDebug() bool
	GetBind() string
	GetLoggers() []*logger.Config
	// GetDynamoDBConfig() *ServerConfig
	// GetRedirectConfig() *RedirectConfig
	// GetCipherMetaInfo(name string) *CipherMetaInfo
	// GetCipherMetaInfos() *map[string]*CipherMetaInfo
}

var env Environment

func SetUp(e Environment) {
	env = e
	logger.Configure(e.GetLoggers())

	//logrus.SetFormatter(&logrus.JSONFormatter{})

	debug := os.Getenv(config.IsDebug) != ""

	switch os.Getenv(config.Env) {
	//case Staging:
	//	fmt.Println("This environment is STG")
	//	env = &LocalEnvironment{
	//		EnvName: Staging,
	//		Debug: debug,
	//	}
	//	if err := env.SetUp(); err != nil {
	//		panic(err)
	//	}
	default:
		fmt.Println("This environment is LOCAL")
		env = &LocalEnvironment{
			EnvName: Local,
			Debug: debug,
		}
		if err := env.SetUp(); err != nil {
			panic(err)
		}
	}

	ruslog.Configure(env.GetLoggers())
}
