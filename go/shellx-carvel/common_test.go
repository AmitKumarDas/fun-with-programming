package shellx_carvel

import "testing"

func requireNoErr(t *testing.T, err error) {
	if err == nil {
		return
	}
	t.Fatal(err)
}

func requireTrue(t *testing.T, given bool) {
	if given {
		return
	}
	t.Fatal("expected true got false")
}
