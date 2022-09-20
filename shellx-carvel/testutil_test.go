package shellx_carvel

import (
	shx "carvel.shellx.dev/internal/sh"
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

func requireContains(t *testing.T, expectedSubStr, actual string) {
	if strings.Contains(actual, expectedSubStr) {
		return
	}
	t.Fatalf("expected substring %q is not part of actual %q", expectedSubStr, actual)
}

func mustJoinPaths(paths ...string) string {
	finalPath, err := shx.JoinPaths(paths...)
	if err != nil {
		panic(err)
	}
	return finalPath
}
