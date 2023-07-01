package config

import "fmt"

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

func GetBookConfig() *BookConfig {
	v := cfgProvider.viper
	var bookConfigPath = fmt.Sprintf("%s.%s", ExternalConfigName, BookConfigName)
	return &BookConfig{
		Host:            v.GetString(bookConfigPath + ".host"),
		Port:            v.GetInt(bookConfigPath + ".port"),
		MaxRetries:      v.GetInt(bookConfigPath + ".max_retries"),
		BackoffDelaysMs: v.GetInt(bookConfigPath + ".backoff_delays_ms"),
	}
}

func initBookConfig() {
	cfgProvider.viper.SetDefault(DBConfigName, map[string]interface{}{
		MySQLConfigName: map[string]interface{}{
			"host":              "localhost",
			"port":              3306,
			"max_retries":       3,
			"backoff_delays_ms": 100,
		},
	})
}
