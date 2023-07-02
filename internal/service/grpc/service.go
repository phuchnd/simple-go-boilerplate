package grpc

import (
	"context"
	"fmt"
	"github.com/phuchnd/simple-go-boilerplate/internal/db/repositories"
	"github.com/phuchnd/simple-go-boilerplate/internal/db/repositories/entities"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc/pb"
)

//go:generate mockery --name=IGRPCService --case=snake --disable-version-string
type IGRPCService interface {
	ListBooks(context.Context, *pb.ListBookRequest) (*pb.ListBookResponse, error)
}

type implGRPCService struct {
	bookRepo repositories.IBookRepository
}

func NewGRPCService(bookRepo repositories.IBookRepository) IGRPCService {
	return &implGRPCService{
		bookRepo: bookRepo,
	}
}

func (s *implGRPCService) ListBooks(ctx context.Context, req *pb.ListBookRequest) (*pb.ListBookResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("empty input")
	}
	listRes, err := s.bookRepo.ListBooks(ctx, int(req.Limit), entities.ID(req.Cursor), nil)
	if err != nil {
		return nil, err
	}
	return ListBookResponseFromDBEntitiesToPB(listRes), nil
}
