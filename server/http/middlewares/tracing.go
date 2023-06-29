package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
)

func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		reqID := c.GetHeader("x-request-id")

		logger := logging.FromContext(c.Request.Context()).WithField("request_id", reqID)
		newCtx := logging.NewContext(ctx, logger)
		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}
}
