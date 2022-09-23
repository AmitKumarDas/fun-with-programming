package shellx_carvel

import (
	shx "carvel.shellx.dev/internal/sh"
	"testing"
)

func tryInstallCarvelCLIs(t *testing.T) {
	requireNoErr(t, installCarvelCLIs())
	requireNoErr(t, ls(EnvBinPathCarvel))
	requireTrue(t, isNoErr(whichKbld(), whichYtt(), whichImgpkg()))
}

func tryPackageRelease(t *testing.T) {
	requireNoErr(t, cutCarvelRelease())

	out, outErr := shx.Output("curl", format("%s:%s/v2/_catalog", EnvRegistryName, EnvRegistryPort))
	requireNoErr(t, outErr)
	var mErr shx.MultiError
	assertContains(t, shx.JoinPathsWithErrHandle(&mErr, "packages", EnvPackageRepoName), out)
	requireNoErr(t, (&mErr).ErrOrNil())
}
