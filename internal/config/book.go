package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	ExternalConfigName = "external"
	BookConfigName     = "book"
)

type BookConfig struct {
	Host            string
	Port            int
	MaxRetries      int
	BackoffDelaysMs int
}

func (c *configImpl) GetBookConfig() *BookConfig {
	var bookConfigPath = fmt.Sprintf("%s.%s", ExternalConfigName, BookConfigName)
	return &BookConfig{
		Host:            c.viper.GetString(bookConfigPath + ".host"),
		Port:            c.viper.GetInt(bookConfigPath + ".port"),
		MaxRetries:      c.viper.GetInt(bookConfigPath + ".max_retries"),
		BackoffDelaysMs: c.viper.GetInt(bookConfigPath + ".backoff_delays_ms"),
	}
}

func initBookConfig(v *viper.Viper) {
	v.SetDefault(DBConfigName, map[string]interface{}{
		MySQLConfigName: map[string]interface{}{
			"host":              "localhost",
			"port":              3306,
			"max_retries":       3,
			"backoff_delays_ms": 100,
		},
	})
}
