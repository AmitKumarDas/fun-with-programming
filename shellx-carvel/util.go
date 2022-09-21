package shellx_carvel

import (
	shx "carvel.shellx.dev/internal/sh"
	"fmt"
	"os"
	"strings"
	"time"
)

// Unix & generic commands as functions
var mkdir = shx.RunCmd("mkdir", "-p")
var curl = shx.RunCmd("curl")
var ls = shx.RunCmd("ls")
var chmod = shx.RunCmd("chmod")

// Docker CLI as function
var docker = shx.RunCmd("docker")

// kubectl CLI as function
var kubectl = shx.RunCmd("kubectl")

// File creation as a function
var file = shx.File

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
		_ = os.Setenv(os.Expand(envKey, passThroughFn), defaultVal)
	}
	return os.ExpandEnv(envKey)
}

func getEnv(envKey string) string {
	return os.ExpandEnv(envKey)
}

func exists(file string) bool {
	return ls(file) == nil
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

type eventuallyConfig struct {
	Attempts *int
	Interval *time.Duration
}

// eventually runs the given fn till it succeeds or eventually
// function times out
func eventually(fn func() error) error {
	return eventuallyWith(fn, eventuallyConfig{})
}

// eventuallyWith runs the given fn till it succeeds or eventuallyWith
// function times out
func eventuallyWith(fn func() error, config eventuallyConfig) error {
	if fn == nil {
		return fmt.Errorf("nil function")
	}
	attempts := 10
	interval := 3 * time.Second
	if config.Attempts != nil && *config.Attempts != 0 {
		attempts = *config.Attempts
	}
	if config.Interval != nil && config.Interval.Seconds() != 0 {
		interval = *config.Interval
	}
	var start = time.Now()
	var final error
	for counter := 1; counter <= attempts; counter++ {
		curErr := fn()
		if curErr == nil {
			return nil
		}
		final = curErr
		time.Sleep(interval)
	}
	return fmt.Errorf("%w: func timed out after %s", final, time.Now().Sub(start))
}
