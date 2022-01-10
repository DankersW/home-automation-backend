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
		info_item.Port = parse_ports(container.Ports)

		docker_info = append(docker_info, info_item)
	}
	return docker_info
}

func parse_ports(ports_struct []types.Port) string {
	ports_array := []string{}
	for _, port := range ports_struct {
		port := strconv.Itoa(int(port.PublicPort))
		if !string_in_slice(port, ports_array) {
			ports_array = append(ports_array, port)
		}
	}
	return strings.Join(ports_array[:], ", ")
}

func string_in_slice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
