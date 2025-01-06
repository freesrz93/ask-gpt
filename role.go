package main

import (
	"errors"

	"github.com/sashabaranov/go-openai"
)

type Role struct {
	Description string
	Prompt      string
}

func (r *Role) ToMsg() openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: r.Prompt,
	}
}

func GetRole(name string) (*Role, error) {
	r, ok := Config.Roles[name]
	if !ok {
		return nil, errors.New("role not exist")
	}
	return r, nil
}

func CreateRole(name, desc, prompt string) error {
	_, ok := Config.Roles[name]
	if ok {
		return errors.New("role already exist")
	}
	Config.Roles[name] = &Role{
		Description: desc,
		Prompt:      prompt,
	}
	return SaveCfg()
}
