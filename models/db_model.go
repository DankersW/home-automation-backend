package models

import (
	"github.com/dankersw/home-automation-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SensorData struct {
	Device_id string             `json:"device_id"`
	Timestamp primitive.DateTime `json:"timestamp"`
	Temp      float32            `json:"temp"`
	Humi      float32            `json:"humi"`
}

func ToSensorData(data bson.M) SensorData {
	sd := SensorData{
		Device_id: utils.ToString(data["device_id"]),
		Temp:      utils.ToFloat32(data["temperature"]),
		Humi:      utils.ToFloat32(data["humidity"]),
		Timestamp: data["timestamp"].(primitive.DateTime),
	}
	return sd
}
