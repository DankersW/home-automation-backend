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

func get_devices_status() DeviceStatus {
	filter := bson.D{}
	cursor := mongo_read("digital_twin", filter)

	data := parse_dev_status_cursor(cursor)

	cursor.Close(context.TODO())
	return data
}

func parse_dev_status_cursor(cursor *mongo.Cursor) DeviceStatus {
	var total int = 0
	var active int = 0

	for cursor.Next(db_ctx) {
		var document_item bson.M
		err := cursor.Decode(&document_item)
		if err != nil {
			log.Fatal(err)
		}

		total++
		if document_item["active"].(bool) {
			active++
		}
	}

	var device_status DeviceStatus
	device_status.Total = total
	device_status.Online = active
	device_status.Offline = total - active

	return device_status
}

func get_temp_info() TempInfo {
	num_items := 5
	current_filter := bson.D{}
	day_filter := generate_timestamp_filter(1, 0)
	week_filter := generate_timestamp_filter(7, 0)
	current_cursor := mongo_read_x_items("device_sensor_data", current_filter, num_items)
	day_cursor := mongo_read("device_sensor_data", day_filter)
	week_cursor := mongo_read("device_sensor_data", week_filter)
	current_temp, current_humi := get_avarage_temp_humi_from_cursor(current_cursor)
	day_temp, day_humi := get_avarage_temp_humi_from_cursor(day_cursor)
	week_temp, week_humi := get_avarage_temp_humi_from_cursor(week_cursor)
	var temp_info TempInfo
	temp_info.Current.Temp = current_temp
	temp_info.Daily.Temp = day_temp
	temp_info.Weekly.Temp = week_temp
	temp_info.Current.Humi = current_humi
	temp_info.Daily.Humi = day_humi
	temp_info.Weekly.Humi = week_humi

	current_cursor.Close(context.TODO())
	day_cursor.Close(context.TODO())
	week_cursor.Close(context.TODO())
	return temp_info
}

func get_avarage_temp_humi_from_cursor(cursor *mongo.Cursor) (float32, float32) {
	var sum_temp float32 = 0
	var sum_humi float32 = 0
	var items int32 = 0

	for cursor.Next(db_ctx) {
		var document_item bson.M
		err := cursor.Decode(&document_item)
		if err != nil {
			log.Fatal(err)
		}
		items++
		sum_humi += cast_to_float32(document_item["humidity"])
		sum_temp += cast_to_float32(document_item["temperature"])
	}

	return sum_temp / float32(items), sum_humi / float32(items)
}

func get_host_info() HostHealthSummary {
	filter := generate_timestamp_filter(1, 0)
	cursor := mongo_read("host_health", filter)

	data := parse_host_health_summary(cursor)

	cursor.Close(context.TODO())
	return data
}

func parse_host_health_summary(cursor *mongo.Cursor) HostHealthSummary {
	var items int32 = 0
	var total_cpu float32 = 0
	var total_temp float32 = 0
	var host_health HostHealthSummary
	for cursor.Next(db_ctx) {
		var document_item bson.M
		err := cursor.Decode(&document_item)
		if err != nil {
			log.Fatal(err)
		}
		total_cpu += cast_to_float32(document_item["cpu_load"])
		total_temp += cast_to_float32(document_item["temperature"])
		if items == 0 {
			host_health.Current.Cpu = total_cpu
			host_health.Current.Temp = total_temp
		}
		items++
	}
	host_health.Daily.Temp = total_temp / float32(items)
	host_health.Daily.Cpu = total_cpu / float32(items)

	return host_health
}

func get_host_info_stream() []HostHealth {
	filter := generate_timestamp_filter(1, 0)
	cursor := mongo_read("host_health", filter)

	stream := parse_host_health_cursor_stream(cursor)

	cursor.Close(context.TODO())
	return stream
}

func parse_host_health_cursor_stream(cursor *mongo.Cursor) []HostHealth {
	data_stream := []HostHealth{}
	for cursor.Next(db_ctx) {
		var document_item bson.M
		err := cursor.Decode(&document_item)
		if err != nil {
			log.Fatal(err)
		}

		var item HostHealth
		item.Cpu = cast_to_float32(document_item["cpu_load"])
		item.Temp = cast_to_float32(document_item["temperature"])
		data_stream = append(data_stream, item)
	}
	return data_stream
}
