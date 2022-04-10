package routes

import (
	"net/http"

	"github.com/ishank838/data-collection/api/handler"
)

type Route struct {
	Name    string
	Path    string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
	Auth    bool
}

func GetAllRoutes() *[]Route {
	allRoutes := []Route{
		{
			Name:    "Insert",
			Method:  http.MethodPost,
			Path:    "/v1/insert",
			Handler: handler.Insert,
			Auth:    false,
		},
		{
			Name:    "Query",
			Method:  http.MethodPost,
			Path:    "/v1/query",
			Handler: handler.Query,
			Auth:    false,
		},
	}

	return &allRoutes
}
