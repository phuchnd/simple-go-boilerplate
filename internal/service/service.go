package service

import (
	"context"
	pb "github.com/phuchnd/core-product-management/server/pb"
)

func NewProductManagementService() *ProductManagementService {
	return &ProductManagementService{}
}

type ProductManagementService struct {
	pb.UnimplementedProductManagementServiceServer
}

func (s *ProductManagementService) ListUnitOfMeasurements(context.Context, *pb.ListUnitOfMeasurementsRequest) (*pb.ListUnitOfMeasurementsResponse, error) {
	return &pb.ListUnitOfMeasurementsResponse{
		Entries: []*pb.UnitOfMeasurement{
			{
				ID:   1,
				Name: "Name 2",
			},
		},
	}, nil
}
