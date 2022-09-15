package shellx_carvel

import "testing"

func tryInstallKindCLI(t *testing.T) {
	requireNoErr(t, installKindCLI())
	requireNoErr(t, ls(EnvBinPathKind))
	requireTrue(t, isNoErr(whichKind()))
}
