package shellx_carvel

import (
	"testing"
)

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
