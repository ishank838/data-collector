package router

import (
	"github.com/gorilla/mux"
	"github.com/ishank838/data-collection/api/config"
	"github.com/ishank838/data-collection/api/webserver/routes"
)

func NewRouter(config *config.Config, routes *[]routes.Route) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	setAllRoutes(router, *routes)

	return router
}

func setAllRoutes(router *mux.Router, routes []routes.Route) {

	for _, v := range routes {
		router.Name(v.Name).Methods(v.Method).Path(v.Path).HandlerFunc(v.Handler)
	}
}
