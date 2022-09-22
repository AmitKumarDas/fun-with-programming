package sh

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

func TestExpandEnvStrict(t *testing.T) {
	// mock env k:v pairs
	EnvWhoAmI := "${WHO_AM_I}"
	EnvVerbose := "${VERBOSE}"
	_ = MaybeSetEnv(EnvWhoAmI, "none")
	_ = MaybeSetEnv(EnvVerbose, "false")
	defer func() {
		UnsetEnv(EnvWhoAmI)
		UnsetEnv(EnvVerbose)
	}()
	var scenarios = []struct {
		name     string
		data     string
		isErr    bool
		expected string
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
			name:     "verify no error given a var that is not env but has $",
			data:     "Hi $ How are you",
			isErr:    false,
			expected: "Hi $ How are you",
		},
		{
			name:     "verify no error given a var with $*",
			data:     "Hi $* How are you",
			isErr:    false,
			expected: "Hi  How are you",
		},
		{
			name:     "verify no error given a var with ${*}",
			data:     "Hi ${*} How are you",
			isErr:    false,
			expected: "Hi  How are you",
		},
		{
			name:     "verify no error given a var with $#",
			data:     "Hi $# How are you",
			isErr:    false,
			expected: "Hi  How are you",
		},
		{
			name:     "verify no error given a var with ${#}",
			data:     "Hi ${#} How are you",
			isErr:    false,
			expected: "Hi  How are you",
		},
		{
			name:     "verify no error given a var with $?",
			data:     "Hi $? How are you",
			isErr:    false,
			expected: "Hi  How are you",
		},
		{
			name:     "verify no error given a var with ${?}",
			data:     "Hi ${?} How are you",
			isErr:    false,
			expected: "Hi  How are you",
		},
		{
			name:     "verify no error given a var with $$",
			data:     "Hi $$ How are you",
			isErr:    false,
			expected: "Hi  How are you",
		},
		{
			name:     "verify no error given a var with ${$}",
			data:     "Hi ${$} How are you",
			isErr:    false,
			expected: "Hi  How are you",
		},
		{
			name:     "verify no error given a var with $!",
			data:     "Hi $! How are you",
			isErr:    false,
			expected: "Hi  How are you",
		},
		{
			name:     "verify no error given a var with ${!}",
			data:     "Hi ${!} How are you",
			isErr:    false,
			expected: "Hi  How are you",
		},
		{
			name:     "verify no error given a var with $@",
			data:     "Hi $@ How are you",
			isErr:    false,
			expected: "Hi  How are you",
		},
		{
			name:     "verify no error given a var with ${@}",
			data:     "Hi ${@} How are you",
			isErr:    false,
			expected: "Hi  How are you",
		},
		{
			name:     "verify no error given a set env",
			data:     "Hello " + EnvWhoAmI,
			isErr:    false,
			expected: "Hello " + os.ExpandEnv(EnvWhoAmI),
		},
		{
			name:     "verify no error given multiple set envs",
			data:     fmt.Sprintf("Hello %s; My verbosity is %s", EnvWhoAmI, EnvVerbose),
			isErr:    false,
			expected: fmt.Sprintf("Hello %s; My verbosity is %s", os.ExpandEnv(EnvWhoAmI), os.ExpandEnv(EnvVerbose)),
		},
	}
	for _, s := range scenarios {
		s := s
		t.Run(s.name, func(t *testing.T) {
			got, err := ExpandStrict(s.data)
			if s.isErr {
				requireErr(t, err)
				return
			}
			requireNoErr(t, err)
			requireNotEmpty(t, got)
			requireEqual(t, s.expected, got)
		})
	}
}

func TestFilterInvalidEnvs(t *testing.T) {
	var tmpEnvKey = "TESTING_FILTER_INVALID_ENVS_FUNC"
	var tmpEnvGOOS = "TMP_GOOS"
	var tmpEnvGOARCH = "TMP_GOARCH"

	os.Setenv(tmpEnvKey, "")
	os.Setenv(tmpEnvGOOS, runtime.GOOS)
	os.Setenv(tmpEnvGOARCH, runtime.GOARCH)

	defer func() {
		os.Unsetenv(tmpEnvKey)
		os.Unsetenv(tmpEnvGOOS)
	}()

	var scenarios = []struct {
		name  string
		envs  []string
		isErr bool
	}{
		{
			name:  "given all valid envs",
			envs:  []string{tmpEnvGOARCH, tmpEnvGOOS},
			isErr: false,
		},
		{
			name:  "given all valid envs & one env with empty value",
			envs:  []string{tmpEnvGOARCH, tmpEnvGOOS, "$" + tmpEnvKey},
			isErr: false,
		},
		{
			name:  "given all envs that can not be looked up",
			envs:  []string{"$NOOP", "$DONT_EXIST", "${NO}"},
			isErr: true,
		},
		{
			name:  "given all envs with invalid syntax",
			envs:  []string{"${}", "${"},
			isErr: true,
		},
		{
			name:  "given one valid & then one invalid env",
			envs:  []string{tmpEnvGOARCH, "${}"},
			isErr: true,
		},
		{
			name:  "given many valid envs & then one invalid env",
			envs:  []string{tmpEnvGOOS, tmpEnvGOARCH, "${}"},
			isErr: true,
		},
		{
			name:  "given one invalid env & then many valid envs",
			envs:  []string{"${", tmpEnvGOOS, tmpEnvGOARCH},
			isErr: true,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			_, got := ExpandStrictAll(s.envs...)
			if s.isErr {
				requireErr(t, got)
			} else {
				requireNoErr(t, got)
			}
		})
	}
}
