package entities

type ListBookRequest struct {
	Limit  uint32
	Cursor uint64
}

type ListBookResponse struct {
	Entries    []*Book
	Total      uint32
	NextCursor uint64
}

type Book struct {
	ID              uint64
	Title           string
	Author          string
	PublicationYear uint32
	Price           uint64
	Description     string
	Type            BookType
	CreatedAt       string
	UpdatedAt       string
}

type BookType string

const (
	BookType_BOOK_TYPE_UNKNOWN    BookType = "Unknown"
	BookType_BOOK_TYPE_FICTION    BookType = "Fiction"
	BookType_BOOK_TYPE_NONFICTION BookType = "Non-fiction"
	BookType_BOOK_TYPE_SCI_FI     BookType = "Sci-fi"
	BookType_BOOK_TYPE_MYSTERY    BookType = "Mystery"
	BookType_BOOK_TYPE_THRILLER   BookType = "Thriller"
)
