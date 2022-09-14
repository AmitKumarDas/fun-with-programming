package shellx_carvel

import (
	"fmt"
	"github.com/magefile/mage/sh"
)

// This file provides functions that cater to releasing a Carvel package

func getDefaultReleaseDir() string {
	return "package/release"
}

func getDefaultPackageImgpkgDir() string {
	return getDefaultReleaseDir() + "/.imgpkg"
}

func getDefaultPackageRepoDir() string {
	return getDefaultReleaseDir() + "/packages"
}

func getDefaultPackageDir() string {
	return getDefaultPackageRepoDir() + "/" + getEnvTrimKey(EnvPackageName)
}

func getDefaultPackageMetadataFile() string {
	return getDefaultPackageDir() + "/package-metadata.yml"
}

func getDefaultPackageTemplateFile() string {
	return getDefaultPackageDir() + "/package-template.yml"
}

func getDefaultPackageVersionFile() string {
	return getDefaultPackageDir() + "/" + getEnvTrimKey(EnvPackageVersion) + ".yml"
}

func getDefaultPackageImgpkgFile() string {
	return getDefaultPackageImgpkgDir() + "/images.yml"
}

func createReleaseDirs(pkgDir, pkgImgpkgDir string) error {
	if err := mkdir(pkgDir); err != nil {
		return err
	}
	return mkdir(pkgImgpkgDir)
}

func createPackageMetadata(pkgMetadataFile string) error {
	return file(pkgMetadataFile, packageMetadataYML, 0644)
}

func generateOpenAPISchema(inputValuesFile, outputValuesFile string) error {
	if !exists(inputValuesFile) {
		return fmt.Errorf("file %s not found", inputValuesFile) // TODO: Check if this error message is helpful
	}
	out, outErr := sh.Output("${BIN_PATH_CARVEL}/ytt", "-f", inputValuesFile, "--data-values-schema-inspect", "-o", "openapi-v3")
	if outErr != nil {
		return outErr
	}
	return file(outputValuesFile, out+"\n", 0644)
}

func createPackageTemplate(pkgTemplateFile string) error {
	return file(pkgTemplateFile, packageTemplateYML, 0644)
}

func createPackageVersion(pkgTemplateFile, pkgValuesFile, pkgVersion, pkgVersionFile string) error {
	out, outErr := sh.Output("${BIN_PATH_CARVEL}/ytt", "-f", pkgTemplateFile, "--data-value-file", "openapi="+pkgValuesFile, "-v", "version="+pkgVersion)
	if outErr != nil {
		return outErr
	}
	return file(pkgVersionFile, out+"\n", 0644)
}

func createPackageRepoBundle(pkgRepoDir, pkgImgpkgFile string) error {
	return kbld("-f", pkgRepoDir, "--imgpkg-lock-output", pkgImgpkgFile)
}

func publishPackageRepoBundle(releaseDir string) error {
	return imgpkg("push", "-b", "${REGISTRY_NAME}:${REGISTRY_PORT}/packages/${PACKAGE_REPO_NAME}:${PACKAGE_REPO_VERSION}", "-f", releaseDir)
}
