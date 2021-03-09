package main

import "github.com/gin-gonic/gin"

var api_gateway *gin.Engine

func init() {
	api_gateway = setup_api_gateway()
}

func main() {
	api_gateway.Run(":8090")
}

func setup_api_gateway() *gin.Engine {
	api_engine := gin.Default()
	api_router_group := api_engine.Group("/api")
	setup_api_endpoints(api_router_group)
	return api_engine
}

func setup_api_endpoints(router_group *gin.RouterGroup) {
	router_group.GET("/", handler_get_all_apis)
	router_group.GET("/random", handler_get_random)
}
