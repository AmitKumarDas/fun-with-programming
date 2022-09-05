package main

import (
	"github.com/rogpeppe/go-internal/testscript"
	"testing"
)

func TestScript(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata",
	})
}
