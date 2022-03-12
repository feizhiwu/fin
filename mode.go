package fin

import (
	"os"
)

const EnvFinMode = "FIN_MODE"

const (
	DebugMode   = "debug"
	ReleaseMode = "release"
	TestMode    = "test"
)

var modeName = DebugMode

func init() {
	mode := os.Getenv(EnvFinMode)
	SetMode(mode)
}

func SetMode(value string) {
	switch value {
	case DebugMode, "":
	case ReleaseMode:
	case TestMode:
	default:
		panic("fin mode unknown: " + value)
	}
	if value == "" {
		value = DebugMode
	}
	modeName = value
}

func Mode() string {
	return modeName
}
