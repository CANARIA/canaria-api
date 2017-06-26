package logger

import (
	"os"
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type (
	// Config for logging
	Config struct {
		// Name is log name unique and required.
		Name string
		// Level is log level on zapcore.Level
		Level zapcore.Level
		// EncodeLogsAsJSON makes the log framework log JSON
		EncodeLogsAsJSON bool
		// FileLoggingEnabled makes the framework log to a file
		// the fields below can be skipped if this value is false!
		FileLoggingEnabled bool
		// EnabledCaller is true if output file's path and row number log
		EnabledCaller bool
		// EncoderConfig is zapcore.EncoderConfig
		EncoderConfig *zapcore.EncoderConfig
		// Directory to log to to when filelogging is enabled
		Directory string
		// Filename is the name of the logfile which will be placed inside the directory
		Filename string
		// MaxSize the max size in MB of the logfile before it's rolled
		MaxSize int
		// MaxBackups the max number of rolled files to keep
		MaxBackups int
		// MaxAge the max age in days to keep a logfile
		MaxAge int
	}
)

var (
	logs                 = map[string]*zap.SugaredLogger{}
	defaultEncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "file",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
)

// Configure is Configure
func Configure(config []Config) {
	for _, c := range config {
		var writer zapcore.WriteSyncer = os.Stdout
		if c.FileLoggingEnabled {
			writer = newRollingFile(c)
		}
		log := newZapLogger(c.EncodeLogsAsJSON, zap.NewAtomicLevelAt(c.Level), c.EnabledCaller, c.EncoderConfig, zapcore.NewMultiWriteSyncer(writer))
		logs[c.Name] = log.Sugar()
	}
}

// GetLogger return zaplogger
func GetLogger(name string) *zap.SugaredLogger {
	l, exists := logs[name]
	if !exists {
		l = newZapLogger(false, zap.NewAtomicLevelAt(zapcore.DebugLevel), true, nil, os.Stdout).Sugar()
	}
	return l
}

func newRollingFile(config Config) zapcore.WriteSyncer {
	if err := os.MkdirAll(config.Directory, 0); err != nil {
		panic(err)
	}

	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   path.Join(config.Directory, config.Filename),
		MaxSize:    config.MaxSize,    //megabytes
		MaxAge:     config.MaxAge,     //days
		MaxBackups: config.MaxBackups, //files
	})
}

func newZapLogger(encodeAsJSON bool, level zapcore.LevelEnabler, enabledCaller bool, encoderConfig *zapcore.EncoderConfig, output zapcore.WriteSyncer) *zap.Logger {

	encCfg := defaultEncoderConfig
	if encoderConfig != nil {
		encCfg = *encoderConfig
	}

	encoder := zapcore.NewConsoleEncoder(encCfg)
	if encodeAsJSON {
		encoder = zapcore.NewJSONEncoder(encCfg)
	}

	zapLogger := zap.New(zapcore.NewCore(encoder, output, level))
	if enabledCaller {
		zapLogger = zapLogger.WithOptions(zap.AddCaller())
	}

	return zapLogger
}
