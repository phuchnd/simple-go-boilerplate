package middlewares

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

type requestLoggingWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func RequestLogging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		blw := &requestLoggingWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		var requestBodyBytes []byte
		if c.Request.Body != nil {
			requestBodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBodyBytes))
		c.Next()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		logging.FromContext(c.Request.Context()).
			WithField("type", "request-logging").
			WithFields(logrus.Fields{
				"request":       string(requestBodyBytes),
				"response":      blw.body.String(),
				"keys":          c.Keys,
				"timeStamp":     time.Now(),
				"latency_ms":    time.Since(start).Milliseconds(),
				"client_id":     c.ClientIP(),
				"method":        c.Request.Method,
				"status_code":   c.Writer.Status(),
				"error_message": c.Errors.ByType(gin.ErrorTypePrivate).String(),
				"body_size":     c.Writer.Size(),
				"path":          c.Request.URL.Path,
				"full_path":     path,
			}).Info(fmt.Sprintf("[%s][%d] %s %s", c.Request.Method, c.Writer.Status(), c.Request.URL.Path, time.Since(start).String()))
	}
}
