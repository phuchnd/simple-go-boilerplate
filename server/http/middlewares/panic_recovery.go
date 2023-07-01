package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	"github.com/phuchnd/simple-go-boilerplate/server/http/dto"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
)

func PanicRecovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		var errorMessage string
		if err, ok := recovered.(string); ok {
			errorMessage = err
		}
		c.JSON(http.StatusInternalServerError, &dto.Error{
			Code:  http.StatusInternalServerError,
			Error: errorMessage,
		})
		c.AbortWithStatus(http.StatusInternalServerError)
		stack := debug.Stack()

		logging.FromContext(c.Request.Context()).
			WithField("log_type", "panic-recovery").
			WithFields(logrus.Fields{
				"error":     errorMessage,
				"stack":     string(stack),
				"recovered": recovered,
			}).Panic(fmt.Sprintf("panic recovery in %s %s", c.Request.Method, c.Request.URL.Path))
	})
}
