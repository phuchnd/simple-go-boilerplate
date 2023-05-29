package config

import (
	"github.com/spf13/viper"
	"strings"
)

var cfgProvider *configProvider

type configProvider struct {
	registry IConfigRegistry
	viper    *viper.Viper
}

func init() {
	cfgProvider = &configProvider{
		registry: NewConfigRegistry(),
		viper:    viper.New(),
	}
	initViper(cfgProvider.viper)
	_ = cfgProvider.viper.ReadInConfig()
	cfgProvider.registry.SetDefaultConfigs(cfgProvider.viper)
	initServerConfig()
	initDBConfig()
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
