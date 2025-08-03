package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type convoType string

const (
	Interview  convoType = "interview"
	DailyConvo convoType = "daily conversation"
)

type convoBasicInfo struct {
	// Conversation types selected by Client (ex.Interview, Daily Conversation, Business...)
	ConvoType convoType `json: "convo_type" binding: "required"`
	// some information (Not planned yet)
}

type request struct {
	Prompt string
}

func getConvoInfo(c *gin.Context) {
	var info convoBasicInfo
	fmt.Println("Success to get some informations: ", info)
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
	fmt.Println("Communicate Success: ", req)
}

func main() {
	router := gin.Default()

	// Endpoint for getting information about conversation
	router.POST("/convos", getConvoInfo)

	// Endpoint for communication with LLM.
	router.POST("/queries", sendToLLM)

	// Run server as localhost:8080
	router.Run(":8080")
}
