package shellx_carvel

import "testing"

func tryInstallCarvelCLIs(t *testing.T) {
	requireNoErr(t, installCarvelCLIs())
	requireNoErr(t, ls(EnvBinPathCarvel))
	requireTrue(t, isNoErr(whichKbld(), whichYtt(), whichImgpkg()))
}
