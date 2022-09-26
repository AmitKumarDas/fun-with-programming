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

func tryCutAppBundle(t *testing.T) {
	requireNoErr(t, cutAppBundle())

	out, outErr := shx.Output("curl", format("%s:%s/v2/_catalog", EnvRegistryName, EnvRegistryPort))
	requireNoErr(t, outErr)
	var mErr shx.MultiError
	assertContains(t, shx.JoinPathsWithErrHandle(&mErr, "packages", EnvAppBundleName), out)
	requireNoErr(t, (&mErr).ErrOrNil())
}

func tryCutPackageRelease(t *testing.T) {
	requireNoErr(t, cutCarvelRelease())

	out, outErr := shx.Output("curl", format("%s:%s/v2/_catalog", EnvRegistryName, EnvRegistryPort))
	requireNoErr(t, outErr)
	var mErr shx.MultiError
	assertContains(t, shx.JoinPathsWithErrHandle(&mErr, "packages", EnvPackageRepoName), out)
	requireNoErr(t, (&mErr).ErrOrNil())
}
