package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/sashabaranov/go-openai"
)

var SessionDir = filepath.Join(CfgDir, sessionDir)

func ListSessions() []*Session {
	entries, err := os.ReadDir(SessionDir)
	if err != nil {
		return nil
	}
	res := make([]*Session, 0)
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		bytes, err := os.ReadFile(filepath.Join(SessionDir, e.Name()))
		if err != nil {
			continue
		}
		var r Session
		err = json.Unmarshal(bytes, &r)
		if err != nil {
			continue
		}
		res = append(res, &r)
	}
	return res
}

func GetSession(id string) (*Session, error) {
	if id == tempSession {
		return &Session{
			tmp:      true,
			Messages: make([]openai.ChatCompletionMessage, 0),
		}, nil
	}

	path := filepath.Join(SessionDir, safeName(id))
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return createSession(id, path)
		}
		return nil, err
	}
	var s Session
	err = json.Unmarshal(b, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func createSession(id string, path string) (*Session, error) {
	s := &Session{
		ID:       id,
		Messages: make([]openai.ChatCompletionMessage, 0),
	}
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(path, b, 0o666)
	if err != nil {
		return nil, err
	}
	return s, nil
}

type Session struct {
	tmp      bool
	ID       string
	Messages []openai.ChatCompletionMessage
}

func (s *Session) String() string {
	b := strings.Builder{}
	for _, m := range s.Messages {
		b.WriteString(m.Role + ": " + m.Content + "\n")
	}
	return b.String()
}

func (s *Session) UseRole(r *Role) {
	if len(s.Messages) > 0 {
		return
	}
	s.Append(r.ToMsg())
}

func (s *Session) Get() []openai.ChatCompletionMessage {
	return s.Messages
}

func (s *Session) Save() error {
	if s.tmp {
		return nil
	}

	b, err := json.Marshal(s)
	if err != nil {
		return err
	}
	path := filepath.Join(SessionDir, safeName(s.ID))
	return os.WriteFile(path, b, 0o666)
}

func (s *Session) Append(msg ...openai.ChatCompletionMessage) {
	s.Messages = append(s.Messages, msg...)
}
