package server

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type RestRoute struct {
	method   string
	uri      string
	callback ginCallback
}
type RestRoutes []RestRoute

type server struct {
	restServer GinI
}

type Server interface {
	Start()
	Close()
}

func New(restServerPort string) (Server, error) {
	log.Info("New Server created")

	restServer := NewGin(restServerPort)
	handlers := NewHandlers(restServer.GetRoutes)
	restRoutes := handlers.getRestRoutes()
	s := &server{
		restServer: restServer,
	}
	s.restServer.AddRoutes(restRoutes)

	return s, nil
}

func (s *server) Start() {

	log.Infof("%v", s.restServer.GetRoutes())

	log.Info("Server started")
	go s.restServer.Start()
}

func (s *server) Close() {
	log.Info("Closing server")
	s.restServer.Close(context.Background())
}
