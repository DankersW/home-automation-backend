package api

import (
	log "github.com/sirupsen/logrus"
)

type api struct {
}

type Api interface {
}

func New() Api {
	a := &api{}
	return a
}

func (a *api) dbCall() {
	log.Info("a db call")
}
