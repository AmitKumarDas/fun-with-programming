package tests

import (
	"fmt"
	"testing"
)

func TestExpandEnvStrict(t *testing.T) {
	// mock env k:v pairs
	EnvWhoAmI := "${WHO_AM_I}"
	EnvVerbose := "${VERBOSE}"
	_ = maybeSetEnv(EnvWhoAmI, "none")
	_ = maybeSetEnv(EnvVerbose, "false")
	defer func() {
		unsetEnv(EnvWhoAmI)
		unsetEnv(EnvVerbose)
	}()
	var scenarios = []struct {
		name      string
		data      string
		isErr     bool
		isOutEqIn bool
	}{
		{
			name:  "verify error given an unset env",
			data:  "Hi ${there}",
			isErr: true,
		},
		{
			name:  "verify error given an env with missing }",
			data:  "Hi ${there",
			isErr: true,
		},
		{
			name:      "verify no error given a var that is not env but has $",
			data:      "Hi $ How are you",
			isErr:     false,
			isOutEqIn: true,
		},
		{
			name:      "verify no error given a var with $*",
			data:      "Hi $* How are you",
			isErr:     false,
			isOutEqIn: true,
		},
		{
			name:      "verify no error given a var with $#",
			data:      "Hi $# How are you",
			isErr:     false,
			isOutEqIn: true,
		},
		{
			name:      "verify no error given a var with $?",
			data:      "Hi $? How are you",
			isErr:     false,
			isOutEqIn: true,
		},
		{
			name:      "verify no error given a var with $$",
			data:      "Hi $$ How are you",
			isErr:     false,
			isOutEqIn: true,
		},
		{
			name:      "verify no error given a var with $!",
			data:      "Hi $! How are you",
			isErr:     false,
			isOutEqIn: true,
		},
		{
			name:      "verify no error given a var with $@",
			data:      "Hi $@ How are you",
			isErr:     false,
			isOutEqIn: true,
		},
		{
			name:      "verify no error given a set env",
			data:      "Hello " + EnvWhoAmI,
			isErr:     false,
			isOutEqIn: false,
		},
		{
			name:      "verify no error given multiple set envs",
			data:      fmt.Sprintf("Hello %s; My verbosity is %s", EnvWhoAmI, EnvVerbose),
			isErr:     false,
			isOutEqIn: false,
		},
	}
	for _, s := range scenarios {
		s := s
		t.Run(s.name, func(t *testing.T) {
			got, err := ExpandEnvStrict(s.data)
			if s.isErr {
				requireErr(t, err)
				return
			} else {
				requireNoErr(t, err)
				requireNotEmpty(t, got)
			}
			if s.isOutEqIn {
				requireEqual(t, s.data, got)
			} else {
				requireNotEqual(t, s.data, got)
			}
		})
	}
}
