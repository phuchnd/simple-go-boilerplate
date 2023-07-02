package tests

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/phuchnd/simple-go-boilerplate/internal/db/repositories"
	"github.com/phuchnd/simple-go-boilerplate/internal/db/repositories/entities"
	"gorm.io/gorm"
	"sync"
)

var _ = Describe("repositories/BookRepository", func() {
	var (
		db         *gorm.DB
		idg        entities.IDGenerator
		repository repositories.IBookRepository
	)

	BeforeEach(func() {
		db = databaseSetup.DB
		idg = databaseSetup.IDGenerator
		repository, _ = repositories.NewBookRepository(db, idg, databaseSetup.Config.MySQL)
	})

	Describe("Create()", func() {
		var (
			bookID entities.ID
			input  *entities.Book
		)

		BeforeEach(func() {
			bookID = idg.Next()
			input = &entities.Book{
				Model: entities.Model{
					ID: bookID,
				},
				Title:           gofakeit.BookTitle(),
				Author:          gofakeit.Name(),
				PublicationYear: uint32(gofakeit.Year()),
				Price:           20000,
				Description:     "Description",
				Type:            entities.BookType_Non_Fiction,
			}
		})

		It("should create successfully", func() {
			result, err := repository.Create(context.TODO(), input)
			fmt.Println("Create ID51", input.ID)

			Expect(err).Should(BeNil())
			Expect(result).ShouldNot(BeNil())
			Expect(result.ID).ShouldNot(Equal(input.ID))
		})

		It("should fill in book id when not input", func() {
			input.ID = 0
			Expect(input.ID).Should(Equal(entities.ID(0)))

			result, err := repository.Create(context.TODO(), input)
			fmt.Println("Create ID62", input.ID)

			Expect(err).Should(BeNil())
			Expect(result).ShouldNot(BeNil())
			Expect(result.ID).ShouldNot(Equal(input.ID))
		})

		Describe("idempotency", func() {
			It("should return only one book with multiple creates", func() {
				count := 10

				var ref *entities.Book

				var deliveries = make([]*entities.Book, count)
				var errors = make([]error, count)

				mu := sync.Mutex{}

				wg := sync.WaitGroup{}
				wg.Add(count)

				for i := 0; i < count; i++ {
					idx := i
					go func() {
						defer func() {
							GinkgoRecover()
							mu.Unlock()
							wg.Done()
						}()

						newInput := *input
						newBook, err := repository.Create(context.TODO(), &newInput)

						mu.Lock()

						deliveries[idx] = newBook
						errors[idx] = err

						if ref == nil {
							ref = newBook
						}
					}()
				}

				wg.Wait()

				for i := 0; i < count; i++ {
					Expect(errors[i]).Should(BeNil())
					Expect(deliveries[i].ID).Should(Equal(ref.ID))
				}
			})
		})
	})

	Describe("GetByID()", func() {
		var (
			bookID entities.ID
			input  *entities.Book
		)

		BeforeEach(func() {
			bookID = idg.Next()
			input = &entities.Book{
				Model: entities.Model{
					ID: bookID,
				},
				Title:           gofakeit.BookTitle(),
				Author:          gofakeit.Name(),
				PublicationYear: uint32(gofakeit.Year()),
				Price:           150000,
				Description:     "Description Thriller",
				Type:            entities.BookType_Thriller,
			}
			_, _ = repository.Create(context.TODO(), input)

			fmt.Println("Create ID136", input.ID)
		})

		It("should return successfully", func() {
			result, err := repository.GetByID(context.TODO(), bookID)
			Expect(err).Should(BeNil())
			Expect(result).ShouldNot(BeNil())
			Expect(result.ID).Should(Equal(bookID))
		})
	})

	Describe("ListOrders", func() {
		Context("with filter author", func() {
			var (
				bookFiction1, bookFiction2, bookScifi *entities.Book
				ctx                                   context.Context
				limit                                 int
				author                                string
			)

			BeforeEach(func() {
				ctx = context.TODO()
				limit = 10
				author = gofakeit.Name() + gofakeit.LetterN(5)
				bookFiction1 = &entities.Book{
					Model: entities.Model{
						ID: idg.Next(),
					},
					Title:           gofakeit.BookTitle(),
					Author:          author,
					PublicationYear: uint32(gofakeit.Year()),
					Price:           90000,
					Description:     "Description Fiction",
					Type:            entities.BookType_Fiction,
				}
				bookFiction2 = &entities.Book{
					Model: entities.Model{
						ID: idg.Next(),
					},
					Title:           gofakeit.BookTitle(),
					Author:          author,
					PublicationYear: uint32(gofakeit.Year()),
					Price:           110000,
					Description:     "Description Fiction",
					Type:            entities.BookType_Fiction,
				}
				bookScifi = &entities.Book{
					Model: entities.Model{
						ID: idg.Next(),
					},
					Title:           gofakeit.BookTitle(),
					Author:          author,
					PublicationYear: uint32(gofakeit.Year()),
					Price:           80000,
					Description:     "Description Scifi",
					Type:            entities.BookType_Sci_fi,
				}
				db.Create(bookFiction1)
				db.Create(bookFiction2)
				db.Create(bookScifi)
			})

			It("should return just created books (same author)", func() {
				filter := &entities.ListBookFilter{
					Author: author,
				}
				resultBooks, err := repository.ListBooks(ctx, limit, 0, filter)
				Expect(err).Should(BeNil())
				Expect(resultBooks).ShouldNot(BeNil())
				Expect(len(resultBooks.Items)).Should(Equal(3))
				Expect(resultBooks.Total).Should(Equal(3))
				Expect(resultBooks.Items[0].ID).Should(Equal(bookFiction1.ID))
				Expect(resultBooks.Items[1].ID).Should(Equal(bookFiction2.ID))
				Expect(resultBooks.Items[2].ID).Should(Equal(bookScifi.ID))
			})

			It("should return just created books (same author) and book type", func() {
				filter := &entities.ListBookFilter{
					Author:   author,
					BookType: []entities.BookType{entities.BookType_Fiction},
				}
				resultBooks, err := repository.ListBooks(ctx, limit, 0, filter)
				Expect(err).Should(BeNil())
				Expect(resultBooks).ShouldNot(BeNil())
				Expect(len(resultBooks.Items)).Should(Equal(2))
				Expect(resultBooks.Total).Should(Equal(2))
				Expect(resultBooks.Items[0].ID).Should(Equal(bookFiction1.ID))
				Expect(resultBooks.Items[1].ID).Should(Equal(bookFiction2.ID))
			})

			It("should return all books in the expected order by", func() {
				filter := &entities.ListBookFilter{
					Author:           author,
					OrderBy:          "id",
					OrderByDirection: "desc",
				}
				resultBooks, err := repository.ListBooks(ctx, limit, 0, filter)
				Expect(err).Should(BeNil())
				Expect(resultBooks).ShouldNot(BeNil())
				Expect(len(resultBooks.Items)).Should(Equal(3))
				Expect(resultBooks.Total).Should(Equal(3))
				Expect(resultBooks.Items[0].ID).Should(Equal(bookScifi.ID))
				Expect(resultBooks.Items[1].ID).Should(Equal(bookFiction2.ID))
				Expect(resultBooks.Items[2].ID).Should(Equal(bookFiction1.ID))
			})
		})

		Context("with no filter", func() {
			var (
				ctx   context.Context
				limit int
			)

			BeforeEach(func() {
				ctx = context.TODO()
				limit = 10
			})

			It("should return all books just created", func() {
				filter := &entities.ListBookFilter{}
				resultBooks, err := repository.ListBooks(ctx, limit, 0, filter)
				Expect(err).Should(BeNil())
				Expect(resultBooks).ShouldNot(BeNil())
				Expect(len(resultBooks.Items)).Should(Equal(10))
				Expect(resultBooks.Total).Should(BeNumerically(">=", 10))
			})
		})
	})
})
