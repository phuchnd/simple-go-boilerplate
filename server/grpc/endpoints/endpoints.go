package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/phuchnd/simple-go-boilerplate/internal/service/grpc"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc/pb"
)

const (
	endpointMetadataKey = "endpoints-metadata"
)

type EndpointBuilder func(svc grpc.IGRPCService) (endpoint.Endpoint, *EndpointMetadata)

type EndpointMetadata struct {
	Name string
}

type Endpoints struct {
	ListBooks endpoint.Endpoint
}

func GetEndpointMetadataFromContext(ctx context.Context) (*EndpointMetadata, bool) {
	metadata, ok := ctx.Value(endpointMetadataKey).(*EndpointMetadata)
	return metadata, ok
}

func SetEndpointMetadataToContext(ctx context.Context, metadata *EndpointMetadata) context.Context {
	return context.WithValue(ctx, endpointMetadataKey, metadata)
}

func MakeEndPoints(svc grpc.IGRPCService, middlewares ...endpoint.Middleware) *Endpoints {
	return &Endpoints{
		ListBooks: makeEndpoint(svc, makeListBooksEndpoint, middlewares...),
	}
}

func makeEndpoint(svc grpc.IGRPCService, builder EndpointBuilder, middlewares ...endpoint.Middleware) endpoint.Endpoint {
	e, metadata := builder(svc)
	for _, m := range middlewares {
		e = m(e)
	}

	var metadataMiddleware endpoint.Middleware
	metadataMiddleware = func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			ctx = context.WithValue(ctx, endpointMetadataKey, metadata)
			return next(ctx, request)
		}
	}

	return metadataMiddleware(e)
}

func makeListBooksEndpoint(svc grpc.IGRPCService) (endpoint.Endpoint, *EndpointMetadata) {
	metadata := &EndpointMetadata{
		Name: "ListBooks",
	}

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		response, err := svc.ListBooks(ctx, request.(*pb.ListBookRequest))
		return response, err
	}, metadata
}
