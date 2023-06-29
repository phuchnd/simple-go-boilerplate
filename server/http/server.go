package http

import (
	"fmt"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	http2 "github.com/phuchnd/simple-go-boilerplate/internal/service/http"
	service "github.com/phuchnd/simple-go-boilerplate/internal/service/http"
	"net/http"
)

type IHTTPServer interface {
	Start() error
	Stop() error
}

type httpServerImpl struct {
	isReady   bool
	server    *http.Server
	handler   service.IHTTPService
	appConfig *config.ServerConfig
}

func NewHTTPServer() IHTTPServer {
	serverCfg := config.GetServerConfig()
	httpServerHandler := http2.NewHTTPService()

	s := &httpServerImpl{
		handler:   httpServerHandler,
		appConfig: serverCfg,
	}
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.appConfig.HTTPPort),
		Handler: s.initRouter(),
	}

	return s
}

func (s *httpServerImpl) Start() error {
	return s.server.ListenAndServe()
}

func (s *httpServerImpl) Stop() error {
	return nil
}
