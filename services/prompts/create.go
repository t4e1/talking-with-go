package prompts

type PromptCreateService struct{}

func NewPromptCreateService() *PromptCreateService {
	return &PromptCreateService{}
}

func (pc *PromptCreateService) BuildPromptMessage(convoType, role, message, language string) []string {

	return []string{
		"Using language: ",
		language,
		"\nYou are participating in a conversation of the following type:",
		convoType,
		"\nYour role is: ",
		role,
		"\nThe client's message is: ",
		message,
		"\nInstructions: ",
		"1. Respond appropriately to the client's message.",
		"2. Summarize the current conversation at the end.",
		"",
		"Return the response using the following JSON format:",
		"[response]: your answer",
		"[summary]: summary of the entire conversation",
	}
}
