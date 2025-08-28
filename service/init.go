<<<<<<< HEAD
package service

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
=======
package service

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
>>>>>>> 6dd8c99975de2f870161f490b1e4d559d57c2652
