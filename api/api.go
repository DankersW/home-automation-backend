package api

import (
	"context"
	"net/http"

	"github.com/dankersw/home-automation-backend/db"
	"github.com/dankersw/home-automation-backend/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type api struct {
	dbi db.Db
}

type Api interface {
	GetDbCollectionNames(*gin.Context)
	GetSensorData(*gin.Context)
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

func (a *api) GetDbCollectionNames(gc *gin.Context) {
	names, err := a.dbi.FetchCollectionNames()
	if err != nil {
		Reply(gc, http.StatusInternalServerError, nil, err)
	} else {
		Reply(gc, http.StatusOK, names, err)
	}
}

func (a *api) GetSensorData(gc *gin.Context) {
	filter := a.dbi.TimestampBetween(7, 0)
	cursor, err := a.dbi.GetWithFilter("device_sensor_data", filter)
	if err != nil {
		fail(gc, "Failed to get sensor data", err)
		return
	}
	data := []models.SensorData{}
	for cursor.Next(context.TODO()) {
		var item bson.M
		if err := cursor.Decode(&item); err != nil {
			fail(gc, "Failed to decode data", err)
			return
		}
		data = append(data, models.ToSensorData(item))
	}
	Reply(gc, http.StatusOK, data, nil)
}
