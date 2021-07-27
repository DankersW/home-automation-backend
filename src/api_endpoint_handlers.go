package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func handler_get_all_apis(context *gin.Context) {
	apis := map[string]string{
		"all":                      "/api/",
		"get all collection names": "/api/iotDbCollectionNames",
		"get one week of temperature stream data":                      "/api/temp/stream",
		"get info about the temperature and humidity":                  "/api/temp/info",
		"predicted outdoor temp":                                       "/api/temp/predicted",
		"docker info":                                                  "/api/docker_info",
		"digital twin info":                                            "/api/devices/digital_twin",
		"connected devices status summary":                             "/api/devices/status",
		"summary of information about the host running the docker app": "/api/host_health/info",
		"stream of host health information over the last day":          "/api/host_health/stream",
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

func handler_get_temp_stream(context *gin.Context) {
	temp_data_raw := get_device_sensor_data()
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": temp_data_raw,
	})
}

func handler_get_temp_info(context *gin.Context) {
	data := get_temp_info()
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": data,
	})
}
func handler_get_outdoor_temp_prediction(context *gin.Context) {
	data := get_external_weather()
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": data,
	})
}

func handler_get_docker_info(context *gin.Context) {
	docker_info := get_docker_info()
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": docker_info,
	})
}

func handler_get_devices_digital_twin(context *gin.Context) {
	device_info := get_digital_twin()
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": device_info,
	})
}

func handler_get_devices_status(context *gin.Context) {
	device_info := get_devices_status()
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": device_info,
	})
}

func handler_get_host_health_info(context *gin.Context) {
	device_info := get_host_info()
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": device_info,
	})
}

func handler_get_host_health_stream(context *gin.Context) {
	device_info := get_host_info_stream()
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": device_info,
	})
}
