package models

import (
	"github.com/dankersw/home-automation-backend/utils"
	"github.com/docker/docker/api/types"
)

type DockerInfo struct {
	Name    string `json:"name"`
	Repo    string `json:"repo"`
	Version string `json:"version"`
	Uptime  string `json:"uptime"`
	Status  string `json:"status"`
	Port    string `json:"port"`
}

func ToDockerInfo(container types.Container) DockerInfo {
	dockerInfo := DockerInfo{
		Name: utils.RmFirstChar(container.Names[0]),
	}
	return dockerInfo
}

func parseDockerPorts(ports []types.Port) string {
	return ""
}
