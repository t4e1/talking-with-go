package service

import (
	"strings"

	"github.com/t4e1/talking-with-go.git/models"
)

type PromptService struct {
}

// NewPromptService creates a new PromptService at server startup
func NewPromptService() *PromptService {
	return &PromptService{}
}

// BuildPromptMessage returns a prompt which has basic information and message for conversation.
func (pc *PromptService) BuildPromptMessage(infos *models.RequestMessage) string {

	temp := []string{"Using language: ",
		infos.Language,
		"\nYou are participating in a conversation of the following type:",
		infos.ConvoType,
		"\nYour role is: ",
		infos.Role,
		"\nThe client's message is: ",
		infos.Message,
		"\nInstructions: ",
		"1. Respond appropriately to the client's message.",
		"2. Summarize the current conversation at the end.",
		"",
		"Return the response using the following JSON format:",
		"[response]: your answer",
		"[summary]: summary of the entire conversation"}

	result := strings.Join(temp, "")
	return result
}

func (pc *PromptService) Conversation(prompt string) string {

	// call llm connector

	return "test string"
}

func (pc *PromptService) BuildResponseMessage() *models.ResponseMessage {

	return &models.ResponseMessage{
		Message: "Test Response Message",
	}
}
