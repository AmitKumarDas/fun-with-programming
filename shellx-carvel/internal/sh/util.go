package sh

import (
	"fmt"
	"os"
	"path"
)

// isAlphaNum reports whether the byte is an ASCII letter, number, or underscore
//
// Note: Borrowed from os package
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
				// Encountered invalid syntax
				return "", &InvalidEnvError{
					Context:     fmt.Sprintf("failed to expand [%s]", s),
					InvalidEnvs: []string{s[j+1:]},
				}
			} else if name == "" {
				// Valid syntax, but $ was not followed by a
				// name. Leave the dollar character untouched.
				buf = append(buf, s[j])
			} else {
				val, found := os.LookupEnv(name)
				if !found {
					return "", &InvalidEnvError{
						Context:     fmt.Sprintf("failed to lookup [%s]", s),
						InvalidEnvs: []string{name},
					}
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

// ExpandAllStrict returns error if the provided data makes use of unset
// environment variables
func ExpandAllStrict(data ...string) ([]string, error) {
	var invalid = make([]error, 0, len(data))
	var expanded = make([]string, 0, len(data))
	for _, d := range data {
		exp, err := ExpandEnvStrict(d)
		if err != nil {
			invalid = append(invalid, err)
			continue
		}
		expanded = append(expanded, exp)
	}
	if len(invalid) == 0 {
		return expanded, nil
	}
	return nil, &MultiError{Errors: invalid}
}

func File(name, data string, perm os.FileMode) error {
	expanded, err := ExpandAllStrict(name, data)
	if err != nil {
		return err
	}
	return os.WriteFile(expanded[0], []byte(expanded[1]), perm)
}

func JoinPaths(paths ...string) (string, error) {
	if len(paths) == 0 {
		return "", fmt.Errorf("no paths given")
	}
	expanded, err := ExpandAllStrict(paths...)
	if err != nil {
		return "", err
	}
	return path.Join(expanded...), nil
}
