package cron

import (
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
)

//go:generate mockery --name=IJobRunner --case=snake --disable-version-string
type IJobRunner interface {
	RegisterJob(job IJob) error
	Start()
	Stop()
}

type jobRunnerImpl struct {
	logger logging.Logger
	cron   ICron
}

func NewJobRunner(logger logging.Logger, cron ICron) IJobRunner {
	return &jobRunnerImpl{
		logger: logger,
		cron:   cron,
	}
}

func (impl *jobRunnerImpl) RegisterJob(job IJob) error {
	_, err := impl.cron.AddFunc(job.CronSpec(), func() {
		defer func() {
			if r := recover(); r != nil {
				impl.logger.Errorf("panic during job run %v", r)
			}
		}()

		if !job.IsEnable() {
			return
		}

		job.Run()
	})
	return err
}

func (impl *jobRunnerImpl) Start() {
	impl.cron.Start()
}

func (impl *jobRunnerImpl) Stop() {
	impl.cron.Stop()
}
