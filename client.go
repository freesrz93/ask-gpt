package main

import (
	"context"
	"io"
	"strings"

	"github.com/sashabaranov/go-openai"
)

func NewClient(opt *BackendOption) *Client {
	cfg := openai.DefaultConfig(opt.APIKey)
	cfg.BaseURL = opt.BaseURL
	return &Client{BackendOption: opt, Client: openai.NewClientWithConfig(cfg)}
}

type Client struct {
	*BackendOption
	*openai.Client
}

func (c *Client) Stream(s *Session, input string) error {
	defer Pln()

	s.Append(openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: input,
	})

	// https://platform.openai.com/docs/api-reference/chat
	stream, err := c.CreateChatCompletionStream(context.Background(), openai.ChatCompletionRequest{
		Messages:         s.Get(),
		Model:            c.Model,
		MaxTokens:        c.MaxTokens,
		Temperature:      c.Temperature,
		TopP:             c.TopP,
		Stream:           true,
		PresencePenalty:  c.PresencePenalty,
		FrequencyPenalty: c.FrequencyPenalty,
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
