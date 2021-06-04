package main

import (
	"context"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func trim_first_char(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func get_docker_info() []DockerInfo {
	ctx := context.Background()
	docker_cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	data_containers, err := docker_cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	docker_info := []DockerInfo{}
	for _, container := range data_containers {
		var info_item DockerInfo
		info_item.Name = trim_first_char(container.Names[0])
		image := strings.Split(container.Image, ":")
		info_item.Repo = image[0]
		info_item.Version = image[1]
		info_item.Status = container.State
		info_item.Uptime = container.Status

		var ports string = ""
		for _, port := range container.Ports {
			ports += strconv.Itoa(int(port.PublicPort)) + " "
		}
		info_item.Port = ports

		docker_info = append(docker_info, info_item)
	}
	return docker_info
}