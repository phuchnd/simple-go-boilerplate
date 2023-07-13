package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	DBConfigName    = "db"
	MySQLConfigName = "mysql"
)

type DBConfig struct {
	MySQL *MySQLConfig
}

type MySQLConfig struct {
	Host            string
	Port            int
	Username        string
	Password        string
	Database        string
	MaxIdleConns    int
	MaxOpenConns    int
	MaxRetries      int
	BackoffDelaysMs int
}

func (c *configImpl) GetDBConfig() *DBConfig {
	var mySQLConfigPath = fmt.Sprintf("%s.%s", DBConfigName, MySQLConfigName)
	return &DBConfig{
		MySQL: &MySQLConfig{
			Host:         c.viper.GetString(mySQLConfigPath + ".host"),
			Port:         c.viper.GetInt(mySQLConfigPath + ".port"),
			Username:     c.viper.GetString(mySQLConfigPath + ".username"),
			Password:     c.viper.GetString(mySQLConfigPath + ".password"),
			Database:     c.viper.GetString(mySQLConfigPath + ".database"),
			MaxIdleConns: c.viper.GetInt(mySQLConfigPath + ".max_idle_conns"),
			MaxRetries:   c.viper.GetInt(mySQLConfigPath + ".max_retry"),
			MaxOpenConns: c.viper.GetInt(mySQLConfigPath + ".backoff_delays_ms"),
		},
	}
}

func initDBConfig(v *viper.Viper) {
	v.SetDefault(DBConfigName, map[string]interface{}{
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
