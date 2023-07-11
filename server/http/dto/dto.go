package dto

type Error struct {
	Error  string `json:"error"`
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

type ListBookRequest struct {
	Limit  uint32 `json:"limit" uri:"limit"`
	Cursor uint64 `json:"cursor"  uri:"cursor"`
}

type ListBookResponse struct {
	Entries    []*Book `json:"entries"`
	Total      uint32  `json:"total"`
	NextCursor string  `json:"next_cursor"`
}

type Book struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	Author          string   `json:"author"`
	PublicationYear uint32   `json:"publication_year"`
	Price           uint64   `json:"price"`
	Description     string   `json:"description"`
	Type            BookType `json:"type"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
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
