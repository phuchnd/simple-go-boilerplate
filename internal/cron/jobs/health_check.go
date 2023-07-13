package jobs

import (
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"github.com/phuchnd/simple-go-boilerplate/internal/cron"
	"github.com/phuchnd/simple-go-boilerplate/internal/health"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	"time"
)

//go:generate mockery --name=ICronHealthCheck --case=snake --disable-version-string
type ICronHealthCheck interface {
	cron.IJob
}

type healthCheckImpl struct {
	config         *config.CronConfig
	logger         logging.Logger
	healthCheckSvc health.IHealthCheck
}

func NewCronHealthCheck(healthCheckSvc health.IHealthCheck, logger logging.Logger, cfg *config.CronConfig) ICronSimpleExample {
	return &healthCheckImpl{
		logger:         logger,
		config:         cfg,
		healthCheckSvc: healthCheckSvc,
	}
}

func (s *healthCheckImpl) CronSpec() string {
	return s.config.CronSpec
}

func (s *healthCheckImpl) IsEnable() bool {
	return s.config.Enable
}

func (s *healthCheckImpl) Run() {
	logger := s.logger
	logger.WithField("time", time.Now().String()).Infof("ICronHealthCheck has been executed")
	if err := s.healthCheckSvc.Check(); err != nil {
		logger.WithField("time", time.Now().String()).Errorf("ICronHealthCheck detect dependency error")
	}
}
