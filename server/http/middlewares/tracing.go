package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	"github.com/phuchnd/simple-go-boilerplate/internal/tracing"
)

func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		tracingMetadata := tracing.FromContext(ctx)
		if tracingMetadata == nil {
			reqID := c.GetHeader(tracing.DefaultContextKeyRequestID)
			if reqID == "" {
				reqID = uuid.New().String()
			}
			tracingMetadata = &tracing.RequestTracing{
				RequestID: reqID,
			}
		}

		logger := logging.FromContext(c.Request.Context()).WithField("request_id", tracingMetadata.RequestID)
		newCtx := tracing.NewContext(ctx, tracingMetadata)
		newCtx = logging.NewContext(ctx, logger)
		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}
}
