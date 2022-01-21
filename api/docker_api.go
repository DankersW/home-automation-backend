package api

import (
	"context"
	"net/http"
	"time"

	"github.com/dankersw/home-automation-backend/models"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

func (a *api) GetDockerInfo(gc *gin.Context) {
	dockerData, err := getDockerInfo()
	if err != nil {
		fail(gc, "Failed to docker info", err)
	}
	Reply(gc, http.StatusOK, dockerData, nil)
}

func getDockerInfo() ([]models.DockerContainerInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*1))
	defer cancel()

	dockerCli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	containers, err := dockerCli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	dockerData := []models.DockerContainerInfo{}
	for _, container := range containers {
		item := models.ToDockerContainerInfo(container)
		dockerData = append(dockerData, item)
	}

	return dockerData, nil
}
