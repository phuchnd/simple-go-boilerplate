package http

import (
	"context"
	bookservice "github.com/phuchnd/simple-go-boilerplate/external/book-grpc-service"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"github.com/phuchnd/simple-go-boilerplate/internal/service/http/entities"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc/pb"
)

//go:generate mockery --name=IHTTPService --case=snake --disable-version-string
type IHTTPService interface {
	ListBooks(context.Context, *entities.ListBookRequest) (*entities.ListBookResponse, error)
}

type implHTTPService struct {
	bookService bookservice.IBookService
}

func NewHTTPService() (IHTTPService, error) {
	bookService, err := bookservice.NewService(config.GetBookConfig())
	if err != nil {
		return nil, err
	}
	return &implHTTPService{
		bookService: bookService,
	}, nil
}

func (s *implHTTPService) ListBooks(ctx context.Context, req *entities.ListBookRequest) (*entities.ListBookResponse, error) {
	listResp, err := s.bookService.ListBooks(ctx, &pb.ListBookRequest{
		Limit:  req.Limit,
		Cursor: req.Cursor,
	})
	if err != nil {
		return nil, err
	}
	return ListBookResponseFromPBToEntities(listResp), nil
}
