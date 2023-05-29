package server

import (
	"fmt"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"github.com/phuchnd/simple-go-boilerplate/internal/service"
	pb "github.com/phuchnd/simple-go-boilerplate/server/pb"
	"google.golang.org/grpc"
	"net"
)

type IServer interface {
	Start() error
	Stop() error
}

type serverImpl struct {
	grpcServer pb.ProductManagementServiceServer
}

func NewServer() (IServer, error) {
	serverCfg := config.GetServerConfig()

	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverCfg.GRPCPort))
	if err != nil {
		return nil, err
	}

	grpcServerHandler := service.NewProductManagementService()
	server := &serverImpl{
		grpcServer: grpcServerHandler,
	}
	baseServer := grpc.NewServer()
	pb.RegisterProductManagementServiceServer(baseServer, server.grpcServer)

	fmt.Printf("gRPC server is running on :%d", serverCfg.GRPCPort)
	if err := baseServer.Serve(grpcListener); err != nil {
		fmt.Println("failed to start gRPC server", "err", err)
	}

	return server, nil
}

func (s *serverImpl) Start() error {
	return nil
}

func (s *serverImpl) Stop() error {
	return nil
}
