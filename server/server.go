package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type server struct {
	restServer GinI
}

type Server interface {
	Start()
	Close()
}

func New(httpApiPort string) (Server, error) {
	log.Info("New Server created")

	restServer := NewGin(httpApiPort)

	s := &server{
		restServer: restServer,
	}
	return s, nil
}
func c(context *gin.Context) {
	log.Warn("oke")
}

func (s *server) Start() {
	s.restServer.AddHandler(http.MethodGet, "/hi", c)
	log.Infof("%v", s.restServer.GetRoutes())
	log.Info("Server started")
	go s.restServer.Start()
}

func (s *server) Close() {
	log.Info("Closing server")
	s.restServer.Close(context.Background())
}
