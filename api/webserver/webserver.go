package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ishank838/data-collection/api/config"
	"github.com/ishank838/data-collection/api/webserver/router"
	"github.com/ishank838/data-collection/api/webserver/routes"
)

type WebServer struct {
	Router  *mux.Router
	Routes  *[]routes.Route
	Address string
}

func NewWebserver(config *config.Config) *WebServer {
	routes := routes.GetAllRoutes()
	router := router.NewRouter(config, routes)
	address := fmt.Sprintf(":%v", config.WebPort)
	return &WebServer{
		Router:  router,
		Routes:  routes,
		Address: address,
	}
}

func (w *WebServer) StartServer() error {

	log.Println("starting server at port: ", w.Address)

	err := http.ListenAndServe(w.Address, w.Router)

	if err != nil {
		return fmt.Errorf("error at start server: %v", err)
	}

	return nil
}
