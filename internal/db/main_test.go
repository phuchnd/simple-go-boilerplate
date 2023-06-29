package db

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	genA()
}

func genA() {
	rand.Seed(time.Now().UnixNano())

	// Possible book types
	bookTypes := []string{"Fiction", "Non-fiction", "Sci-fi", "Mystery", "Thriller"}

	for i := 0; i < 100; i++ {
		id := generateID(i)
		title := "Book Title " + strconv.Itoa(rand.Intn(1000))
		author := "Author Name " + strconv.Itoa(rand.Intn(1000))
		publicationYear := rand.Intn(123) + 1900
		price := fmt.Sprintf("%.2f", rand.Float32()*90.0+10.0)
		description := "This is a description of the book."
		bookType := bookTypes[rand.Intn(len(bookTypes))]

		// SQL statement
		fmt.Printf("INSERT INTO books (id, title, author, publication_year, price, description, type) VALUES ('%s', '%s', '%s', %d, %s, '%s', '%s');\n", id, title, author, publicationYear, price, description, bookType)
	}
}

// Simple Snowflake-like ID generator
func generateID(i int) string {
	// Time - 41 bits (millisecond precision w/ a custom epoch gives us 69 years)
	timeBits := time.Now().UnixNano() / int64(time.Millisecond) << 22
	// Node - 10 bits - gives us up to 1024 nodes
	nodeBits := int64(i%1024) << 12
	// Sequence number - 12 bits - rolls over every 4096 per node
	seqBits := int64(i / 1024)
	// Concatenate all parts
	id := timeBits | nodeBits | seqBits
	return strconv.FormatInt(id, 10)
}
