package tests

import (
	"fmt"
	"os"
	"testing"
)

// isAlphaNum reports whether the byte is an ASCII letter, number, or underscore
//
// Note: Borrowed from os package
// Note: These seem to be valid characters to define an ENV variable
func isAlphaNum(c uint8) bool {
	return c == '_' || '0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}

// isShellSpecialVar reports whether the character identifies a special
// shell variable such as $*.
//
// Note: Borrowed from os package
func isShellSpecialVar(c uint8) bool {
	switch c {
	case '*', '#', '$', '@', '!', '?', '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	}
	return false
}

// getShellName returns the name that begins the string and the number of bytes
// consumed to extract it. If the name is enclosed in {}, it's part of a ${}
// expansion and two more bytes are needed than the length of the name.
//
// Note: Borrowed from os package
func getShellName(s string) (string, int) {
	switch {
	case s[0] == '{':
		if len(s) > 2 && isShellSpecialVar(s[1]) && s[2] == '}' {
			return s[1:2], 3
		}
		// Scan to closing brace
		for i := 1; i < len(s); i++ {
			if s[i] == '}' {
				if i == 1 {
					return "", 2 // Bad syntax; eat "${}"
				}
				return s[1:i], i + 1
			}
		}
		return "", 1 // Bad syntax; eat "${"
	case isShellSpecialVar(s[0]):
		return s[0:1], 1
	}
	// Scan alphanumerics.
	var i int
	for i = 0; i < len(s) && isAlphaNum(s[i]); i++ {
	}
	return s[:i], i
}

type invalidEnvError struct{ Msg string }

func (e *invalidEnvError) Error() string { return e.Msg }

// ExpandEnvStrict replaces ${var} or $var in the string. It returns
// error for invalid use of $ or env expansion returns empty
//
// Note: Borrowed from os.Expand
func ExpandEnvStrict(s string) (string, error) {
	var buf []byte
	// ${} is all ASCII, so bytes are fine for this operation.
	i := 0
	for j := 0; j < len(s); j++ {
		if s[j] == '$' && j+1 < len(s) {
			if buf == nil {
				buf = make([]byte, 0, 2*len(s))
			}
			buf = append(buf, s[i:j]...)
			name, w := getShellName(s[j+1:])
			if name == "" && w > 0 {
				// Encountered invalid ENV syntax
				return "", &invalidEnvError{fmt.Sprintf("invalid syntax %s", s)}
			} else if name == "" {
				// Valid syntax, but $ was not followed by a
				// name. Leave the dollar character untouched.
				buf = append(buf, s[j])
			} else if len(name) == 1 && isShellSpecialVar(name[0]) {
				// A shell special variable.
				// Leave the dollar as well as special character untouched.
				buf = append(buf, s[j], name[0])
			} else {
				// This is most likely an ENV variable
				val, found := os.LookupEnv(name)
				if !found {
					return "", &invalidEnvError{fmt.Sprintf("lookup failed for %s", name)}
				}
				buf = append(buf, val...)
			}
			j += w
			i = j + 1
		}
	}
	if buf == nil {
		return s, nil
	}
	return string(buf) + s[i:], nil
}

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
