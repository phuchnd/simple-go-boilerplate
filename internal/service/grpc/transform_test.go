package grpc

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/phuchnd/simple-go-boilerplate/internal/db/repositories/entities"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var _ = Describe("Transform", func() {
	Context("BookTypeFromDBEntitiesToPB", func() {
		It("should return expected result", func() {
			out := BookTypeFromDBEntitiesToPB(entities.BookType_Unknown)
			Expect(out).Should(Equal(pb.BookType_BOOK_TYPE_UNKNOWN))
			out = BookTypeFromDBEntitiesToPB(entities.BookType_Fiction)
			Expect(out).Should(Equal(pb.BookType_BOOK_TYPE_FICTION))
			out = BookTypeFromDBEntitiesToPB(entities.BookType_Non_Fiction)
			Expect(out).Should(Equal(pb.BookType_BOOK_TYPE_NONFICTION))
			out = BookTypeFromDBEntitiesToPB(entities.BookType_Sci_fi)
			Expect(out).Should(Equal(pb.BookType_BOOK_TYPE_SCI_FI))
			out = BookTypeFromDBEntitiesToPB(entities.BookType_Mystery)
			Expect(out).Should(Equal(pb.BookType_BOOK_TYPE_MYSTERY))
			out = BookTypeFromDBEntitiesToPB(entities.BookType_Thriller)
			Expect(out).Should(Equal(pb.BookType_BOOK_TYPE_THRILLER))
		})
		It("should return default value", func() {
			out := BookTypeFromDBEntitiesToPB(entities.BookType(""))
			Expect(out).Should(Equal(pb.BookType_BOOK_TYPE_UNKNOWN))
			out = BookTypeFromDBEntitiesToPB(entities.BookType("123123"))
			Expect(out).Should(Equal(pb.BookType_BOOK_TYPE_UNKNOWN))
		})
	})
	Context("BookFromDBEntitiesToPB", func() {
		It("should return nil when input nil", func() {
			out := BookFromDBEntitiesToPB(nil)
			Expect(out).Should(BeNil())
		})
		It("should return expected result", func() {
			loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
			time := time.Date(2023, 07, 1, 14, 30, 40, 0, loc)
			out := BookFromDBEntitiesToPB(&entities.Book{
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
			})
			Expect(out).ShouldNot(BeNil())
			pbTime := timestamppb.New(time)
			Expect(out).Should(Equal(&pb.Book{
				ID:              234234,
				Title:           "Book 1",
				Author:          "Author 1",
				PublicationYear: 1991,
				Price:           2000,
				Description:     "Description 1",
				Type:            pb.BookType_BOOK_TYPE_SCI_FI,
				CreatedAt:       pbTime,
				UpdatedAt:       pbTime,
			}))
		})
	})
})
