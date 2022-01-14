package server

import (
	"net/http"

	"github.com/dankersw/home-automation-backend/api"
	"github.com/dankersw/home-automation-backend/models"
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
		route(http.MethodGet, "/iotDbCollectionNames", r.api.GetDbCollectionNames),
		route(http.MethodGet, "/sensor/data", r.api.GetSensorData),
	}
	return routes
}

func (r *routeHandler) allRoutes(context *gin.Context) {
	x := models.Routes{V: "hee"}
	log.Info(x)
	log.Infof("All active routes: %v", r.activeRoutes())
}
