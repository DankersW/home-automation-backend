package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var iot_db *mongo.Database
var db_ctx = context.TODO()

func connect_to_mongo() {
	mongo_url := fmt.Sprintf("mongodb://admin:mongo_admin_iot@%s:%d/", config.Mongo.Address, config.Mongo.Port)
	clientOptions := options.Client().ApplyURI(mongo_url)

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

func mongo_read(collection string, filter primitive.D) *mongo.Cursor {
	sensor_data_collection := iot_db.Collection(collection)
	data_cursor, err := sensor_data_collection.Find(db_ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return data_cursor
}

func generate_timestamp_filter(oldest_day_limit int, max_day_limit int) primitive.D {
	filter := bson.D{
		primitive.E{
			Key: "timestamp", Value: bson.D{primitive.E{
				Key: "$gte", Value: primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -oldest_day_limit))}}},
		primitive.E{
			Key: "timestamp", Value: bson.D{primitive.E{
				Key: "$lte", Value: primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -max_day_limit))}}},
	}
	return filter
}
