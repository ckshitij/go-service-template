package main

import (
	"log"

	"github.io/ckshitij/go-service-template/config"
	"github.io/ckshitij/go-service-template/server"
)

func main() {
	// Load configuration
	conf, err := config.LoadConfig("./resources/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Start the Gin server
	log.Printf("Starting server on %s:%d", conf.Server.Host, conf.Server.Port)
	if err := server.StartServer(conf); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
