package grpc

import (
	"github.com/phuchnd/simple-go-boilerplate/internal/db/repositories/entities"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ListBookResponseFromDBEntitiesToPB(in *entities.BookPaginator) *pb.ListBookResponse {
	if in == nil {
		return nil
	}

	books := make([]*pb.Book, 0, len(in.Items))
	if in.Items != nil {
		for _, entry := range in.Items {
			books = append(books, BookFromDBEntitiesToPB(entry))
		}
	}
	return &pb.ListBookResponse{
		Entries:    books,
		Total:      uint32(in.Total),
		NextCursor: uint64(in.NextCursor),
	}
}

func BookFromDBEntitiesToPB(in *entities.Book) *pb.Book {
	if in == nil {
		return nil
	}

	out := &pb.Book{
		ID:              uint64(in.ID),
		Title:           in.Title,
		Author:          in.Author,
		PublicationYear: in.PublicationYear,
		Price:           in.Price,
		Description:     in.Description,
		Type:            BookTypeFromDBEntitiesToPB(in.Type),
		CreatedAt:       timestamppb.New(in.CreatedAt),
		UpdatedAt:       timestamppb.New(in.UpdatedAt),
	}
	return out
}

func BookTypeFromDBEntitiesToPB(in entities.BookType) pb.BookType {
	mapValue := map[entities.BookType]pb.BookType{
		entities.BookType_Unknown:     pb.BookType_BOOK_TYPE_UNKNOWN,
		entities.BookType_Fiction:     pb.BookType_BOOK_TYPE_FICTION,
		entities.BookType_Non_Fiction: pb.BookType_BOOK_TYPE_NONFICTION,
		entities.BookType_Sci_fi:      pb.BookType_BOOK_TYPE_SCI_FI,
		entities.BookType_Mystery:     pb.BookType_BOOK_TYPE_MYSTERY,
		entities.BookType_Thriller:    pb.BookType_BOOK_TYPE_THRILLER,
	}
	if v, ok := mapValue[in]; ok {
		return v
	}
	return pb.BookType_BOOK_TYPE_UNKNOWN
}
