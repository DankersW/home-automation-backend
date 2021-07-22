package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var api_gateway *gin.Engine
var config Config

func init() {
	config = get_config()
	api_gateway = setup_api_gateway()
	connect_to_mongo()
}

func main() {
	api_port := fmt.Sprintf(":%d", config.Api.Port)
	api_gateway.Run(api_port)
}

func setup_api_gateway() *gin.Engine {
	api_engine := gin.Default()
	api_engine.Use(cors.Default())

	api_router_group := api_engine.Group("/api")
	setup_api_endpoints(api_router_group)
	return api_engine
}

func setup_api_endpoints(router_group *gin.RouterGroup) {
	router_group.GET("/", handler_get_all_apis)
	router_group.GET("/iotDbCollectionNames", handler_get_iot_db_collection_names)
	router_group.GET("/temp/stream", handler_get_temp_stream)
	router_group.GET("/temp/info", handler_get_temp_info)
	router_group.GET("/docker_info", handler_get_docker_info)
	router_group.GET("/devices/digital_twin", handler_get_devices_digital_twin)
	router_group.GET("/devices/status", handler_get_devices_status)
	router_group.GET("/host_health/info", handler_get_host_health_info)
	router_group.GET("/host_heath/stream", handler_get_host_health_stream)
}
