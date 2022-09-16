package shellx_carvel

import (
	"carvel.shellx.dev/internal/sh"
	"fmt"
	"os"
	"path"
	"strings"
)

// Unix & generic commands as functions
var mkdir = sh.RunCmdStrict("mkdir", "-p")
var curl = sh.RunCmdStrict("curl")
var ls = sh.RunCmdStrict("ls")
var chmod = sh.RunCmdStrict("chmod")

// Docker CLI as function
var docker = sh.RunCmdStrict("docker")

// File creation as a function
var file = sh.File

// passThroughFn returns the provided input. It is useful
// as a custom mapper function for os.Expand
func passThroughFn(in string) string {
	return in
}

func maybeSetEnv(envKey, defaultVal string) string {
	// set default only if provided env key is not set
	if value := os.ExpandEnv(envKey); value == "" {
		// envKey is first expanded such that "$key" or "${key}" if any
		// is trimmed to produce "key" & then this trimmed key is
		// set as an environment variable
		os.Setenv(os.Expand(envKey, passThroughFn), defaultVal)
	}
	return os.ExpandEnv(envKey)
}

func getEnv(envKey string) string {
	return os.ExpandEnv(envKey)
}

func exists(file string) bool {
	return ls(file) == nil
}

func joinPaths(elem ...string) string {
	var out = make([]string, len(elem))
	for _, i := range elem {
		out = append(out, getEnv(i))
	}
	return path.Join(out...)
}

func isErr(err error, more ...error) bool {
	if err != nil {
		return true
	}
	for _, e := range more {
		if e != nil {
			return true
		}
	}
	return false
}

func isNoErr(err error, more ...error) bool {
	return !isErr(err, more...)
}

func isEq(a, b string) bool {
	return strings.ToLower(getEnv(a)) == strings.ToLower(getEnv(b))
}

func isNotEq(a, b string) bool {
	return !isEq(a, b)
}

func format(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}
