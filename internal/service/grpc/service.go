package grpc

import (
	"context"
	pb2 "github.com/phuchnd/simple-go-boilerplate/server/grpc/pb"
)

type implGRPCService struct {
	pb2.UnimplementedSimpleGoBoilerplateServiceServer
}

func NewGRPCService() *implGRPCService {
	return &implGRPCService{}
}

func (s *implGRPCService) ListBooks(context.Context, *pb2.ListBookRequest) (*pb2.ListBookResponse, error) {
	return &pb2.ListBookResponse{
		Entries: []*pb2.Book{
			{
				ID:    1,
				Title: "Book 1",
			},
		},
	}, nil
}
