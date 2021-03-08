package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler_get_all_apis(context *gin.Context) {
	apis := map[string]string{
		"all":    "/api/",
		"random": "/api/random",
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": apis,
	})
}

func handler_get_random(context *gin.Context) {
	random_number := rand.Intn(50)
	context.JSON(200, gin.H{
		"Random": random_number,
	})
}
