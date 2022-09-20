package tests

import (
	"os"
	"strings"
	"testing"
)

func requireCount(t *testing.T, expected int, actual int) {
	if expected == actual {
		return
	}
	t.Fatalf("expected %d got %d", expected, actual)
}

func requireErr(t *testing.T, err error) {
	if err != nil {
		return
	}
	t.Fatal("expected error got none")
}

func requireNoErr(t *testing.T, err error) {
	if err == nil {
		return
	}
	t.Fatalf("expected no err got: %+v", err)
}

func requireTrue(t *testing.T, actual bool) {
	if actual {
		return
	}
	t.Fatal("expected true got false")
}

func requireEqual(t *testing.T, expected, actual string) {
	if expected == actual {
		return
	}
	t.Fatalf("expected %q actual %q", expected, actual)
}

func requireNotEqual(t *testing.T, expected, actual string) {
	if expected != actual {
		return
	}
	t.Fatalf("expected is same as actual %q", actual)
}

func requireContains(t *testing.T, expectedSubStr, actual string) {
	if strings.Contains(actual, expectedSubStr) {
		return
	}
	t.Fatalf("expected substring %q is not part of actual %q", expectedSubStr, actual)
}

func requireNotEmpty(t *testing.T, actual string) {
	if actual != "" {
		return
	}
	t.Fatalf("expected non empty got empty")
}

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

func unsetEnv(envKey string) {
	_ = os.Unsetenv(os.Expand(envKey, passThroughFn))
}
