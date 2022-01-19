package main

import (
	"context"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"

	//"github.com/dankersw/home-automation-backend/server"

	"github.com/dankersw/home-automation-backend/config"
	"github.com/dankersw/home-automation-backend/server"
)

/*
func setup_api_endpoints(router_group *gin.RouterGroup) {
	router_group.GET("/", handler_get_all_apis)                                    // done
	router_group.GET("/iotDbCollectionNames", handler_get_iot_db_collection_names) // done
	router_group.GET("/temp/stream", handler_get_temp_stream)
	router_group.GET("/temp/info", handler_get_temp_info)
	router_group.GET("/temp/predicted", handler_get_outdoor_temp_prediction)
	router_group.GET("/docker_info", handler_get_docker_info)
	router_group.GET("/devices/digital_twin", handler_get_devices_digital_twin)
	router_group.GET("/devices/status", handler_get_devices_status)
	router_group.GET("/host_health/info", handler_get_host_health_info)
	router_group.GET("/host_heath/stream", handler_get_host_health_stream)
}
*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := config.Get()

	log.Info(config)

	server, err := server.New(ctx, config)
	if err != nil {
		log.Error("Failed to create server")
	}
	server.Start()

	quit := make(chan os.Signal, 10)
	signal.Notify(quit, os.Interrupt)
	<-quit

	server.Close()
}
