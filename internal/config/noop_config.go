package config

type noopConfigImpl struct {
}

func NewNoopConfig() IConfig {
	return &noopConfigImpl{}
}

func (c *noopConfigImpl) GetServerConfig() *ServerConfig {
	return &ServerConfig{
		Name: "simple-go-service",
		Env:  LocalEnv,
	}
}

func (c *noopConfigImpl) GetDBConfig() *DBConfig {
	return nil
}

func (c *noopConfigImpl) GetBookConfig() *BookConfig {
	return nil
}

func (c *noopConfigImpl) GetCronSimpleExampleConfig() *CronConfig {
	return nil
}
