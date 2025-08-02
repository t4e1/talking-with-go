package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var request struct { 
	
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

	fmt.Println(c.))
	c.JSON(200, gin.H{
		"message": "Hello",
	})
	fmt.Println("Communicate Success")
}
