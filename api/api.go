package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type api struct {
}

type Api interface {
	DbCall(*gin.Context)
}

func New() Api {
	a := &api{}
	return a
}

func (a *api) DbCall(context *gin.Context) {
	log.Info("a db call")
}
