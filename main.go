<<<<<<< HEAD
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
=======
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
>>>>>>> 6dd8c99975de2f870161f490b1e4d559d57c2652
