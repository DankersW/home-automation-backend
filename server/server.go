package server

import (
	"context"

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

func (s *server) Start() {
	log.Info("Server started")
	go s.restServer.Start()
}

func (s *server) Close() {
	log.Info("Closing server")
	s.restServer.Close(context.Background())
}
