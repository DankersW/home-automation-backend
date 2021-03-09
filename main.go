package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"
)

var api_gateway *gin.Engine
var iot_db *mongo.Database
var db_ctx = context.TODO()

func init() {
	api_gateway = setup_api_gateway()
	connect_to_mongo()
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
	router_group.GET("/iotDbCollectionNames", handler_get_iot_db_collection_names)
}

func connect_to_mongo() {
	clientOptions := options.Client().ApplyURI("mongodb://admin:mongo_admin_iot@192.168.1.140:27017/")
	client, err := mongo.Connect(db_ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(db_ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	iot_db = client.Database("iot_db")
}
