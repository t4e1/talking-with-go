package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/t4e1/talking-with-go.git/services"
)

type Server struct {
	getConvoTypeService *services.GetConvoTypeService
	tryConvoService     *services.TryConvoService
	createPromptService *services.CreatePromptService
}

func NewServer() {
	router := gin.Default()
	router.Run(":10001")
}
