package main

import (
	"context"
	"io"
	"strings"

	"github.com/sashabaranov/go-openai"
)

var (
	Client *C
)

func InitClient() {
	cfg := openai.DefaultConfig(Config.APIKey)
	cfg.BaseURL = Config.BaseURL
	Client = &C{Client: openai.NewClientWithConfig(cfg)}
}

type C struct {
	*openai.Client
}

func (c *C) Stream(s *Session, input string) error {
	defer Pln()

	s.Append(openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: input,
	})

	// https://platform.openai.com/docs/api-reference/chat
	stream, err := c.CreateChatCompletionStream(context.Background(), openai.ChatCompletionRequest{
		Messages:         s.Get(),
		Model:            Config.Model,
		MaxTokens:        Config.MaxTokens,
		Temperature:      Config.Temperature,
		TopP:             Config.TopP,
		Stream:           true,
		PresencePenalty:  Config.PresencePenalty,
		FrequencyPenalty: Config.FrequencyPenalty,
	})
	if err != nil {
		return err
	}
	defer stream.Close()

	builder := strings.Builder{}
	for {
		r, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if len(r.Choices) == 0 {
			continue
		}
		choice := r.Choices[0]
		if str := choice.Delta.Content; str != "" {
			P(str)
			builder.WriteString(str)
		}
	}

	s.Append(openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: builder.String(),
	})
	return s.Save()
}
