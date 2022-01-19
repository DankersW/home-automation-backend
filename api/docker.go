package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *api) GetDockerInfo(gc *gin.Context) {
	Reply(gc, http.StatusOK, "all good", nil)
}
