package shellx_carvel

import "testing"

func tryInstallKindCLI(t *testing.T) {
	requireNoErr(t, installKindCLI())
	requireNoErr(t, ls(EnvBinPathKind))
	requireTrue(t, isNoErr(whichKind()))
}

func trySetupKindCluster(t *testing.T) {
	requireNoErr(t, setupKindCluster())
}
