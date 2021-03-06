package env

import (
	"os"

	"fmt"
	"log"

	"github.com/CANARIA/canaria-api/core/config"
	//"github.com/CANARIA/canaria-api/core/logger"
)

const (
	Local      string = "LOCAL"
	Staging    string = "STG"
	Production string = "PRD"
)

type (
	Environment interface {
		SetUp() error
		GetEnvName() string
		GetDebug() bool
		GetBind() string
		//GetLoggers() []logger.Config
		GetDBConfig() *DBConfig
		// GetDynamoDBConfig() *ServerConfig
		// GetRedirectConfig() *RedirectConfig
		// GetCipherMetaInfo(name string) *CipherMetaInfo
		// GetCipherMetaInfos() *map[string]*CipherMetaInfo
	}

	DB struct {
		User               string
		Password           string
		Host               string
		Port               int
		DBName             string
		MaxConnections     int
		MaxIdleConnections int
	}

	DBConfig struct {
		Master *DB
		Slaves []*DB
		LogMode bool
	}
)

func (d *DB) GetMySQLDataSource() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Local", d.User, d.Password, d.Host, d.Port, d.DBName)
}

var env Environment

func SetUp() {

	debug := os.Getenv(config.IsDebug) != ""

	switch os.Getenv(config.Env) {
	case Staging:
		log.Printf("This environment is %s", Staging)
		env = &StgEnvironment{
			EnvName: Staging,
			Debug:   debug,
		}
		if err := env.SetUp(); err != nil {
			panic(err)
		}
	case Production:
		log.Printf("This environment is %s", Production)
		env = &StgEnvironment{
			EnvName: Production,
			Debug:   debug,
		}
		if err := env.SetUp(); err != nil {
			panic(err)
		}
	default:
		log.Printf("This environment is %s", Local)
		env = &LocalEnvironment{
			EnvName: Local,
			Debug:   debug,
		}
		if err := env.SetUp(); err != nil {
			panic(err)
		}
	}

	//logger.Configure(env.GetLoggers())
}

func GetEnvName() string {
	return env.GetEnvName()
}

func GetDebug() bool {
	return env.GetDebug()
}

func GetBind() string {
	return env.GetBind()
}

//func GetLoggers() []logger.Config {
//	return env.GetLoggers()
//}

func GetDBConfig() *DBConfig {
	return env.GetDBConfig()
}
