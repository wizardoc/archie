package env_utils

import (
	"log"
	"os"
)

const (
	DEVELOPMENT     = "development"
	PRODUCTION      = "production"
	RUNTIME_ENV_KEY = "RUNTIME_ENV"
)

var Env Environment

type Environment struct {
	RuntimeEnv string
}

func (e *Environment) IsDev() bool {
	return e.RuntimeEnv == DEVELOPMENT
}

func (e *Environment) IsProd() bool {
	return e.RuntimeEnv == PRODUCTION
}

func init() {
	// The default value of runtime env is dev
	var parsedEnv = DEVELOPMENT
	runtimeEnv := os.Getenv(RUNTIME_ENV_KEY)

	if runtimeEnv != "" {
		if runtimeEnv != DEVELOPMENT && runtimeEnv != PRODUCTION {
			log.Fatalf("Unknown value %s from RUNTIME_ENV", runtimeEnv)
			return
		}

		parsedEnv = runtimeEnv
	}

	Env = Environment{
		RuntimeEnv: parsedEnv,
	}
}
