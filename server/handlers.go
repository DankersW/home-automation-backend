package server

import (
	"net/http"

	"github.com/dankersw/home-automation-backend/api"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type handler struct {
	activeRoutes func() RestRoutes
	api          api.Api
}

func NewHandlers(activeRoutes func() RestRoutes, endpoints api.Api) *handler {
	h := &handler{
		activeRoutes: activeRoutes,
		api:          endpoints,
	}
	return h
}

func (h *handler) getRestRoutes() RestRoutes {
	route := func(method string, uri string, callback ginCallback) RestRoute {
		return RestRoute{method: method, uri: uri, callback: callback}
	}
	routes := RestRoutes{
		route(http.MethodGet, "/", h.allRoutes),
		route(http.MethodGet, "/hello", helloWorld),
		route(http.MethodGet, "/db", h.api.DbCall),
	}
	return routes
}

func (h *handler) allRoutes(context *gin.Context) {
	log.Infof("All active routes: %v", h.activeRoutes())
}

func helloWorld(context *gin.Context) {
	log.Info("world")
}
