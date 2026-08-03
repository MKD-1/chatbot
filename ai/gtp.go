package gtp

import (
	"context"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func TalkText(apiKey string, baseUrl string, model string, systemPrompt string, content string) (string, error) {
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
		option.WithBaseURL(baseUrl),
	)
	completion, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Model: model,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemPrompt),
			openai.UserMessage(content),
		},
	})
	if err != nil {
		return "", err
	}
	return completion.Choices[0].Message.Content, nil
}
