package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (t *ttess) c_(context *gin.Context) {
	log.Warnf("oke %s", t.t)
	log.Info(context)
}

type ttess struct {
	t string
}

func getRestRoutes() RestRoutes {

	tt := ttess{t: "cool"}

	route := func(method string, uri string, callback GinCallback) RestRoute {
		return RestRoute{method: method, uri: uri, callback: callback}
	}
	routes := RestRoutes{
		route(http.MethodGet, "/hi", c),
		route(http.MethodGet, "/hi_you", tt.c_),
	}
	return routes
}

func c(context *gin.Context) {
	log.Warn("oke")
	log.Info(context)
}
