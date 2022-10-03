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
	tryInstallCarvelCLIs(t)
	tryInstallKindCLI(t)
	trySetupRegistryAsLocalDockerContainer(t)
	tryCutAppBundle(t)
	tryCutPackageRelease(t)
	trySetupKindCluster(t)
	tryVerifyApplication(t)
}

func TestCarvelPublish(t *testing.T) {
	tryInstallCarvelCLIs(t)
	tryCutAppBundle(t)
	tryCutPackageRelease(t)
}
