package cron

import (
	"context"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	"github.com/robfig/cron/v3"
)

type cronLogger struct {
	logger logging.Logger
}

func (l *cronLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Infof(msg, keysAndValues...)
}

func (l *cronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	keysAndValues = append(keysAndValues, "error", err)
	l.logger.Errorf(msg, keysAndValues...)
}

//go:generate mockery --name=ICron --case=snake --disable-version-string
type ICron interface {
	AddFunc(spec string, cmd func()) (cron.EntryID, error)
	Start()
	Stop() context.Context
}

func NewCron(logger logging.Logger) ICron {
	return cron.New(cron.WithChain(cron.Recover(&cronLogger{logger: logger})))
}
