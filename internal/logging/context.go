package logging

import (
	"context"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
)

type contextKey struct{}

var (
	defaultLogger = NewLogger(config.GetServerConfig())
)

// NewContext returns a new Context with given logger.
func NewContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

// FromContext returns the logger associated with a context.
// It returns the default Logger if no Logger exists.
func FromContext(ctx context.Context) Logger {
	if logger, ok := ctx.Value(contextKey{}).(Logger); ok {
		return logger
	}
	return defaultLogger
}
