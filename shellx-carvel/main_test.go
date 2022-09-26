package shellx_carvel

import "testing"

func TestReleaseOfCarvelPackage(t *testing.T) {
	tryInstallCarvelCLIs(t)
	tryInstallKindCLI(t)
	trySetupRegistryAsLocalDockerContainer(t)
	tryCutAppBundle(t)
	tryCutPackageRelease(t)
	trySetupKindCluster(t)
	tryVerifyApplication(t)
}
