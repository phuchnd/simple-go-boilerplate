package middlewares

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	"github.com/phuchnd/simple-go-boilerplate/internal/tracing"
	"google.golang.org/grpc/metadata"
	"strings"
)

func Tracing() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			tracingMetadata := tracing.FromContext(ctx)
			if tracingMetadata == nil {
				if metaData, ok := metadata.FromIncomingContext(ctx); ok {
					reqID := ""
					if v, ok := metaData[tracing.DefaultContextKeyRequestID]; ok {
						reqID = strings.Join(v, "")
					}
					if reqID == "" {
						reqID = uuid.New().String()
					}
					tracingMetadata = &tracing.RequestTracing{
						RequestID: reqID,
					}
				}
			}
			logger := logging.FromContext(ctx).WithField("request_id", tracingMetadata.RequestID)
			requestCtx := tracing.NewContext(ctx, tracingMetadata)
			requestCtx = logging.NewContext(ctx, logger)

			return next(requestCtx, request)
		}
	}
}
