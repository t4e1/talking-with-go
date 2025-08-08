package main

import (
	"log"

	"github.com/t4e1/talking-with-go.git/servers"
	"github.com/t4e1/talking-with-go.git/services"
)

func main() {
	// Initialize Services
	services, err := services.InitServices()
	if err != nil || services == nil {
		log.Fatalf("Failed to initialize services: %v", err)
	}

	// Setup Routers
	srv := &servers.Server{}
	srv.SetupRouters()

	// Start server
	log.Printf("Server is starting on http://localhost:%d", 10001)
	srv.RunServer()
}
