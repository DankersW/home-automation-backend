package api_routes

import "github.com/gin-gonic/gin"

// https://github.com/gin-gonic/examples/blob/master/group-routes/routes/users.go

func setup_routes(router_group *gin.RouterGroup) {
	rg := router_group.Group("/api")
	rg.GET("/ping", endpoint_handler_get_ping)
	rg.GET("/hello", endpoint_handler_get_hello)

}

func endpoint_handler_get_ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}

func endpoint_handler_get_hello(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Hello world",
	})
}
