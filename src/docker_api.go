package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

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

		// todo: fix port entry
		fmt.Println(container.Ports)

		var info_item DockerInfo
		info_item.Name = container.Names[0]
		info_item.Version = container.Image
		info_item.Status = container.State
		info_item.Uptime = container.Status

		docker_info = append(docker_info, info_item)
	}
	fmt.Println(docker_info)
	return docker_info
}
