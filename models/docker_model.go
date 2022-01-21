package models

import (
	"strconv"
	"strings"

	"github.com/dankersw/home-automation-backend/utils"
	"github.com/docker/docker/api/types"
)

type DockerContainerInfo struct {
	Name    string `json:"name"`
	Repo    string `json:"repo"`
	Version string `json:"version"`
	Uptime  string `json:"uptime"`
	Status  string `json:"status"`
	Port    string `json:"port"`
}

func ToDockerContainerInfo(container types.Container) DockerContainerInfo {
	image := strings.Split(container.Image, ":")
	dockerInfo := DockerContainerInfo{
		Name:    utils.RmFirstChar(container.Names[0]),
		Repo:    image[0],
		Version: image[1],
		Status:  container.State,
		Uptime:  container.Status,
		Port:    parseDockerPorts(container.Ports),
	}
	return dockerInfo
}

func parseDockerPorts(containerPorts []types.Port) string {
	ports := []string{}
	for _, containerPort := range containerPorts {
		port := strconv.Itoa(int(containerPort.PublicPort))
		if !utils.StrInSlice(port, ports) {
			ports = append(ports, port)
		}
	}
	return strings.Join(ports[:], ", ")
}
