package shellx_carvel

import "testing"

func TestReleaseOfCarvelPackage(t *testing.T) {

	// Note:
	// -------------------------------------------------------
	// Naming convention followed for each subtest:
	// Running this testcase will verify {name of subtest}
	//
	// Where {name of subtest} is mentioned as shown below:
	// t.Run("name of subtest", func(t *testing.T){...})
	// -------------------------------------------------------

	t.Run("installing carvel CLIs", func(t *testing.T) {
		tryInstallCarvelCLIs(t)
	})
	t.Run("installing KIND cli", func(t *testing.T) {
		tryInstallKindCLI(t)
	})
	t.Run("setting up local docker container as an OCI registry", func(t *testing.T) {
		trySetupRegistryAsLocalDockerContainer(t)
	})
	t.Run("publishing the app bundle", func(t *testing.T) {
		tryCutAppBundle(t)
	})
	t.Run("publishing the package release", func(t *testing.T) {
		tryCutPackageRelease(t)
	})
	t.Run("setting up KIND cluster", func(t *testing.T) {
		trySetupKindCluster(t)
	})
	t.Run("verifying app deployment", func(t *testing.T) {
		tryVerifyApplication(t)
	})
}
