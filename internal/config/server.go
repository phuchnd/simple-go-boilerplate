package config

import (
	"fmt"
)

const (
	ServerConfigName = "server"
)

type ServerConfig struct {
	GRPCPort int
	Name     string
	Env      string
}

func GetServerConfig() *ServerConfig {
	v := cfgProvider.viper
	a := &ServerConfig{
		GRPCPort: v.GetInt(ServerConfigName + ".grpc_port"),
		Name:     v.GetString(ServerConfigName + ".name"),
		Env:      v.GetString(ServerConfigName + ".env"),
	}
	fmt.Println(a)
	return a
}

func initServerConfig() {
	cfgProvider.viper.SetDefault(ServerConfigName, map[string]interface{}{
		"grpc_port": 3033,
		"name":      "product-service",
		"env":       "local",
	})
}
