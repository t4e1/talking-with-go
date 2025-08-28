package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/t4e1/talking-with-go.git/models"
	"github.com/t4e1/talking-with-go.git/service"
)

// APIHandler handles all API endpoints
type APIHandler struct {
	promptSvc *service.PromptService
}

func NewAPIHandler(promptSvc *service.PromptService) *APIHandler {
	return &APIHandler{
		promptSvc: promptSvc,
	}
}

func (h *APIHandler) Conversation(ctx *gin.Context) {
	req := &models.RequestMessage{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"Error is occured: ": err})
	}

	// build message
	h.promptSvc.BuildPromptMessage(req)

	// connect to llm (not yet)

	// build response message
	resp := h.promptSvc.BuildResponseMessage()

	// create a response to client
	ctx.JSON(200, "Response Message: "+resp.Message)
}
