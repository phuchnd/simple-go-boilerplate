package logging

import (
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
)

func NewLogger(cfg *config.ServerConfig) Logger {
	logger := NewStructuredLogger(cfg)
	return logger.WithField("service", cfg.Name)
}
