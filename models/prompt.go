package models

type RequestMessage struct {
	ConvoType string `json:"type"`
	Language  string `json:"language"`
	Role      string `json:"role"`
	Message   string `json:"message"`
}

type ResponseMessage struct {
	Message string `json: "message"`
	// Should check LLM's response format
}
