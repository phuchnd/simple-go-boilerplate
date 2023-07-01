package grpc

import (
	"context"
	service "github.com/phuchnd/simple-go-boilerplate/internal/service/grpc"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc/pb"
)

type grpcEndpointsImpl struct {
	pb.UnimplementedSimpleGoBoilerplateServiceServer
	handler service.IGRPCService
}

func NewGRPCEndpoints(handler service.IGRPCService) pb.SimpleGoBoilerplateServiceServer {
	return &grpcEndpointsImpl{
		handler: handler,
	}
}

func (ep *grpcEndpointsImpl) ListBooks(ctx context.Context, req *pb.ListBookRequest) (*pb.ListBookResponse, error) {
	return ep.handler.ListBooks(ctx, req)
}
