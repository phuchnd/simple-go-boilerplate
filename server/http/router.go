package http

import (
	"github.com/gin-gonic/gin"
	"github.com/phuchnd/simple-go-boilerplate/internal/service/http/entities"
	"github.com/phuchnd/simple-go-boilerplate/server/http/dto"
	"github.com/phuchnd/simple-go-boilerplate/server/http/middlewares"
	"net/http"
)

func (s *httpServerImpl) initRouter() *gin.Engine {
	r := gin.New()

	// Heath Check Router
	r.GET("/health", func(c *gin.Context) {
		if s.isReady {
			c.JSON(http.StatusOK, map[string]string{
				"message": "service ready",
			})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
				"error": "service not ready",
			})
		}
		return
	})

	// Logic Router
	v0 := r.Group("/v0")
	v0.Use(
		middlewares.Tracing(),
		middlewares.RequestLogging(),
	)
	{
		v0.GET("/books", s.listBookV0)
	}

	return r
}

func (s *httpServerImpl) listBookV0(c *gin.Context) {
	var req dto.ListBookRequest
	if err := c.BindUri(&req); err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}
	ctx := c.Request.Context()
	resp, err := s.handler.ListBooks(ctx, &entities.ListBookRequest{
		Limit:  req.Limit,
		Cursor: req.Cursor,
	})
	if err != nil {
		s.handleResponseError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (i *httpServerImpl) handleResponseError(c *gin.Context, err error) {
	// Todo parsing internal error into server error and return
	c.AbortWithStatusJSON(http.StatusInternalServerError, &dto.Error{
		Error: err.Error(),
		Code:  http.StatusInternalServerError,
	})
}
