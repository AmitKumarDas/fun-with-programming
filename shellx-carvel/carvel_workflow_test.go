package shellx_carvel

import "testing"

func TestReleaseCarvelPackage(t *testing.T) {
	tryInstallCarvelCLIs(t)
	tryInstallKindCLI(t)
	trySetupRegistryAsLocalDockerContainer(t)
	tryAppBundleCreateAndPublish(t)
	tryPackageRelease(t)
}
