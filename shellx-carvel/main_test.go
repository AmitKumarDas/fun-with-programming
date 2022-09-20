package shellx_carvel

import "testing"

func TestReleaseOfCarvelPackage(t *testing.T) {
	tryInstallCarvelCLIs(t)
	tryInstallKindCLI(t)
	trySetupRegistryAsLocalDockerContainer(t)
	tryAppBundleCreateAndPublish(t)
	tryPackageRelease(t)
	trySetupKindCluster(t)
	tryVerifyApplication(t)
}
