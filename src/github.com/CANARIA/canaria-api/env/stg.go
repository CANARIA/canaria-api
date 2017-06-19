package env

import (
	"github.com/CANARIA/canaria-api/config"
	"github.com/CANARIA/canaria-api/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type StgEnvironment struct {
	EnvName  string
	Debug    bool
	Bind     string
	Loggers  []logger.Config
	DBConfig *DBConfig
	// DynamoDBConfig  ServerConfig
	// RedisConfig     ServerConfig
	// RedirectConfig  RedirectConfig
	// CipherMetaInfos map[string]*CipherMetaInfo
}

func (e *StgEnvironment) SetUp() error {
	e.Bind = ":5000"

	logDir := "/var/log/canaria/"

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
			Name:  config.QueryLoggerName,
			Level: zap.InfoLevel,
			EncoderConfig: &zapcore.EncoderConfig{
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
			Name:  config.MutationLoggerName,
			Level: zap.InfoLevel,
			EncoderConfig: &zapcore.EncoderConfig{
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

	e.DBConfig = &DBConfig{
		Master: &DB{
			Host:               "127.0.0.1",
			Port:               3306,
			User:               "root",
			Password:           "xo2FrGiwmFqIC2kA",
			DBName:             "canaria",
			MaxConnections:     5,
			MaxIdleConnections: 5,
		},
		Slaves: []*DB{
			&DB{
				Host:               "127.0.0.1",
				Port:               3306,
				User:               "root",
				Password:           "",
				DBName:             "canaria",
				MaxConnections:     5,
				MaxIdleConnections: 5,
			},
		},
		LogMode: true,
	}

	return nil
}

func (e *StgEnvironment) GetEnvName() string {
	return e.EnvName
}

func (e *StgEnvironment) GetDebug() bool {
	return e.Debug
}

func (e *StgEnvironment) GetBind() string {
	return e.Bind
}

// GetLoggers returns loggers
func (e *StgEnvironment) GetLoggers() []logger.Config {
	return e.Loggers
}

func (e *StgEnvironment) GetDBConfig() *DBConfig {
	return e.DBConfig
}
