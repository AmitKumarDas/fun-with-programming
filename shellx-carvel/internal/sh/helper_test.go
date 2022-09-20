package sh

import "testing"

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
	t.Fatal(err)
}

func requireTrue(t *testing.T, actual bool) {
	if actual {
		return
	}
	t.Fatal("expected true got false")
}

func requireNotEmpty(t *testing.T, actual string) {
	if actual != "" {
		return
	}
	t.Fatalf("expected non empty got empty")
}

func requireNotEqual(t *testing.T, expected, actual string) {
	if expected != actual {
		return
	}
	t.Fatalf("expected is same as actual %q", actual)
}

func requireEqual(t *testing.T, expected, actual string) {
	if expected == actual {
		return
	}
	t.Fatalf("expected %q actual %q", expected, actual)
}
