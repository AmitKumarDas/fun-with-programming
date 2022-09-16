package shellx_carvel

import (
	"fmt"
	"github.com/magefile/mage/sh"
)

// This file provides functions that cater to releasing a Carvel package

// TODO: Should these be environment variables?
func getDefaultReleaseDir() string {
	return joinPaths(getDefaultCarvelPackagingDir(), "release")
}

func getDefaultPackageImgpkgDir() string {
	return joinPaths(getDefaultReleaseDir(), ".imgpkg")
}

func getDefaultPackageRepoDir() string {
	return joinPaths(getDefaultReleaseDir(), "packages")
}

func getDefaultPackageDir() string {
	return joinPaths(getDefaultPackageRepoDir(), EnvPackageName)
}

func getDefaultPackageMetadataFile() string {
	return joinPaths(getDefaultPackageDir(), "package-metadata.yml")
}

func getDefaultPackageTemplateFile() string {
	return joinPaths(getDefaultPackageDir(), "package-template.yml")
}

func getDefaultPackageVersionFile() string {
	return joinPaths(getDefaultPackageDir(), EnvPackageVersion+".yml")
}

func getDefaultPackageImgpkgFile() string {
	return joinPaths(getDefaultPackageImgpkgDir(), "images.yml")
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
	out, outErr := sh.Output(format("%s/ytt", EnvBinPathCarvel), "-f", inputValuesFile, "--data-values-schema-inspect", "-o", "openapi-v3")
	if outErr != nil {
		return outErr
	}
	return file(outputValuesFile, out+"\n", 0644)
}

func createPackageTemplate(pkgTemplateFile string) error {
	return file(pkgTemplateFile, packageTemplateYML, 0644)
}

func createPackageVersion(pkgTemplateFile, pkgValuesFile, pkgVersion, pkgVersionFile string) error {
	out, outErr := sh.Output(format("%s/ytt", EnvBinPathCarvel), "-f", pkgTemplateFile, "--data-value-file", "openapi="+pkgValuesFile, "-v", "version="+pkgVersion)
	if outErr != nil {
		return outErr
	}
	return file(pkgVersionFile, out+"\n", 0644)
}

func createPackageRepoBundle(pkgRepoDir, pkgImgpkgFile string) error {
	return kbld("-f", pkgRepoDir, "--imgpkg-lock-output", pkgImgpkgFile)
}

func publishPackageRepoBundle(releaseDir string) error {
	return imgpkg("push", "-b", format("%s:%s/packages/%s:%s", EnvRegistryName, EnvRegistryPort, EnvPackageRepoName, EnvPackageRepoVersion), "-f", releaseDir)
}
