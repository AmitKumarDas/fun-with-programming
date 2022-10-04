package sh

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// passThroughFn returns the provided input. It is useful
// as a custom mapper function for os.Expand
func passThroughFn(in string) string {
	return in
}

func MaybeSetEnv(envKey, defaultVal string) string {
	// This treats the expanded forms such as "$key" or "${key}" if any
	// to "key". There is no change to key chars if it was not in any
	// of above expansion forms.
	trimmedKey := os.Expand(envKey, passThroughFn)
	// set default only if provided env key is not set
	if value := os.Getenv(trimmedKey); value == "" {
		_ = os.Setenv(trimmedKey, defaultVal)
	}
	return os.Getenv(trimmedKey)
}

func UnsetEnv(envKey string) {
	_ = os.Unsetenv(os.Expand(envKey, passThroughFn))
}

// isAlphaNum reports whether the byte is an ASCII letter, number, or underscore.
// These are the valid characters to define an ENV variable
//
// Note: Borrowed from os package
func isAlphaNum(c uint8) bool {
	return c == '_' || '0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}

// isShellSpecialVar reports whether the character identifies a special
// shell variable such as $*
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

// ExpandStrict replaces ${var} or $var in the string. It returns
// error for invalid use of $ or env expansion returns empty. This
// is mostly useful in env substitution of file content without
// the need for additional tools
//
// Note: Borrowed from os.Expand
func ExpandStrict(s string) (string, error) {
	var buf []byte
	var removals []string
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
				// Encountered invalid syntax
				return "", &InvalidEnvError{
					Context:     fmt.Sprintf("parse [%s]", s[j:]),
					InvalidEnvs: []string{s[j+1:]},
				}
			} else if name == "" {
				// Valid syntax, but $ was not followed by a
				// name. Leave the dollar character untouched.
				buf = append(buf, s[j])
			} else if len(name) == 1 && isShellSpecialVar(name[0]) {
				// A shell special variable. E.g.: $#, $@, $%, $?, ${?} etc.
				// Do nothing i.e. dollar & special character are removed
				if IsDebug() {
					if removals == nil {
						removals = make([]string, 0, len(s))
					}
					removals = append(removals, s[j:j+w+1])
				}
			} else {
				// This is most likely an ENV variable
				val, found := os.LookupEnv(name)
				if !found {
					return "", &InvalidEnvError{
						Context:     fmt.Sprintf("lookup [%s]", s[j:]),
						InvalidEnvs: []string{name},
					}
				}
				buf = append(buf, val...)
			}
			j += w
			i = j + 1
		}
	}
	if IsDebug() && len(removals) != 0 {
		log.Println("given:", s, "removed:", strings.Join(removals, " "))
	}
	if buf == nil {
		return s, nil
	}
	return string(buf) + s[i:], nil
}

// ExpandStrictAll returns error if the provided data makes use of unset
// environment variables
func ExpandStrictAll(data ...string) ([]string, error) {
	var invalid = make([]error, 0, len(data))
	var expanded = make([]string, 0, len(data))
	for _, d := range data {
		exp, err := ExpandStrict(d)
		if err != nil {
			invalid = append(invalid, err)
			continue
		}
		expanded = append(expanded, exp)
	}
	if len(invalid) == 0 {
		return expanded, nil
	}
	return nil, (&Error{}).AddAll(invalid)
}
