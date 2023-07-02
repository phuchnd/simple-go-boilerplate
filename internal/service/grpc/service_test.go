package grpc

import (
	"context"
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/phuchnd/simple-go-boilerplate/internal/db/repositories/entities"
	"github.com/phuchnd/simple-go-boilerplate/internal/db/repositories/mocks"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var _ = Describe("NewGRPCService", func() {
	It("should not be nil when init", func() {
		repo := new(mocks.IBookRepository)
		s := NewGRPCService(repo)
		Expect(s).ShouldNot(BeNil())
	})
})

var _ = Describe("ListBooks", func() {
	Context("with invalid input", func() {
		var (
			ctx context.Context
		)

		BeforeEach(func() {
			ctx = context.TODO()
		})
		It("should return error", func() {
			s := &implGRPCService{}
			resp, err := s.ListBooks(ctx, nil)
			Expect(resp).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("with valid input", func() {
		var (
			ctx    context.Context
			req    *pb.ListBookRequest
			repo   *mocks.IBookRepository
			s      *implGRPCService
			book   *entities.Book
			filter *entities.ListBookFilter
			pbBook *pb.Book
		)

		BeforeEach(func() {
			ctx = context.TODO()
			repo = new(mocks.IBookRepository)
			req = &pb.ListBookRequest{
				Limit:  10,
				Cursor: 123123,
			}
			loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
			time := time.Date(2023, 07, 1, 14, 30, 40, 0, loc)
			book = &entities.Book{
				Model: entities.Model{
					ID:        234234,
					CreatedAt: time,
					UpdatedAt: time,
				},
				Title:           "Book 1",
				Author:          "Author 1",
				PublicationYear: 1991,
				Price:           2000,
				Description:     "Description 1",
				Type:            entities.BookType_Sci_fi,
			}
			pbTime := timestamppb.New(time)
			pbBook = &pb.Book{
				ID:              234234,
				Title:           "Book 1",
				Author:          "Author 1",
				PublicationYear: 1991,
				Price:           2000,
				Description:     "Description 1",
				Type:            pb.BookType_BOOK_TYPE_SCI_FI,
				CreatedAt:       pbTime,
				UpdatedAt:       pbTime,
			}
		})

		It("should return success when repo return success", func() {
			repo.On("ListBooks", ctx, 10, entities.ID(123123), filter).Return(&entities.BookPaginator{
				Total:      1,
				NextCursor: 234234,
				Items: []*entities.Book{
					book,
				},
			}, nil)
			s = &implGRPCService{
				bookRepo: repo,
			}
			resp, err := s.ListBooks(ctx, req)
			Expect(resp).ShouldNot(BeNil())
			Expect(err).Should(BeNil())
			Expect(resp.Total).Should(Equal(uint32(1)))
			Expect(resp.NextCursor).Should(Equal(uint64(234234)))
			Expect(len(resp.Entries)).Should(Equal(1))
			Expect(resp.Entries[0]).Should(Equal(pbBook))
		})

		It("should return error when repo return error", func() {
			repo.On("ListBooks", ctx, 10, entities.ID(123123), filter).Return(nil, errors.New("err ListBooks"))
			s = &implGRPCService{
				bookRepo: repo,
			}
			resp, err := s.ListBooks(ctx, req)
			Expect(resp).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
		})
	})
})
