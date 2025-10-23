package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger() *logrus.Logger {
	log := logrus.New()

	logWriter := &lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10,
		MaxBackups: 7,
		MaxAge:     30,
		Compress:   true,
		LocalTime:  true,
	}

	multiWriter := io.MultiWriter(os.Stdout, logWriter)
	log.SetOutput(multiWriter)

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		PadLevelText:    true,
	})

	log.SetLevel(logrus.InfoLevel)

	return log
}

func NewLoggerWithConfig(config LogConfig) *logrus.Logger {
	log := logrus.New()

	logWriter := &lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
		LocalTime:  true,
	}

	multiWriter := io.MultiWriter(os.Stdout, logWriter)
	log.SetOutput(multiWriter)

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		PadLevelText:    true,
	})

	log.SetLevel(config.Level)

	return log
}

type LogConfig struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	Level      logrus.Level
}

func DefaultConfig() LogConfig {
	return LogConfig{
		Filename:   "logs/app.log",
		MaxSize:    10,
		MaxBackups: 7,
		MaxAge:     30,
		Compress:   true,
		Level:      logrus.InfoLevel,
	}
}
