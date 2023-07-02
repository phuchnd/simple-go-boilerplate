package mysql

import (
	"fmt"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(mySQLConfig *config.MySQLConfig) (*gorm.DB, error) {
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

	return db, nil
}
