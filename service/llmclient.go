<<<<<<< HEAD
package service

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/ollama/ollama/api"
)

// Interface for all kind of llm connecting
type ConvoToLLM interface {
	Messaging(prompt string) (string, error)
}

type OllamaConnector struct {
	client *api.Client
}

func NewOllamaConnector() (*OllamaConnector, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
	url, err1 := url.Parse(os.Getenv("HOST"))
	if err1 != nil {
		log.Fatal("Error occured :", err1)
	}
	client := api.NewClient(url, nil)
	return &OllamaConnector{client: client}, nil
}

func (ol *OllamaConnector) Messaging(message string) (string, error) {

	req := &api.GenerateRequest{
		Model:  "", // Add Ollama model
		Prompt: message,
		Stream: new(bool),
	}

	ctx := context.Background()
	var fullResp strings.Builder

	respFunc := func(resp api.GenerateResponse) error {
		fmt.Println(resp.Response)
		fullResp.WriteString(resp.Response)
		return nil
	}

	err := ol.client.Generate(ctx, req, respFunc)
	if err != nil {
		log.Fatal("Response error occured: ", err)
	}

	return fullResp.String(), nil
}

// type OpenAIConnector struct {
// 	client *openai.Client
// }

// func NewOpenAIConnector(apiKey string) *OpenAIConnector {
// 	return &OpenAIConnector{
// 		client: openai.NewClient(option.WithAPIKey(apiKey)),
// 	}
// }

// func (o *OpenAIConnector) Messaging(message string) (string, error) {
// 	chatCompletion, err := o.client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
// 		Messages: openai.F(
// 			[]openai.ChatCompletionMessageParamUnion{
// 				openai.UserMessage(message),
// 			}),
// 		Model: openai.F(openai.ChatModelGPT4o),
// 	})
// 	if err != nil {
// 		return "", err
// 	}

// 	return chatCompletion.Choices[0].Message.Content, nil
// }

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
=======
package service

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/ollama/ollama/api"
)

// Interface for all kind of llm connecting
type ConvoToLLM interface {
	Messaging(prompt string) (string, error)
}

type OllamaConnector struct {
	client *api.Client
}

func NewOllamaConnector() (*OllamaConnector, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
	url, err1 := url.Parse(os.Getenv("HOST"))
	if err1 != nil {
		log.Fatal("Error occured :", err1)
	}
	client := api.NewClient(url, nil)
	return &OllamaConnector{client: client}, nil
}

func (ol *OllamaConnector) Messaging(message string) (string, error) {

	req := &api.GenerateRequest{
		Model:  "", // Add Ollama model
		Prompt: message,
		Stream: new(bool),
	}

	ctx := context.Background()
	var fullResp strings.Builder

	respFunc := func(resp api.GenerateResponse) error {
		fmt.Println(resp.Response)
		fullResp.WriteString(resp.Response)
		return nil
	}

	err := ol.client.Generate(ctx, req, respFunc)
	if err != nil {
		log.Fatal("Response error occured: ", err)
	}

	return fullResp.String(), nil
}

// type OpenAIConnector struct {
// 	client *openai.Client
// }

// func NewOpenAIConnector(apiKey string) *OpenAIConnector {
// 	return &OpenAIConnector{
// 		client: openai.NewClient(option.WithAPIKey(apiKey)),
// 	}
// }

// func (o *OpenAIConnector) Messaging(message string) (string, error) {
// 	chatCompletion, err := o.client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
// 		Messages: openai.F(
// 			[]openai.ChatCompletionMessageParamUnion{
// 				openai.UserMessage(message),
// 			}),
// 		Model: openai.F(openai.ChatModelGPT4o),
// 	})
// 	if err != nil {
// 		return "", err
// 	}

// 	return chatCompletion.Choices[0].Message.Content, nil
// }

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
>>>>>>> 6dd8c99975de2f870161f490b1e4d559d57c2652
