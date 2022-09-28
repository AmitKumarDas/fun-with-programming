package shellx_carvel

// Each TestXYZ function is used as an entrypoint to define a set
// of functions that cab be executed to achieve the desired outcome.
// In this case its enabling Build, Test & Release features to be
// consumed by CI pipelines.
//
// Note that this could have been a bash script or a main() function
// as well. However, `go test` CLI provides a simpler alternative to:
// 	1/ define a blueprint,
//	2/ accept environment variables, &
//	3/ be easily triggered from a Makefile
//
// The only need is to have a go environment
import "testing"

func TestCarvelReleaseE2E(t *testing.T) {

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
	t.Run("publishing the package", func(t *testing.T) {
		tryCutPackageRelease(t)
	})
	t.Run("setting up KIND cluster", func(t *testing.T) {
		trySetupKindCluster(t)
	})
	t.Run("verifying app deployment", func(t *testing.T) {
		tryVerifyApplication(t)
	})
}

func TestCarvelPublish(t *testing.T) {

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
	t.Run("publishing app bundle", func(t *testing.T) {
		tryCutAppBundle(t)
	})
	t.Run("publishing package", func(t *testing.T) {
		tryCutPackageRelease(t)
	})
}
