package sh

import (
	"fmt"
	"github.com/magefile/mage/sh"
	"os"
	"strings"
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

type InvalidArgError struct{ Msg string }

func (e *InvalidArgError) Error() string { return e.Msg }

// ExpandEnvStrict replaces ${var} or $var in the string. It returns
// error for invalid use of $ or env expansion returns empty
//
// Note: Borrowed from os.ExpandEnv
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
				return "", &InvalidArgError{fmt.Sprintf("invalid %s", s)}
			} else if name == "" {
				// Valid syntax, but $ was not followed by a
				// name. Leave the dollar character untouched.
				buf = append(buf, s[j])
			} else {
				val, found := os.LookupEnv(name)
				if !found {
					return "", &InvalidArgError{fmt.Sprintf("%s lookup failed", name)}
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

// VerifyArgs returns error if the arguments makes use of unset
// environment variables
func VerifyArgs(args ...string) error {
	if len(args) == 0 {
		return nil
	}
	var invalidArgs = make([]string, 0, len(args))
	for _, arg := range args {
		if _, err := ExpandEnvStrict(arg); err != nil {
			invalidArgs = append(invalidArgs, err.Error())
		}
	}
	if len(invalidArgs) == 0 {
		return nil
	}
	return &InvalidArgError{fmt.Sprintf("verify args ['%s']: %s", strings.Join(args, "', '"), strings.Join(invalidArgs, ", "))}
}

// RunCmdStrict is a wrapper over sh.Run. The returned function
// returns error if either verification of arguments fails or
// command execution fails
func RunCmdStrict(cmd string, args ...string) func(args ...string) error {
	return func(args2 ...string) error {
		allArgs := append(args, args2...)
		if err := VerifyArgs(allArgs...); err != nil {
			return err
		}
		return sh.Run(cmd, allArgs...)
	}
}

func File(name, data string, perm os.FileMode) error {
	if err := VerifyArgs(data); err != nil {
		return err
	}
	out, outErr := sh.Output("echo", data)
	if outErr != nil {
		return outErr
	}
	return os.WriteFile(name, []byte(out), perm)
}
