package env

import (
	"github.com/Sirupsen/logrus"
	"github.com/dogenzaka/ruslog"
)

type LocalEnvironment struct {
	EnvName string
	Debug   bool
	Bind    string
	Loggers []*ruslog.Logger
	// DynamoDBConfig  ServerConfig
	// RedisConfig     ServerConfig
	// RedirectConfig  RedirectConfig
	// CipherMetaInfos map[string]*CipherMetaInfo
}

func (e *LocalEnvironment) SetUp() error {
	e.EnvName = Local
	e.Bind = ":5000"
	println("~~~~~~~~~~~~~~~~~~~~~~~~")

	logDir := "/tmp"
	e.Loggers = append(e.Loggers, &ruslog.Logger{
		Name:        "Default",
		Type:        ruslog.APPENDER_DAILY,
		Level:       logrus.DebugLevel.String(),
		Format:      ruslog.FORMATTER_SIMPLE,
		FilePath:    logDir + "/app.log",
		MaxRotation: 10,
		AddFileInfo: false,
	})

	return nil
}

func (e *LocalEnvironment) GetEnvName() string {
	return e.EnvName
}

func (e *LocalEnvironment) GetDebug() bool {
	return e.Debug
}

func (e *LocalEnvironment) GetBind() string {
	return e.Bind
}

// GetLoggers returns loggers
func (e *LocalEnvironment) GetLoggers() []*ruslog.Logger {
	return e.Loggers
}
