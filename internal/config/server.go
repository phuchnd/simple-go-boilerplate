package config

import "github.com/spf13/viper"

const (
	ServerConfigName = "server"
	LocalEnv         = "local"
)

type ServerConfig struct {
	HTTPPort    int
	HTTPDocPort int
	GRPCPort    int
	Name        string
	Env         string
}

func (c *configImpl) GetServerConfig() *ServerConfig {
	return &ServerConfig{
		HTTPPort:    c.viper.GetInt(ServerConfigName + ".http_port"),
		HTTPDocPort: c.viper.GetInt(ServerConfigName + ".http_doc_port"),
		GRPCPort:    c.viper.GetInt(ServerConfigName + ".grpc_port"),
		Name:        c.viper.GetString(ServerConfigName + ".name"),
		Env:         c.viper.GetString(ServerConfigName + ".env"),
	}
}

func initServerConfig(v *viper.Viper) {
	v.SetDefault(ServerConfigName, map[string]interface{}{
		"http_port":     8088,
		"http_doc_port": 8099,
		"grpc_port":     3033,
		"name":          "simple-go-service",
		"env":           LocalEnv,
	})
}
