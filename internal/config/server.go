package config

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

func GetServerConfig() *ServerConfig {
	v := cfgProvider.viper
	return &ServerConfig{
		HTTPPort:    v.GetInt(ServerConfigName + ".http_port"),
		HTTPDocPort: v.GetInt(ServerConfigName + ".http_doc_port"),
		GRPCPort:    v.GetInt(ServerConfigName + ".grpc_port"),
		Name:        v.GetString(ServerConfigName + ".name"),
		Env:         v.GetString(ServerConfigName + ".env"),
	}
}

func initServerConfig() {
	cfgProvider.viper.SetDefault(ServerConfigName, map[string]interface{}{
		"http_port":     8088,
		"http_doc_port": 8099,
		"grpc_port":     3033,
		"name":          "simple-go-service",
		"env":           LocalEnv,
	})
}
