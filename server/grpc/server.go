package grpc

import (
	"fmt"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"github.com/phuchnd/simple-go-boilerplate/internal/db/repositories"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	grpc2 "github.com/phuchnd/simple-go-boilerplate/internal/service/grpc"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc/endpoints"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc/middlewares"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc/pb"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

//go:generate mockery --name=IServer --case=snake --disable-version-string
type IServer interface {
	Start() error
	Stop() error
}

type serverImpl struct {
	grpcServer pb.SimpleGoBoilerplateServiceServer

	serverCfg *config.ServerConfig
	logger    logging.Logger
	isRunning bool
	quit      chan os.Signal
}

func NewServer(logger logging.Logger) (IServer, error) {
	serverCfg := config.GetServerConfig()

	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverCfg.GRPCPort))
	if err != nil {
		return nil, err
	}

	bookRepo, err := repositories.NewBookRepository()
	if err != nil {
		return nil, err
	}
	serviceHandler := grpc2.NewGRPCService(bookRepo)
	grpcEndpoints := endpoints.MakeEndPoints(serviceHandler,
		middlewares.Tracing(),
		middlewares.RequestLogging(),
		middlewares.PanicRecovery(),
	)
	server := &serverImpl{
		grpcServer: NewGRPCServer(grpcEndpoints),

		serverCfg: serverCfg,
		logger:    logger,
	}
	baseServer := grpc.NewServer()
	pb.RegisterSimpleGoBoilerplateServiceServer(baseServer, server.grpcServer)

	logger.Infof("gRPC server is running on :%d", serverCfg.GRPCPort)
	if err := baseServer.Serve(grpcListener); err != nil {
		logger.Errorf("failed to start gRPC server", "err", err)
		return nil, err
	}

	return server, nil
}

func (s *serverImpl) Start() error {
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.serverCfg.GRPCPort))
	if err != nil {
		s.logger.Errorf("GRPC server initialization failed: %s", err.Error())
		return err
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterSimpleGoBoilerplateServiceServer(baseServer, s.grpcServer)
		s.logger.Infof("GRPC server started successfully at port %d", s.serverCfg.GRPCPort)

		if err := baseServer.Serve(grpcListener); err != nil {
			s.logger.Errorf("failed to start GRPC server %s", err.Error())
		}
	}()

	s.quit = make(chan os.Signal)
	signal.Notify(s.quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
	<-s.quit

	s.logger.Info("received terminated signal, server is shutting down")
	return s.Stop()
}

func (s *serverImpl) Stop() error {
	return nil
}
