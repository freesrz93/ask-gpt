package consts

import (
	"fmt"
	"runtime"
)

const AppName = "ask-gpt"

var (
	Version = "0.0.0"
	BuiltBy = "unknown"
	BuiltAt = "unknown"
	VerInfo = fmt.Sprintf("%s %s\nRuntime: %s/%s\nBuiltBy: %s\nBuiltAt: %s\n", AppName, Version, runtime.GOOS, runtime.GOARCH, BuiltBy, BuiltAt)
)
