package http

import (
	"fmt"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
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
	// isRunning for health check TODO
	isRunning bool
	quit      chan os.Signal

	server    *http.Server
	handler   service.IHTTPService
	serverCfg *config.ServerConfig
	logger    logging.Logger
}

func NewServer(logger logging.Logger) (IServer, error) {
	serverCfg := config.GetServerConfig()
	httpServerHandler, err := http2.NewHTTPService()
	if err != nil {
		return nil, err
	}

	s := &httpServerImpl{
		handler:   httpServerHandler,
		serverCfg: serverCfg,
		logger:    logger,
	}
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.serverCfg.HTTPPort),
		Handler: s.initRouter(),
	}

	return s, nil
}

func (s *httpServerImpl) Start() error {
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
