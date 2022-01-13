package server

import (
	"net/http"

	"github.com/dankersw/home-automation-backend/api"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type routeHandler struct {
	activeRoutes func() RestRoutes
	api          api.Api
}

func NewRouteHandler(activeRoutes func() RestRoutes, endpoints api.Api) *routeHandler {
	h := &routeHandler{
		activeRoutes: activeRoutes,
		api:          endpoints,
	}
	return h
}

func (r *routeHandler) getRestRoutes() RestRoutes {
	route := func(method string, uri string, callback ginCallback) RestRoute {
		return RestRoute{method: method, uri: uri, callback: callback}
	}
	routes := RestRoutes{
		route(http.MethodGet, "/", r.allRoutes),
		route(http.MethodGet, "/hello", helloWorld),
		route(http.MethodGet, "/iotDbCollectionNames", r.api.GetDbCollectionNames),
	}
	return routes
}

func (r *routeHandler) allRoutes(context *gin.Context) {
	log.Infof("All active routes: %v", r.activeRoutes())
}

func helloWorld(context *gin.Context) {
	log.Info("world")
}
