package logging

import (
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
)

func NewLogger(cfgProvider config.IConfig) Logger {
	cfg := cfgProvider.GetServerConfig()
	logger := NewStructuredLogger(cfg)
	return logger.WithField("service", cfg.Name)
}
