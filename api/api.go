package api

import (
	"context"
	"net/http"

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

func (a *api) GetDbCollectionNames(gc *gin.Context) {
	names, err := a.dbi.FetchCollectionNames()
	if err != nil {
		response(gc, http.StatusInternalServerError, nil, err)
	} else {
		response(gc, http.StatusOK, names, err)
	}
}

func response(gc *gin.Context, code int, data interface{}, err error) {
	var content gin.H
	if err != nil {
		content = gin.H{"error": err.Error()}
	} else {
		content = gin.H{"data": data}
	}
	gc.JSON(code, content)
}
