package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	webPort = "8081"
)

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting Broker Server on port %s\n", webPort)

	// Define Server
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// Start the Server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic("Error occured ", err)
	}
}
