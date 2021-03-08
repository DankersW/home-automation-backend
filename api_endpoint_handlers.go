package main

import "github.com/gin-gonic/gin"

// https://github.com/gin-gonic/examples/blob/master/group-routes/routes/users.go

func handler_get_all_apis(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}

func handler_get_random(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Hello world",
	})
}
