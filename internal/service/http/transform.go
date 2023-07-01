package http

import (
	"github.com/phuchnd/simple-go-boilerplate/internal/service/http/entities"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc/pb"
	"github.com/phuchnd/simple-go-boilerplate/utils"
	"time"
)

func ListBookResponseFromPBToEntities(in *pb.ListBookResponse) *entities.ListBookResponse {
	if in == nil {
		return nil
	}
	books := make([]*entities.Book, 0, len(in.Entries))
	if in.Entries != nil {
		for _, entry := range in.Entries {
			books = append(books, BookFromPBToEntities(entry))
		}
	}
	return &entities.ListBookResponse{
		Entries:    books,
		Total:      in.Total,
		NextCursor: in.NextCursor,
	}
}

func BookFromPBToEntities(in *pb.Book) *entities.Book {
	if in == nil {
		return nil
	}
	out := &entities.Book{
		ID:              in.ID,
		Title:           in.Title,
		Author:          in.Author,
		PublicationYear: in.PublicationYear,
		Price:           in.Price,
		Description:     in.Description,
		Type:            BookTypeFromPBToEntities(in.Type),
	}
	// Todo dynamic timezone
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	out.CreatedAt = utils.ConvertTimeStr(in.CreatedAt, loc)
	out.UpdatedAt = utils.ConvertTimeStr(in.UpdatedAt, loc)

	return out
}

func BookTypeFromPBToEntities(in pb.BookType) entities.BookType {
	mapValue := map[pb.BookType]entities.BookType{
		pb.BookType_BOOK_TYPE_UNKNOWN:    entities.BookType_BOOK_TYPE_UNKNOWN,
		pb.BookType_BOOK_TYPE_FICTION:    entities.BookType_BOOK_TYPE_FICTION,
		pb.BookType_BOOK_TYPE_NONFICTION: entities.BookType_BOOK_TYPE_NONFICTION,
		pb.BookType_BOOK_TYPE_SCI_FI:     entities.BookType_BOOK_TYPE_SCI_FI,
		pb.BookType_BOOK_TYPE_MYSTERY:    entities.BookType_BOOK_TYPE_MYSTERY,
		pb.BookType_BOOK_TYPE_THRILLER:   entities.BookType_BOOK_TYPE_THRILLER,
	}
	if v, ok := mapValue[in]; ok {
		return v
	}
	return entities.BookType_BOOK_TYPE_UNKNOWN
}
