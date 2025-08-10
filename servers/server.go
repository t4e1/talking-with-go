package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/t4e1/talking-with-go.git/handlers"
	"github.com/t4e1/talking-with-go.git/services"
)

type Server struct {
	router    *gin.Engine
	promptSvc *services.PromptService
}

// SetupRouters makes new router and enrolls every handlers.
func (s *Server) SetupRouters() {
	routers := gin.Default()
	apiHandler := handlers.NewAPIHandler(s.promptSvc)

	// Setup route
	routers.POST("/msgs", apiHandler.Conversation)

	s.router = routers
}

// Starting server
func (s *Server) RunServer() {
	s.router.Run(":10001")
}
