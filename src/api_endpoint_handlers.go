package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func handler_get_all_apis(context *gin.Context) {
	apis := map[string]string{
		"all":                       "/api/",
		"get all collection names":  "/api/iotDbCollectionNames",
		"temperature":               "/api/temp",
		"docker info":               "/api/docker_info",
		"connected devices info":    "/api/devices/status_info",
		"connected devices summary": "/api/devices/status_summary",
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

func handler_get_docker_info(context *gin.Context) {
	docker_info := get_docker_info()
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": docker_info,
	})
}

func handler_get_devices_status_detail(context *gin.Context) {
	device_info := "device info"
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": device_info,
	})
}

func handler_get_devices_status_summary(context *gin.Context) {
	device_info := "device summary"
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": device_info,
	})
}
