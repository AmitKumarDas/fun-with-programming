package shellx_carvel

import "testing"

func TestReleaseCarvelPackage(t *testing.T) {
	tryInstallCarvelCLIs(t)
	tryInstallKindCLI(t)
	tryAppBundleCreateAndPublish(t)
	tryPackageRelease(t)
}
