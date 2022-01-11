package server

import (
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	API_GROUP = "/api"
	NAME      = "Gin REST server"
)

type ginWF struct {
	addr   string
	router *gin.Engine
	group  *gin.RouterGroup
	srv    *http.Server
}

type GinI interface {
	Start()
	Close(ctx context.Context)
	GetRoutes() gin.RoutesInfo
}

func NewGin(port string) GinI {
	router := gin.Default()
	router.Use(cors.Default())
	apiGroup := router.Group(API_GROUP)

	srv := &http.Server{Addr: port, Handler: router}

	g := &ginWF{
		addr:   port,
		group:  apiGroup,
		router: router,
		srv:    srv,
	}
	return g
}

func (g *ginWF) Start() {
	log.Infof("%s started on port %s", NAME, g.addr)
	if err := g.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to server %s, %s\n", NAME, err)
	}
}

func (g *ginWF) Close(ctx context.Context) {
	log.Infof("Stopping %s...", NAME)
	if err := g.srv.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to Shutdown, %s", err.Error())
	}
	log.Infof("%s stopped", NAME)
}

func (g *ginWF) GetRoutes() gin.RoutesInfo {
	return g.router.Routes()
}
