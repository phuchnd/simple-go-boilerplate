package grpc

import (
	"context"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc/endpoints"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc/pb"
)

type grpcServerImpl struct {
	pb.UnimplementedSimpleGoBoilerplateServiceServer
	listBooks gt.Handler
}

func NewGRPCServer(grpcEndpoints *endpoints.Endpoints) pb.SimpleGoBoilerplateServiceServer {
	return &grpcServerImpl{
		listBooks: gt.NewServer(
			grpcEndpoints.ListBooks,
			decodeListBookRequest,
			encodeListBookResponse,
		),
	}
}

func (i *grpcServerImpl) ListBooks(ctx context.Context, req *pb.ListBookRequest) (*pb.ListBookResponse, error) {
	_, resp, err := i.listBooks.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ListBookResponse), nil
}

func decodeListBookRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ListBookRequest)
	return req, nil
}

func encodeListBookResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*pb.ListBookResponse)
	return res, nil
}
