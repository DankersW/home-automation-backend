package server

import (
	"context"
	"fmt"

	"github.com/dankersw/home-automation-backend/api"
	"github.com/dankersw/home-automation-backend/models"
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

func New(ctx context.Context, config models.Config) (Server, error) {
	log.Info("New Server created")

	restServer := NewGin(fmt.Sprintf(":%d", config.Rest.Port))

	endpoints := api.New(ctx, config)

	handlers := NewRouteHandler(restServer.GetRoutes, endpoints)
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
