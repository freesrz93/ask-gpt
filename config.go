package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	ConfigPath = filepath.Join(CfgDir, configFile)
	Config     = newDefault()
)

type config struct {
	DefaultBackend string                    `yaml:"default_backend"`
	Editor         string                    `yaml:"editor"`
	EditorArg      string                    `yaml:"editor_arg"`
	Backends       map[string]*BackendOption `yaml:"backends"`
	Roles          map[string]Role           `yaml:"roles"`
}

type BackendOption struct {
	Description      string  `yaml:"description"`
	DefaultRole      string  `yaml:"default_role"`
	BaseURL          string  `yaml:"url"`
	APIKey           string  `yaml:"api_key"`
	Model            string  `yaml:"model"`
	MaxTokens        int     `yaml:"max_tokens"`
	Temperature      float32 `yaml:"temperature"`
	TopP             float32 `yaml:"top_p"`
	FrequencyPenalty float32 `yaml:"frequency_penalty"`
	PresencePenalty  float32 `yaml:"presence_penalty"`
}

func (c *config) String() string {
	r, err := yaml.Marshal(c)
	if err != nil {
		return ""
	}
	return string(r)
}

func newDefault() *config {
	return &config{
		DefaultBackend: backendOpenai,
		Editor:         "code",
		EditorArg:      "%path",
		Backends: map[string]*BackendOption{
			backendOpenai: {
				Description:      "",
				DefaultRole:      defaultRole,
				BaseURL:          "https://api.openai.com/v1",
				APIKey:           "",
				Model:            "gpt-4o-mini",
				MaxTokens:        4096,
				Temperature:      0.5,
				TopP:             1.0,
				FrequencyPenalty: 0,
				PresencePenalty:  0,
			},
		},
		Roles: map[string]Role{
			defaultRole: {
				Description: "",
				Prompt:      defaultPrompt,
			}},
	}
}

func LoadCfg() {
	b, err := os.ReadFile(ConfigPath)
	if err != nil {
		if os.IsNotExist(err) {
			_ = SaveCfg()
		}
		return
	}
	err = yaml.Unmarshal(b, Config)
	if err != nil {
		return
	}
}

func SaveCfg() error {
	b, err := yaml.Marshal(Config)
	if err != nil {
		return err
	}
	return os.WriteFile(ConfigPath, b, 0o666)
}

func EditCfg() error {
	_, err := os.Lstat(ConfigPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = SaveCfg()
			if err != nil {
				return err
			}
		}
		return err
	}
	err = exec.Command(Config.Editor, strings.ReplaceAll(Config.EditorArg, "%path", ConfigPath)).Run()
	if err != nil {
		return err
	}
	return nil
}
