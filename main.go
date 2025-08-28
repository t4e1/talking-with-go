package main

import (
	"log"

	"github.com/t4e1/talking-with-go.git/server"
	"github.com/t4e1/talking-with-go.git/service"
)

func main() {
	// Initialize Services
	service, err := service.InitServices()
	if err != nil || service == nil {
		log.Fatalf("Failed to initialize services: %v", err)
	}

	// Setup Routers
	srv := &server.Server{}
	srv.SetupRouters()

	// Start server
	log.Printf("Server is starting on http://localhost:%d", 10001)
	srv.RunServer()
}
