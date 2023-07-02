package book_grpc_service

import (
	"context"
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/phuchnd/simple-go-boilerplate/external/book-grpc-service/mocks"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc/pb"
)

var _ = Describe("NewService", func() {
	It("should not be nil when init", func() {
		s, err := NewService(&config.BookConfig{})
		Expect(s).ShouldNot(BeNil())
		Expect(err).Should(BeNil())
	})
})

var _ = Describe("ListBooks", func() {
	var ctx context.Context
	var req *pb.ListBookRequest
	var grpcClient *mocks.IServiceClient

	BeforeEach(func() {
		req = &pb.ListBookRequest{}
	})

	It("should not return err", func() {
		grpcClient = new(mocks.IServiceClient)
		grpcClient.On("ListBooks", ctx, req).Return(&pb.ListBookResponse{}, nil)
		s := &bookServiceImpl{
			client: grpcClient,
		}

		res, err := s.ListBooks(ctx, req)
		Expect(res).NotTo(BeNil())
		Expect(err).Should(BeNil())
	})

	It("should return err", func() {
		grpcClient = new(mocks.IServiceClient)
		grpcClient.On("ListBooks", ctx, req).Return(nil, errors.New("err ListBooks"))
		s := &bookServiceImpl{
			client: grpcClient,
		}

		res, err := s.ListBooks(ctx, req)
		Expect(res).Should(BeNil())
		Expect(err).NotTo(BeNil())
	})
})
