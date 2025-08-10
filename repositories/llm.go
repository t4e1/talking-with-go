package repositories

import (
	"context"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// Interface for all kind of llm connecting
type ConvoToLLM interface {
	Messaging(prompt string) (string, error)
}

type OpenAIConnector struct {
	client *openai.Client
}

func NewOpenAIConnector(apiKey string) *OpenAIConnector {
	return &OpenAIConnector{
		client: openai.NewClient(option.WithAPIKey(apiKey)),
	}
}

func (o *OpenAIConnector) Messaging(message string) (string, error) {
	chatCompletion, err := o.client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F(
			[]openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(message),
			}),
		Model: openai.F(openai.ChatModelGPT4o),
	})
	if err != nil {
		return "", err
	}

	return chatCompletion.Choices[0].Message.Content, nil
}

// To using API, I need to pay... Not nessacery now

// func OpenAIConnector() {
// 	client := openai.NewClient(
//     option.WithAPIKey("My API Key"), // defaults to os.LookupEnv("OPENAI_API_KEY")
// 	)
// 	chatCompletion, err := client.Chat.Completions.New(
//     context.TODO(), openai.ChatCompletionNewParams{
// 		Messages: openai.F(
//         []openai.ChatCompletionMessageParamUnion{
// 			openai.UserMessage("Say this is a test"),
//         }),
// 		Model: openai.F(openai.ChatModelGPT4o),
// 	}
// )
// }
