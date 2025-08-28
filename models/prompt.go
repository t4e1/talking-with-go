<<<<<<< HEAD
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
=======
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
>>>>>>> 6dd8c99975de2f870161f490b1e4d559d57c2652
