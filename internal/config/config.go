package config

import (
	"github.com/spf13/viper"
	"strings"
)

//go:generate mockery --name=IConfig --case=snake --disable-version-string
type IConfig interface {
	GetServerConfig() *ServerConfig
	GetDBConfig() *DBConfig
	GetBookConfig() *BookConfig
	GetCronSimpleExampleConfig() *CronConfig
	GetCronHealthCheckConfig() *CronConfig
}

type configImpl struct {
	viper *viper.Viper
}

func NewConfig() IConfig {
	cfgProvider := &configImpl{
		viper: viper.New(),
	}
	initViper(cfgProvider.viper)
	_ = cfgProvider.viper.ReadInConfig()
	initServerConfig(cfgProvider.viper)
	initDBConfig(cfgProvider.viper)
	initBookConfig(cfgProvider.viper)
	initCronSimpleExampleConfig(cfgProvider.viper)
	initHealthCheckConfig(cfgProvider.viper)
	return cfgProvider
}

func initViper(v *viper.Viper) {
	v.SetConfigName("app-config")
	v.SetConfigType("yaml")
	v.AddConfigPath("$APP_CONFIG_DIR")
	v.AddConfigPath(".")
	v.AddConfigPath("$HOME")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()
}
