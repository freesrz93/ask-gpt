package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	"github.com/freesrz93/ask-gpt/consts"
)

const (
	configFile    = "config.yaml"
	sessionDir    = "sessions"
	roleDir       = "roles"
	tempSession   = "temp"
	defaultRole   = "default"
	defaultPrompt = "You are a polymath. Your role is to synthesize accurate information from various domains while offering insightful analysis and explanations. When responding, strive for clarity and depth, and encourage further inquiry by providing context and related concepts."

	AIPrefix   = "Assistant: "
	UserPrefix = "User: "
)

var CfgDir = func() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(home, ".config", consts.AppName)
}()

func safeName(raw string) string {
	return base64.URLEncoding.EncodeToString([]byte(raw))
}

func P(s string) {
	_, _ = os.Stdout.WriteString(s)
}

func Pln() {
	P("\n")
}

func PErr(v any) {
	P("error: " + fmt.Sprint(v))
}

func PFatal(v any) {
	PErr(v)
	os.Exit(1)
}
