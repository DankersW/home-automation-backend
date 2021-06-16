package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func get_device_sensor_data() []SensorData {
	filter := generate_timestamp_filter(7, 0)
	cursor := mongo_read("device_sensor_data", filter)

	data := parse_sensor_data_cursor(cursor)

	cursor.Close(context.TODO())
	return data
}

func parse_sensor_data_cursor(cursor *mongo.Cursor) []SensorData {
	sensor_data := []SensorData{}
	for cursor.Next(db_ctx) {
		var document_item bson.M
		err := cursor.Decode(&document_item)
		if err != nil {
			log.Fatal(err)
		}

		var sensor_item SensorData
		sensor_item.Device_id = cast_to_string(document_item["device_id"])
		sensor_item.Timestamp = document_item["timestamp"].(primitive.DateTime)
		sensor_item.Temp = cast_to_float32(document_item["temperature"])
		sensor_item.Humi = cast_to_float32(document_item["humidity"])
		sensor_data = append(sensor_data, sensor_item)
	}
	return sensor_data
}

func get_digital_twin() []DigitalTwin {
	filter := bson.D{}
	cursor := mongo_read("digital_twin", filter)

	data := parse_digital_twin_cursor(cursor)

	cursor.Close(context.TODO())
	return data
}

func parse_digital_twin_cursor(cursor *mongo.Cursor) []DigitalTwin {
	digital_twin := []DigitalTwin{}

	for cursor.Next(db_ctx) {
		var document_item bson.M
		err := cursor.Decode(&document_item)
		if err != nil {
			log.Fatal(err)
		}

		var item DigitalTwin
		item.Name = document_item["device_name"].(string)
		item.Active = document_item["active"].(bool)
		item.Location = cast_to_string(document_item["location"])
		item.Technology = cast_to_string(document_item["technology"])
		item.Battery = cast_to_string(document_item["battery_level"])
		digital_twin = append(digital_twin, item)
	}
	return digital_twin
}

func get_devices_status() []DeviceStatus {
	filter := bson.D{}
	cursor := mongo_read("digital_twin", filter)

	data := parse_dev_status_cursor(cursor)

	cursor.Close(context.TODO())
	return data
}

func parse_dev_status_cursor(cursor *mongo.Cursor) []DeviceStatus {
	status := []DeviceStatus{}

	for cursor.Next(db_ctx) {
		var document_item bson.M
		err := cursor.Decode(&document_item)
		if err != nil {
			log.Fatal(err)
		}

		var item DeviceStatus
		//item.Name = document_item["device_name"].(string)
		//item.Active = document_item["active"].(bool)
		//item.Location = cast_to_string(document_item["location"])
		//item.Technology = cast_to_string(document_item["technology"])
		//item.Battery = cast_to_string(document_item["battery_level"])
		status = append(status, item)
	}
	return status
}
