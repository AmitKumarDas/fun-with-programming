package shellx_carvel

import (
	"testing"
)

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

func TestInstallCarvelCLIs(t *testing.T) {
	requireNoErr(t, installCarvelCLIs())
	requireNoErr(t, ls(EnvBinPathCarvel))
	requireTrue(t, isNoErr(whichKbld(), whichYtt(), whichImgpkg()))
}

func TestInstallKindCLI(t *testing.T) {
	requireNoErr(t, installKind())
	requireNoErr(t, ls(EnvBinPathKind))
	requireTrue(t, isNoErr(whichKind()))
}
