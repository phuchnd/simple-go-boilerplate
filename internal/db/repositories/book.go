package repositories

import (
	"context"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/go-sql-driver/mysql"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"github.com/phuchnd/simple-go-boilerplate/internal/db/repositories/entities"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
	"time"
)

//go:generate mockery --name=IBookRepository --case=snake --disable-version-string
type IBookRepository interface {
	Create(ctx context.Context, data *entities.Book) (*entities.Book, error)
	GetByID(ctx context.Context, id entities.ID) (*entities.Book, error)
	ListBooks(ctx context.Context, limit int, cursor entities.ID, filter *entities.ListBookFilter) (*entities.BookPaginator, error)
}

type bookRepository struct {
	config      *config.MySQLConfig
	db          *gorm.DB
	idGenerator entities.IDGenerator
}

func NewBookRepository(db *gorm.DB, idGenerator entities.IDGenerator, cfg *config.MySQLConfig) (IBookRepository, error) {
	return &bookRepository{
		db:          db,
		idGenerator: idGenerator,
		config:      cfg,
	}, nil
}

func (r *bookRepository) Create(ctx context.Context, book *entities.Book) (*entities.Book, error) {
	db := r.db.WithContext(ctx)

	if book.ID == 0 {
		book.SetID(r.idGenerator.Next())
	}

	err := observeWithRetry(ctx, "bookRepository.Create", r.config, func(ctx context.Context) error {
		result := db.Create(&book)
		return result.Error
	})
	if err != nil {
		msqlErr, ok := err.(*mysql.MySQLError)
		if !ok || msqlErr.Number != 1062 {
			return nil, err
		}
	}

	resultBook := &entities.Book{}
	result := db.First(book, book.ID)

	return resultBook, result.Error

}

func (r *bookRepository) GetByID(ctx context.Context, id entities.ID) (*entities.Book, error) {
	db := r.db.WithContext(ctx)

	book := &entities.Book{}
	err := observeWithRetry(ctx, "bookRepository.GetByID", r.config, func(ctx context.Context) error {
		result := db.First(book, id)
		return result.Error
	})

	return book, err
}

func (r *bookRepository) ListBooks(ctx context.Context, limit int, cursor entities.ID, filter *entities.ListBookFilter) (*entities.BookPaginator, error) {
	db := r.db.WithContext(ctx)
	query := db
	orderByField := "id"
	orderByDirection := "ASC"

	var queryFields []string
	var queryConditions []interface{}

	if filter != nil {
		if len(filter.BookType) != 0 {
			queryFields = append(queryFields, "type IN ?")
			queryConditions = append(queryConditions, filter.BookType)
		}
		if filter.Author != "" {
			queryFields = append(queryFields, "author = ?")
			queryConditions = append(queryConditions, filter.Author)
		}
		if filter.OrderBy != "" {
			orderByField = filter.OrderBy
		}
		if filter.OrderByDirection != "" {
			orderByDirection = filter.OrderByDirection
		}
	}

	queryStrWithoutCursor := strings.Join(queryFields, " AND ")

	books := make([]*entities.Book, 0, limit)
	var total int64
	err := observeWithRetry(ctx, "bookRepository.ListBooks", r.config, func(ctx context.Context) error {
		countResult := query.Model(&entities.Book{}).Where(queryStrWithoutCursor, queryConditions...).Count(&total)
		if countResult.Error != nil {
			return countResult.Error
		}
		if cursor != 0 {
			queryFields = append(queryFields, "id > ?")
			queryConditions = append(queryConditions, cursor)
		}
		queryStrWithCursor := strings.Join(queryFields, " AND ")

		result := query.Where(queryStrWithCursor, queryConditions...).
			Order(fmt.Sprintf("%s %s", orderByField, orderByDirection)).
			Limit(limit).
			Find(&books)

		return result.Error
	})

	if err != nil {
		return nil, err
	}

	nextCursor := entities.ID(0)
	if len(books) > 0 {
		nextCursor = books[len(books)-1].ID
	}
	return &entities.BookPaginator{
		Total:      int(total),
		NextCursor: nextCursor,
		Items:      books,
	}, nil
}

func observeWithRetry(ctx context.Context, endpoint string, cfg *config.MySQLConfig, fn func(ctx context.Context) error) (err error) {
	logger := logging.FromContext(ctx)

	defer func() {
		if err != nil {
			logger.WithFields(logrus.Fields{
				"external_endpoint": endpoint,
				"error":             err,
			}).Errorf(fmt.Sprintf("%s: failed", endpoint))
		} else {
			logger.Infof("%s: success", endpoint)
		}
	}()

	return retry.Do(func() error {
		err := fn(ctx)
		if err != nil {
			logger.WithField("error", err).Errorf("%s: failed (inner attempt)", endpoint)
		}

		return err
	},
		retry.Attempts(uint(cfg.MaxRetries)),
		retry.Delay(time.Duration(cfg.BackoffDelaysMs)*time.Millisecond),
		retry.LastErrorOnly(true),
	)
}
