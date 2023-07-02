package jobs

import (
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"github.com/phuchnd/simple-go-boilerplate/internal/cron"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	"time"
)

//go:generate mockery --name=ICronSimpleExample --case=snake --disable-version-string
type ICronSimpleExample interface {
	cron.IJob
}

type cronSimpleExampleImpl struct {
	config *config.CronConfig
	logger logging.Logger
}

func NewCronSimpleExample(logger logging.Logger, cfg *config.CronConfig) ICronSimpleExample {
	return &cronSimpleExampleImpl{
		logger: logger,
		config: cfg,
	}
}

func (s *cronSimpleExampleImpl) CronSpec() string {
	return s.config.CronSpec
}

func (s *cronSimpleExampleImpl) IsEnable() bool {
	return s.config.Enable
}

func (s *cronSimpleExampleImpl) Run() {
	logger := s.logger
	logger.WithField("time", time.Now().String()).Infof("ICronSimpleExample has been executed")
	// TODO implement anything here
}
