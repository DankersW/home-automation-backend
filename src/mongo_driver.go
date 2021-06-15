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

func read_device_sensor_data_collection() []SensorData {
	// todo: make me generic
	sensor_data_collection := iot_db.Collection("device_sensor_data")

	filter := generate_timestamp_filter(7, 0)
	data_cursor, err := sensor_data_collection.Find(db_ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	sensor_data := parse_sensor_data_document(data_cursor)
	data_cursor.Close(context.TODO())
	return sensor_data
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

func parse_sensor_data_document(cursor *mongo.Cursor) []SensorData {
	sensor_data := []SensorData{}
	for cursor.Next(db_ctx) {
		var document_item bson.M
		err := cursor.Decode(&document_item)
		if err != nil {
			log.Fatal(err)
		}

		var sensor_item SensorData
		sensor_item.Device_id = document_item["device_id"].(string)
		sensor_item.Timestamp = document_item["timestamp"].(primitive.DateTime)
		sensor_item.Temp = cast_to_float32(document_item["temperature"])
		sensor_item.Humi = cast_to_float32(document_item["humidity"])
		sensor_data = append(sensor_data, sensor_item)
	}
	return sensor_data
}
