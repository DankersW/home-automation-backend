package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SensorData struct {
	Device_id string             `json:"device_id"`
	Timestamp primitive.DateTime `json:"timestamp"`
	Temp      float32            `json:"temp"`
	Humi      float32            `json:"humi"`
}
