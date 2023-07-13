package test

import (
	"context"
	"errors"
	. "github.com/onsi/ginkgo/v2"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"github.com/phuchnd/simple-go-boilerplate/internal/config/mocks"
	"github.com/phuchnd/simple-go-boilerplate/internal/db/migrations"
	mysqldb "github.com/phuchnd/simple-go-boilerplate/internal/db/mysql"
	"github.com/phuchnd/simple-go-boilerplate/internal/db/repositories/entities"
	"github.com/phuchnd/simple-go-boilerplate/internal/generators"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type DatabaseSetup struct {
	setup    *DockerComposeSetup
	useSetup bool

	Logger      logging.Logger
	Config      *config.DBConfig
	DB          *gorm.DB
	IDGenerator entities.IDGenerator
}

func NewDatabaseSetup(setup *DockerComposeSetup) *DatabaseSetup {
	return &DatabaseSetup{
		setup:    setup,
		useSetup: _viper.GetBool("use.setup"),
	}
}

func (s *DatabaseSetup) Setup() {
	t := WrapGinkgoT(GinkgoT())

	s.Config = &config.DBConfig{
		MySQL: &config.MySQLConfig{
			Host:            "localhost",
			Port:            s.setup.DBPort,
			Username:        "root",
			Password:        "secret",
			Database:        "example",
			MaxIdleConns:    10,
			MaxOpenConns:    100,
			MaxRetries:      3,
			BackoffDelaysMs: 100,
		},
	}
	cfgProvider := new(mocks.IConfig)
	cfgProvider.On("GetServerConfig").Return(&config.ServerConfig{
		Name: "service-test",
	})

	s.Logger = logging.NewLogger(cfgProvider)

	db, err := mysqldb.NewDB(s.Config.MySQL)
	if err != nil {
		t.Error("failed to initialize database")
		return
	}
	s.DB = db

	migrator, err := migrations.NewMigrator(db)
	if err != nil {
		t.Error("failed to initialize database")
		return
	}
	_, err = migrator.Up()

	if err != nil {
		t.Error("failed to migrate database")
		return
	}

	sfGenerator, err := generators.NewSnowflakeIDGenerator()
	if err != nil {
		t.Error("failed to create Snowflake generator")
		return
	}

	s.IDGenerator, err = entities.NewIDGenerator(sfGenerator)
	if err != nil {
		t.Error("failed to create Snowflake generator")
		return
	}
}

func WaitForMySQL(connection string, timeOut time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	readyChan := make(chan struct{})

	go func() {
		for {
			_, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
			if err != nil {
				time.Sleep(time.Second)
			} else {
				close(readyChan)
				return
			}
		}
	}()

	select {
	case <-readyChan:
		return nil
	case <-ctx.Done():
		return errors.New("timeout waiting for SQL service")
	}
}
