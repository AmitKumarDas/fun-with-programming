package shellx_carvel

import (
	"fmt"
	"github.com/magefile/mage/sh"
	"testing"
)

func tryPackageRelease(t *testing.T) {
	sourceDir := "tmp/packaging/source"
	releaseDir := "tmp/packaging/release"

	// config related
	configDir := sourceDir + "/config"
	// release related
	pkgRepoDir := releaseDir + "/packages"
	pkgImgpkgDir := releaseDir + "/.imgpkg"
	pkgDir := pkgRepoDir + "/" + getEnvTrimKey(EnvPackageName)
	pkgVersion := getEnvTrimKey(EnvPackageVersion)

	requireNoErr(t, createReleaseDirs(pkgDir, pkgImgpkgDir))
	requireTrue(t, exists(pkgDir))
	requireTrue(t, exists(pkgImgpkgDir))

	requireNoErr(t, createPackageMetadata(pkgDir+"/package-metadata.yml"))
	requireTrue(t, exists(pkgDir+"/package-metadata.yml"))

	requireNoErr(t, generateOpenAPISchema(configDir+"/values.yml", releaseDir+"/schema-openapi.yml"))
	requireTrue(t, exists(releaseDir+"/schema-openapi.yml"))

	requireNoErr(t, createPackageTemplate(releaseDir+"/package-template.yml"))
	requireTrue(t, exists(releaseDir+"/package-template.yml"))

	pkgVersionFile := pkgDir + "/" + pkgVersion + ".yml"
	requireNoErr(t, createPackageVersion(releaseDir+"/package-template.yml", releaseDir+"/schema-openapi.yml", pkgVersion, pkgVersionFile))
	requireTrue(t, exists(pkgVersionFile))

	requireNoErr(t, createPackageRepoBundle(pkgRepoDir, pkgImgpkgDir+"/images.yml"))
	requireTrue(t, exists(pkgImgpkgDir+"/images.yml"))

	requireNoErr(t, publishPackageRepoBundle(releaseDir))
	out, outErr := sh.Output("curl", "${REGISTRY_NAME}:${REGISTRY_PORT}/v2/_catalog")
	requireNoErr(t, outErr)
	requireContains(t, out, fmt.Sprintf("packages/%s", getEnvTrimKey(EnvPackageRepoName)))
}
