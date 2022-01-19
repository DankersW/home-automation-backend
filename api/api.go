package api

import (
	"context"
	"net/http"

	"github.com/dankersw/home-automation-backend/db"
	"github.com/dankersw/home-automation-backend/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type api struct {
	dbi db.Db
}

type Api interface {
	GetDbCollectionNames(*gin.Context)
	GetSensorData(*gin.Context)
}

func New(ctx context.Context, config models.Config) Api {
	dbi, err := db.New(ctx, config)
	if err != nil {
		log.Errorf("DB setup error. %s", err.Error())
	}

	a := &api{
		dbi: dbi,
	}
	return a
}

func fail(gc *gin.Context, msg string, err error) {
	Reply(gc, http.StatusInternalServerError, nil, err)
	log.Errorf("%s. %s", msg, err.Error())
}

func Reply(gc *gin.Context, code int, data interface{}, err error) {
	var content gin.H
	if err != nil {
		content = gin.H{"error": err.Error()}
	} else {
		content = gin.H{"data": data}
	}
	gc.JSON(code, content)
}
