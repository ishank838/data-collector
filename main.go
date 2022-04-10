package main

import (
	"log"

	"github.com/ishank838/data-collection/api/config"
	"github.com/ishank838/data-collection/api/webserver"
)

func main() {

	config, err := config.InitConfig()

	if err != nil {
		log.Fatal(err)
	}

	webServer := webserver.NewWebserver(config)

	err = webServer.StartServer()

	if err != nil {
		log.Fatal(err)
	}
}
