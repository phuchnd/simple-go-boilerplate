package book_grpc_service

import (
	"context"
	"fmt"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	grpctransport "github.com/phuchnd/simple-go-boilerplate/internal/transport/grpc"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc/pb"
	"google.golang.org/grpc/status"
)

//go:generate mockery --name=IBookService --case=snake --disable-version-string
type IBookService interface {
	ListBooks(ctx context.Context, in *pb.ListBookRequest) (*pb.ListBookResponse, error)
}

type bookServiceImpl struct {
	client pb.SimpleGoBoilerplateServiceClient
}

func NewService() (IBookService, error) {
	conf := config.GetBookConfig()

	cc, err := grpctransport.NewGRCPClientConn(&grpctransport.TransportConfig{
		ServiceName:         "book-api",
		ExternalServiceName: "book-service",
		Host:                conf.Host,
		Port:                conf.Port,
		MaxRetries:          conf.MaxRetries,
		BackoffDelaysMs:     conf.BackoffDelaysMs,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to initialized order service \n config: %v \n error: %w ", conf, err)
	}
	return &bookServiceImpl{
		client: pb.NewSimpleGoBoilerplateServiceClient(cc),
	}, nil
}

func (s *bookServiceImpl) ListBooks(ctx context.Context, in *pb.ListBookRequest) (*pb.ListBookResponse, error) {
	resp, err := s.client.ListBooks(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("book.ListBooks failed with status code: %v", status.Code(err))
	}
	return resp, nil
}
