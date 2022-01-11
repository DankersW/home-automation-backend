package server

import (
	log "github.com/sirupsen/logrus"
)

type server struct {
	httpApi GinI
}

type Server interface {
	Start()
	Close()
}

func New(httpApiPort string) (Server, error) {
	log.Info("New Server created")

	g, err := NewGin(httpApiPort)
	if err != nil {
		log.Error("Failed to create Gin server, %s", err.Error())
		return nil, err
	}

	s := &server{
		httpApi: g,
	}
	return s, nil
}

func (s *server) Start() {
	log.Info("Server started")
	s.httpApi.Start()
}

func (s *server) Close() {
	log.Info("Closing server")
	s.httpApi.Close()
}
