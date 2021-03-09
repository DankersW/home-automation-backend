package main

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func handler_get_all_apis(context *gin.Context) {
	apis := map[string]string{
		"all":                     "/api/",
		"random":                  "/api/random",
		"iot_db_collection_names": "/api/iotDbCollectionNames",
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": apis,
	})
}

func handler_get_random(context *gin.Context) {
	random_number := rand.Intn(50)
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": random_number,
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
