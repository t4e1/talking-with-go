package servers

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
}

func NewServer() {
	router := gin.Default()
	router.Run(":10001")
}
