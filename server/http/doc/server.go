package doc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
)

type DocumentServer struct {
	engine *gin.Engine
	cfg    *config.ServerConfig
}

func NewDocumentServer(cfgProvider config.IConfig) *DocumentServer {
	r := gin.Default()

	r.Static("/internal/static", "internal/static")
	r.LoadHTMLGlob("internal/templates/*.html")
	r.GET("/api-doc", func(c *gin.Context) {
		c.HTML(200, "api-doc.html", gin.H{})
	})

	return &DocumentServer{
		engine: r,
		cfg:    cfgProvider.GetServerConfig(),
	}
}

func (t *DocumentServer) Start() error {
	return t.engine.Run(fmt.Sprintf(":%d", t.cfg.HTTPDocPort))
}
