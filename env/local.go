package env

import (
	"github.com/CANARIA/canaria-api/config"
	"github.com/CANARIA/canaria-api/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LocalEnvironment struct {
	EnvName string
	Debug   bool
	Bind    string
	Loggers []logger.Config
	// DynamoDBConfig  ServerConfig
	// RedisConfig     ServerConfig
	// RedirectConfig  RedirectConfig
	// CipherMetaInfos map[string]*CipherMetaInfo
}

func (e *LocalEnvironment) SetUp() error {
	e.EnvName = Local
	e.Bind = ":5000"

	logDir := "/tmp"

	canariaLogs := []logger.Config{
		logger.Config{
			Name:               config.DefaultLoggerName,
			Level:              zap.DebugLevel,
			EncodeLogsAsJSON:   true,
			FileLoggingEnabled: true,
			EnabledCaller:      true,
			Directory:          logDir,
			Filename:           config.DefaultLoggerName + ".log",
			MaxBackups:         30,
			MaxSize:            100,
			MaxAge:             30,
		},
		logger.Config{
			Name:               config.AccessLoggerName,
			Level:              zap.DebugLevel,
			EncodeLogsAsJSON:   false,
			FileLoggingEnabled: true,
			Directory:          logDir,
			Filename:           config.AccessLoggerName + ".log",
			MaxBackups:         30,
			MaxSize:            100,
			MaxAge:             30,
		},
		logger.Config{
			Name:               config.SlowQueryLoggerName,
			Level:              zap.DebugLevel,
			EncodeLogsAsJSON:   false,
			FileLoggingEnabled: true,
			Directory:          logDir,
			Filename:           config.SlowQueryLoggerName + ".log",
			MaxBackups:         30,
			MaxSize:            100,
			MaxAge:             30,
		},
		logger.Config{
			Name:  				config.QueryLoggerName,
			Level: 				zap.InfoLevel,
			EncoderConfig: 		&zapcore.EncoderConfig{
									MessageKey: "msg",
								},
			EncodeLogsAsJSON:   false,
			FileLoggingEnabled: true,
			Directory:          logDir,
			Filename:           config.QueryLoggerName + ".log",
			MaxBackups:         30,
			MaxSize:            100,
			MaxAge:             30,
		},
		logger.Config{
			Name:  				config.MutationLoggerName,
			Level: 				zap.InfoLevel,
			EncoderConfig: 		&zapcore.EncoderConfig{
									MessageKey: "msg",
								},
			EncodeLogsAsJSON:   false,
			FileLoggingEnabled: true,
			Directory:          logDir,
			Filename:           config.MutationLoggerName + ".log",
			MaxBackups:         30,
			MaxSize:            100,
			MaxAge:             30,
		},
	}
	e.Loggers = canariaLogs

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
func (e *LocalEnvironment) GetLoggers() []logger.Config {
	return e.Loggers
}
