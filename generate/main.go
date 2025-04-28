package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go"
)

// MODEL_RUNNER_BASE_URL=http://localhost:12434 go run main.go
func main() {
	// Docker Model Runner Chat base URL
	llmURL := os.Getenv("MODEL_RUNNER_BASE_URL") + "/engines/llama.cpp/v1/"
	model := os.Getenv("LLM")

	client := openai.NewClient(
		option.WithBaseURL(llmURL),
		option.WithAPIKey(""),
	)

	ctx := context.Background()

	// content = first argument
	filePath := os.Args[1:][0]
	// sourceCode = content of filePath
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalln("😡:", err)
	}
	sourceCode := string(file)

	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage("You are a helpful assistant, expert in Golang Programming."),
		openai.UserMessage("Generate unit tests for the following source code:\n" + sourceCode),
	}

	param := openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    model,
		Temperature: openai.Opt(0.8),
	}

	completion, err := client.Chat.Completions.New(ctx, param)

	if err != nil {
		log.Fatalln("😡:", err)
	}
	fmt.Println(completion.Choices[0].Message.Content)

}
