package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func handler_get_all_apis(context *gin.Context) {
	apis := map[string]string{
		"all":                     "/api/",
		"iot_db_collection_names": "/api/iotDbCollectionNames",
		"temperature":             "/api/temp",
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": apis,
	})
}

func handler_get_iot_db_collection_names(context *gin.Context) {
	var code int
	filter := bson.D{{}}
	collection_names, err := iot_db.ListCollectionNames(db_ctx, filter)
	if err != nil {
		log.Fatal(err)
		code = http.StatusBadGateway
	} else {
		code = http.StatusOK
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": collection_names,
	})
}

func handler_get_temp(context *gin.Context) {
	temp_data_raw := read_device_sensor_data_collection()
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": temp_data_raw,
	})
}
