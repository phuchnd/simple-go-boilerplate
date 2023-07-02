package entities

type BookType string

const (
	BookType_Unknown     BookType = "Unknown"
	BookType_Fiction     BookType = "Fiction"
	BookType_Non_Fiction BookType = "Non-fiction"
	BookType_Sci_fi      BookType = "Sci-fi"
	BookType_Mystery     BookType = "Mystery"
	BookType_Thriller    BookType = "Thriller"
)

type Book struct {
	Model

	Title           string
	Author          string
	PublicationYear uint32
	Price           uint64
	Description     string
	Type            BookType `gorm:"default:unknown"`
}

type ListBookFilter struct {
	Author           string
	BookType         []BookType
	OrderBy          string
	OrderByDirection string
}

type BookPaginator struct {
	Total      int
	NextCursor ID
	Items      []*Book
}
