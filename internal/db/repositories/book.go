package repositories

import (
	"context"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"github.com/phuchnd/simple-go-boilerplate/internal/db/repositories/entities"
	"github.com/phuchnd/simple-go-boilerplate/internal/generators"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

//go:generate mockery --name=IBookRepository --case=snake --disable-version-string
type IBookRepository interface {
	Save(ctx context.Context, data *entities.Book) error
	GetByID(ctx context.Context, id entities.ID) (*entities.Book, error)
	ListBooks(ctx context.Context, limit int, cursor entities.ID, filter *entities.ListBookFilter) (*entities.BookPaginator, error)
}

type bookRepository struct {
	config      *config.MySQLConfig
	db          *gorm.DB
	idGenerator entities.IDGenerator
}

func NewBookRepository() (IBookRepository, error) {
	sf, err := generators.NewSnowflakeIDGenerator()
	if err != nil {
		return nil, err
	}
	idGenerator, err := entities.NewIDGenerator(sf)
	if err != nil {
		return nil, err
	}

	dbConfig := config.GetDBConfig()
	mySQLConfig := dbConfig.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mySQLConfig.Username, mySQLConfig.Password, mySQLConfig.Host, mySQLConfig.Port, mySQLConfig.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(mySQLConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(mySQLConfig.MaxOpenConns)

	return &bookRepository{
		db:          db,
		idGenerator: idGenerator,
		config:      mySQLConfig,
	}, nil
}

func (r *bookRepository) Save(ctx context.Context, book *entities.Book) error {
	db := r.db.WithContext(ctx)

	return observeWithRetry(ctx, "bookRepository.Save", r.config, func(ctx context.Context) error {
		result := db.Model(&book).Updates(&book)
		return result.Error
	})
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
	bookByField := "id"
	bookByDirection := "ASC"

	var queryFields []string
	var queryConditions []interface{}

	if filter != nil {
		if filter.BookType != "" {
			queryFields = append(queryFields, "type IN ?")
			queryConditions = append(queryConditions, filter.BookType)
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
			Order(fmt.Sprintf("%s %s", bookByField, bookByDirection)).
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
