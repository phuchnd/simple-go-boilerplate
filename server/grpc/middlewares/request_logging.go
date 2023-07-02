package middlewares

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc/endpoints"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func RequestLogging() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			logger := logging.FromContext(ctx).WithField("type", "request-logging")

			metadata, ok := endpoints.GetEndpointMetadataFromContext(ctx)
			if ok {
				logger = logger.WithField("endpoint", metadata.Name)
			}
			requestCtx := logging.NewContext(ctx, logger)
			start := time.Now()
			defer func() {
				code := status.Code(err)
				args := logrus.Fields{
					"method":        metadata.Name,
					"request":       request,
					"response":      response,
					"timestamp":     time.Now(),
					"latency_ms":    time.Since(start).Milliseconds(),
					"status_code":   code.String(),
					"error_message": err,
				}

				if code == codes.OK {
					logger.WithFields(args).Infof("%s: %s", metadata.Name, code.String())
				} else {
					logger.WithFields(args).Errorf("%s: %s", metadata.Name, code.String())
				}
			}()

			return next(requestCtx, request)
		}
	}
}
