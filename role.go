package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sashabaranov/go-openai"
	"gopkg.in/yaml.v3"
)

var RoleDir = filepath.Join(CfgDir, roleDir)

func CreateDefaultRole() error {
	_, err := GetRole(defaultRole)
	if err != nil {
		return CreateRole(defaultRole, "", defaultPrompt)
	}
	return nil
}

type Role struct {
	Name        string
	Description string
	Prompt      string
}

func (r *Role) ToMsg() openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: r.Prompt,
	}
}

func (r *Role) String() string {
	return fmt.Sprintf("name: %s\ndescription: %s\nprompt: %s\n", r.Name, r.Description, r.Prompt)
}

func GetRole(name string) (*Role, error) {
	path := filepath.Join(RoleDir, safeName(name))
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("role not exist")
		}
		return nil, err
	}
	var r Role
	err = yaml.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func CreateRole(name, desc, prompt string) error {
	path := filepath.Join(RoleDir, safeName(name))
	r := &Role{
		Name:        name,
		Description: desc,
		Prompt:      prompt,
	}
	b, err := yaml.Marshal(r)
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0644)
}

func ListRoles() []Role {
	entries, err := os.ReadDir(RoleDir)
	if err != nil {
		return nil
	}
	res := make([]Role, 0)
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		bytes, err := os.ReadFile(filepath.Join(RoleDir, e.Name()))
		if err != nil {
			continue
		}
		var r Role
		err = yaml.Unmarshal(bytes, &r)
		if err != nil {
			continue
		}
		res = append(res, r)
	}
	return res
}
