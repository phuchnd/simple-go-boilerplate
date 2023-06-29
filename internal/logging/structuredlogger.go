package logging

import (
	"context"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	WithField(key string, value interface{}) *logrus.Entry
	WithFields(fields logrus.Fields) *logrus.Entry
	WithError(err error) *logrus.Entry
	WithContext(ctx context.Context) *logrus.Entry
	WithTime(t time.Time) *logrus.Entry

	Logf(level logrus.Level, format string, args ...interface{})
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Log(level logrus.Level, args ...interface{})
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

var (
	fieldMap = logrus.FieldMap{
		logrus.FieldKeyMsg: "message",
	}
)

func NewStructuredLogger(cfg *config.ServerConfig) Logger {
	logger := logrus.New()

	if cfg.Env == config.LocalEnv {
		logger.SetFormatter(&logrus.TextFormatter{FieldMap: fieldMap})
	} else {
		logger.SetFormatter(&logrus.JSONFormatter{FieldMap: fieldMap})
	}

	return logger
}
