package http

import (
	"fmt"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"github.com/phuchnd/simple-go-boilerplate/internal/cron"
	"github.com/phuchnd/simple-go-boilerplate/internal/cron/jobs"
	mysqldb "github.com/phuchnd/simple-go-boilerplate/internal/db/mysql"
	"github.com/phuchnd/simple-go-boilerplate/internal/health"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	http2 "github.com/phuchnd/simple-go-boilerplate/internal/service/http"
	service "github.com/phuchnd/simple-go-boilerplate/internal/service/http"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//go:generate mockery --name=IServer --case=snake --disable-version-string
type IServer interface {
	Start() error
	Stop() error
}

type httpServerImpl struct {
	healthCheckSvc health.IHealthCheck
	isReady        bool
	quit           chan os.Signal

	server    *http.Server
	runner    cron.IJobRunner
	handler   service.IHTTPService
	serverCfg *config.ServerConfig
	logger    logging.Logger
}

func NewServer(logger logging.Logger, cfgProvider config.IConfig, db mysqldb.IMySqlDB, handler http2.IHTTPService) (IServer, error) {
	healthCheckSvc := health.NewHealthCheck(db, logger)
	s := &httpServerImpl{
		handler:        handler,
		serverCfg:      cfgProvider.GetServerConfig(),
		logger:         logger,
		healthCheckSvc: healthCheckSvc,
		runner:         initJobRunner(logger, cfgProvider, healthCheckSvc),
	}
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.serverCfg.HTTPPort),
		Handler: s.initRouter(),
	}
	// Init jobs
	s.runner.Start()

	return s, nil
}

func initJobRunner(logger logging.Logger, cfgProvider config.IConfig, hcSvc health.IHealthCheck) cron.IJobRunner {
	cronSystem := cron.NewCron(logger)
	runner := cron.NewJobRunner(logger, cronSystem)
	cfg := cfgProvider.GetCronHealthCheckConfig()
	simpleExampleJob := jobs.NewCronHealthCheck(hcSvc, logger, cfg)
	_ = runner.RegisterJob(simpleExampleJob)

	return runner
}

func (s *httpServerImpl) Start() error {
	if err := s.healthCheckSvc.Check(); err != nil {
		s.logger.Errorf("HTTP server initialization failed, dependency not ready: %s", err.Error())
		return err
	}
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			s.logger.Errorf("failed to start HTTP server %s", err.Error())
		}
	}()
	s.logger.Infof("HTTP server started successfully at port %d", s.serverCfg.HTTPPort)

	s.quit = make(chan os.Signal)
	signal.Notify(s.quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
	<-s.quit

	s.logger.Info("received terminated signal, server is shutting down")
	return s.Stop()
}

func (s *httpServerImpl) Stop() error {
	return nil
}
