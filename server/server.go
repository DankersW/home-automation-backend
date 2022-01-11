package server

import (
	log "github.com/sirupsen/logrus"
)

type server struct {
	addr string
}

type Server interface {
	Start()
}

func New(addr string) (Server, error) {
	log.Info("New Server created")
	s := &server{
		addr: addr,
	}
	return s, nil
}

func (s *server) Start() {
	log.Info("Server started")
}
