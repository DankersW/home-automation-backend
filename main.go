package main

import "github.com/gin-gonic/gin"

func main() {
	api_engine := setup_api_gateway()
	api_engine.Run(":8090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
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

// Example structure at https://github.com/tidwall/gjson
