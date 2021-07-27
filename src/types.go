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

type DockerInfo struct {
	Name    string `json:"name"`
	Repo    string `json:"repo"`
	Version string `json:"version"`
	Uptime  string `json:"uptime"`
	Status  string `json:"status"`
	Port    string `json:"port"`
}

type DigitalTwin struct {
	Name       string `json:"device_name"`
	Active     bool   `json:"active"`
	Location   string `json:"location"`
	Technology string `json:"technology"`
	Battery    string `json:"battery_level"`
}

type DeviceStatus struct {
	Total   int `json:"total"`
	Online  int `json:"online"`
	Offline int `json:"offline"`
}

type Config struct {
	Mongo struct {
		Address string `yaml:"address"`
		Port    uint16 `yaml:"port"`
	}
	Api struct {
		Port uint16 `yaml:"port"`
	}
}

type TempHumi struct {
	Temp float32 `json:"temp"`
	Humi float32 `json:"humi"`
}

type TempInfo struct {
	Current TempHumi
	Daily   TempHumi
	Weekly  TempHumi
}

type HostHealth struct {
	Temp float32 `json:"temp"`
	Cpu  float32 `json:"cpu"`
}

type HostHealthSummary struct {
	Daily   HostHealth
	Current HostHealth
}

type ExternalWeather struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	Weather   string  `json:"weather"`
}
