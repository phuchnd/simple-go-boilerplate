package grpc

import (
	"context"
	retry "github.com/avast/retry-go"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	"github.com/phuchnd/simple-go-boilerplate/internal/tracing"
	"google.golang.org/grpc"
	"time"
)

func propagateAndObservationUnaryClientInterceptor(cfg *TransportConfig) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		newCtx := tracing.PropagateRequestIDToContext(ctx)
		var err error

		err = retry.Do(func() error {
			err := invoker(newCtx, method, req, reply, cc, opts...)
			if err != nil {
				logger := logging.FromContext(newCtx)
				logger.WithField("error", err).Warnf("%s: inner attempt failed", method)
			}
			return err
		},
			retry.Attempts(uint(cfg.MaxRetries)),
			retry.Delay(time.Duration(cfg.BackoffDelaysMs)*time.Millisecond),
			retry.LastErrorOnly(true),
		)
		return err
	}
}
