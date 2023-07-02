package test

import (
	"github.com/phayes/freeport"
	"strings"

	//"github.com/phayes/freeport"

	"github.com/spf13/viper"
)

var (
	_viper = viper.New()
)

func init() {
	_viper.SetDefault("use.setup", true)
	_viper.SetDefault("db.port", freeport.GetPort())

	_viper.SetEnvPrefix("TEST")
	_viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	_viper.AutomaticEnv()
}
