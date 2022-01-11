package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	API_GROUP = "/api"
)

type ginWF struct {
	addr   string
	engine *gin.Engine
	router *gin.RouterGroup
}

type GinI interface {
	Start()
	Close()
}

func NewGin(port string) (GinI, error) {
	engine := gin.Default()
	engine.Use(cors.Default())
	router := engine.Group(API_GROUP)

	g := &ginWF{
		addr:   port,
		engine: engine,
		router: router,
	}
	return g, nil
}

func (g *ginWF) Start() {
	log.Info("Gin HTTP web API framework started on port %s", g.addr)
	log.Warn(g.engine.Routes())
	g.router.GET("/temp/info")
	log.Warn(g.engine.Routes())
	g.engine.Run(g.addr)
}

func (g *ginWF) Close() {
	log.Info("gin stopped")
}
