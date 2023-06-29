package grpc

import (
	"fmt"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	grpc2 "github.com/phuchnd/simple-go-boilerplate/internal/service/grpc"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc/pb"
	"google.golang.org/grpc"
	"net"
)

type IGRPCServer interface {
	Start() error
	Stop() error
}

type grpcServerImpl struct {
	grpcServer pb.SimpleGoBoilerplateServiceServer
}

func NewGRPCServer() (IGRPCServer, error) {
	serverCfg := config.GetServerConfig()

	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverCfg.GRPCPort))
	if err != nil {
		return nil, err
	}

	grpcServerHandler := grpc2.NewGRPCService()
	server := &grpcServerImpl{
		grpcServer: grpcServerHandler,
	}
	logger := logging.NewLogger(serverCfg)
	baseServer := grpc.NewServer()
	pb.RegisterSimpleGoBoilerplateServiceServer(baseServer, server.grpcServer)

	logger.Infof("gRPC server is running on :%d", serverCfg.GRPCPort)
	if err := baseServer.Serve(grpcListener); err != nil {
		logger.Errorf("failed to start gRPC server", "err", err)
		return nil, err
	}

	return server, nil
}

func (s *grpcServerImpl) Start() error {
	return nil
}

func (s *grpcServerImpl) Stop() error {
	return nil
}
