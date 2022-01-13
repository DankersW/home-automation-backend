package api

import (
	"context"

	"github.com/dankersw/home-automation-backend/db"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type api struct {
	dbi db.Db
}

type Api interface {
	GetDbCollectionNames(*gin.Context)
}

func New(ctx context.Context) Api {
	dbi, err := db.New(ctx)
	if err != nil {
		log.Errorf("DB setup error. %s", err.Error())
	}

	a := &api{
		dbi: dbi,
	}
	return a
}

func (a *api) GetDbCollectionNames(context *gin.Context) {
	a.dbi.FetchCollectionNames()
	a.dbi.Get("a")
	log.Info("a db call")
}
