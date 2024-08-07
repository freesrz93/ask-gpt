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
	BaseURL          string  `yaml:"url"`
	APIKey           string  `yaml:"api_key"`
	Model            string  `yaml:"model"`
	MaxTokens        int     `yaml:"max_tokens"`
	Temperature      float32 `yaml:"temperature"`
	TopP             float32 `yaml:"top_p"`
	FrequencyPenalty float32 `yaml:"frequency_penalty"`
	PresencePenalty  float32 `yaml:"presence_penalty"`
	Editor           string  `yaml:"editor"`
	EditorArg        string  `yaml:"editor_arg"`
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
		BaseURL:          "https://api.openai.com/v1",
		APIKey:           "",
		Model:            "gpt-4o-mini",
		MaxTokens:        4096,
		Temperature:      0.5,
		TopP:             1.0,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Editor:           "code",
		EditorArg:        "%path",
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
	return os.WriteFile(ConfigPath, b, os.ModePerm)
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
