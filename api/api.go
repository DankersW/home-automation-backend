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
	Reply(gc, http.StatusOK, "hi", nil)
	/*
		filter := generate_timestamp_filter(7, 0)
		cursor := mongo_read("device_sensor_data", filter)
		cursor.Close(context.TODO())

		sensor_data := []SensorData{}
		for cursor.Next(db_ctx) {
			var document_item bson.M
			err := cursor.Decode(&document_item)
			if err != nil {
				log.Fatal(err)
			}

			var sensor_item SensorData
			sensor_item.Device_id = cast_to_string(document_item["device_id"])
			sensor_item.Timestamp = document_item["timestamp"].(primitive.DateTime)
			sensor_item.Temp = cast_to_float32(document_item["temperature"])
			sensor_item.Humi = cast_to_float32(document_item["humidity"])
			sensor_data = append(sensor_data, sensor_item)
		}

		return data
	*/
}
