package main

import "github.com/gin-gonic/gin"

func main() {
	api_engine := setup_routes()
	api_engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setup_router() *gin.Engine {
	api_engine := gin.Default()
	api_engine.GET("/ping", endpoint_handler_get_ping)
	api_engine.GET("/hello", endpoint_handler_get_hello)
	return api_engine
}

// Example structure at https://github.com/tidwall/gjson
