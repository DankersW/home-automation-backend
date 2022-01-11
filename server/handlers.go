package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type apiRecources struct {
	activeRoutes func() RestRoutes
}

func NewHandlers(activeRoutes func() RestRoutes) *apiRecources {
	apiRecources := &apiRecources{
		activeRoutes: activeRoutes,
	}
	return apiRecources
}

func (a *apiRecources) getRestRoutes() RestRoutes {
	route := func(method string, uri string, callback GinCallback) RestRoute {
		return RestRoute{method: method, uri: uri, callback: callback}
	}
	routes := RestRoutes{
		route(http.MethodGet, "/", a.allRoutes),
		route(http.MethodGet, "/hello", helloWorld),
	}
	return routes
}

func (a *apiRecources) allRoutes(context *gin.Context) {
	log.Infof("All active routes: %v", a.activeRoutes())
}

func helloWorld(context *gin.Context) {
	log.Info("world")
}
