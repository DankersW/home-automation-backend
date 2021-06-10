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

type Config struct {
	Mongo struct {
		Address string `yaml:"address"`
		Port    uint16 `yaml:"port"`
	}
	Api struct {
		Port uint16 `yaml:"port"`
	}
}
