package migrations

import (
	"embed"
	"fmt"
	"github.com/phuchnd/core-product-management/internal/config"
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//go:embed scripts/*.sql
var migrationScripts embed.FS

type IMigrator interface {
	Up() (int, error)
	Down() (int, error)
}

type migratorImpl struct {
	migrations migrate.MigrationSource
	db         *gorm.DB
}

func NewMigrator() (IMigrator, error) {
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

	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationScripts,
		Root:       "scripts",
	}

	return &migratorImpl{
		migrations: migrations,
		db:         db,
	}, nil
}

func (m *migratorImpl) Up() (int, error) {
	return m.migrate(migrate.Up)
}

func (m *migratorImpl) Down() (int, error) {
	return m.migrate(migrate.Down)
}

func (m *migratorImpl) migrate(direction migrate.MigrationDirection) (int, error) {
	db, _ := m.db.DB()
	n, err := migrate.Exec(db, "mysql", m.migrations, direction)

	if err != nil {
		return 0, err
	}

	return n, nil
}
