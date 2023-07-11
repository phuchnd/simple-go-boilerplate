package http

import (
	"fmt"
	"github.com/phuchnd/simple-go-boilerplate/internal/service/http/entities"
	"github.com/phuchnd/simple-go-boilerplate/server/http/dto"
)

func ListBookResponseFromEntitiesToDTO(in *entities.ListBookResponse) *dto.ListBookResponse {
	if in == nil {
		return nil
	}
	books := make([]*dto.Book, 0, len(in.Entries))
	if in.Entries != nil {
		for _, entry := range in.Entries {
			books = append(books, BookFromEntitiesToDTO(entry))
		}
	}
	return &dto.ListBookResponse{
		Entries:    books,
		Total:      in.Total,
		NextCursor: fmt.Sprintf("%d", in.NextCursor),
	}
}

func BookFromEntitiesToDTO(in *entities.Book) *dto.Book {
	if in == nil {
		return nil
	}
	out := &dto.Book{
		ID:              fmt.Sprintf("%d", in.ID),
		Title:           in.Title,
		Author:          in.Author,
		PublicationYear: in.PublicationYear,
		Price:           in.Price,
		Description:     in.Description,
		Type:            BookTypeEntitiesToDTO(in.Type),
		CreatedAt:       in.CreatedAt,
		UpdatedAt:       in.UpdatedAt,
	}
	return out
}

func BookTypeEntitiesToDTO(in entities.BookType) dto.BookType {
	mapValue := map[entities.BookType]dto.BookType{
		entities.BookType_BOOK_TYPE_UNKNOWN:    dto.BookType_BOOK_TYPE_UNKNOWN,
		entities.BookType_BOOK_TYPE_FICTION:    dto.BookType_BOOK_TYPE_FICTION,
		entities.BookType_BOOK_TYPE_NONFICTION: dto.BookType_BOOK_TYPE_NONFICTION,
		entities.BookType_BOOK_TYPE_SCI_FI:     dto.BookType_BOOK_TYPE_SCI_FI,
		entities.BookType_BOOK_TYPE_MYSTERY:    dto.BookType_BOOK_TYPE_MYSTERY,
		entities.BookType_BOOK_TYPE_THRILLER:   dto.BookType_BOOK_TYPE_THRILLER,
	}
	if v, ok := mapValue[in]; ok {
		return v
	}
	return dto.BookType_BOOK_TYPE_UNKNOWN
}
