package services

type Services struct {
	Prompt *PromptService
}

// InitServices initialize all services at server startup
func InitServices() (*Services, error) {
	promptSvc := NewPromptService()

	return &Services{
		Prompt: promptSvc,
	}, nil
}
