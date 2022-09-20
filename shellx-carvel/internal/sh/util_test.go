package sh

import (
	"fmt"
	"github.com/magefile/mage/sh"
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
			got, err := ExpandStrict(s.data)
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

func TestFile(t *testing.T) {
	requireNoErr(t, sh.Run("mkdir", "-p", "tmp"))
	requireErr(t, File("tmp/will-not-create.txt", "I $WILL_NOT_EXIST due to this unset env\n", 0644))
	requireNoErr(t, File("tmp/will-create.txt", "I WILL EXIST\n", 0644))
}
