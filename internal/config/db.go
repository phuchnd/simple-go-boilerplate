package config

import (
	"fmt"
)

const (
	DBConfigName    = "db"
	MySQLConfigName = "mysql"
)

type DBConfig struct {
	MySQL *MySQLConfig
}

type MySQLConfig struct {
	Host         string
	Port         int
	Username     string
	Password     string
	Database     string
	MaxIdleConns int
	MaxOpenConns int
}

func GetDBConfig() *DBConfig {
	v := cfgProvider.viper
	var mySQLConfigPath = fmt.Sprintf("%s.%s", DBConfigName, MySQLConfigName)
	a := &DBConfig{
		MySQL: &MySQLConfig{
			Host:         v.GetString(mySQLConfigPath + ".host"),
			Port:         v.GetInt(mySQLConfigPath + ".port"),
			Username:     v.GetString(mySQLConfigPath + ".username"),
			Password:     v.GetString(mySQLConfigPath + ".password"),
			Database:     v.GetString(mySQLConfigPath + ".database"),
			MaxIdleConns: v.GetInt(mySQLConfigPath + ".max_idle_conns"),
			MaxOpenConns: v.GetInt(mySQLConfigPath + ".max_open_conns"),
		},
	}
	fmt.Println(a.MySQL)
	return a
}

func initDBConfig() {
	cfgProvider.viper.SetDefault(DBConfigName, map[string]interface{}{
		MySQLConfigName: map[string]interface{}{
			"host":           "localhost",
			"port":           3306,
			"username":       "product",
			"password":       "secret",
			"database":       "product",
			"max_idle_conns": 10,
			"max_open_conns": 100,
		},
	})
}
