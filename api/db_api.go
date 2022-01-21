package api

import (
	"context"
	"net/http"

	"github.com/dankersw/home-automation-backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

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
