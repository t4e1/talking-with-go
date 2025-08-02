package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type request struct {
	Prompt string
}

func main() {

	// router
	router := gin.Default()

	// A client communicate with LLM at this end point
	router.POST("/queries", sendToLLM)

	// Run server as localhost:8080
	router.Run(":8080")
}

func sendToLLM(c *gin.Context) {

	var req request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error ": "Invalid JSON"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Hello",
	})
	fmt.Println("Communicate Success : ", req.Prompt)
}
