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
