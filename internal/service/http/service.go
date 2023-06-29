package http

import (
	"context"
	"github.com/phuchnd/simple-go-boilerplate/internal/service/http/entities"
)

type IHTTPService interface {
	ListBooks(context.Context, *entities.ListBookRequest) (*entities.ListBookResponse, error)
}

type implHTTPService struct {
}

func NewHTTPService() IHTTPService {
	return &implHTTPService{}
}

func (s *implHTTPService) ListBooks(context.Context, *entities.ListBookRequest) (*entities.ListBookResponse, error) {
	return nil, nil
}
