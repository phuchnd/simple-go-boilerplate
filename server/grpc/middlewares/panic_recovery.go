package middlewares

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime/debug"
)

func PanicRecovery() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func() {
				if r := recover(); r != nil {
					stack := debug.Stack()
					logging.FromContext(ctx).
						WithFields(logrus.Fields{
							"log_type": "panic-recovery",
							"data":     r,
							"stack":    string(stack),
						}).
						Error("recovered from panic")

					response = nil
					err = status.Errorf(codes.Internal, "internal server error: %v", r)
				}
			}()

			return next(ctx, request)
		}
	}
}
