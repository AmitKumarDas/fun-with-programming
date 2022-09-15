package shellx_carvel

// This file provides functions that cater to Carvel related packaging

// TODO: Should these be environment variables?
func getDefaultCarvelPackagingDir() string {
	return "packaging"
}

func getDefaultSourceDir() string {
	return getDefaultCarvelPackagingDir() + "/source"
}

func getDefaultAppConfigDir() string {
	return getDefaultSourceDir() + "/config"
}

func getDefaultAppImgpkgDir() string {
	return getDefaultSourceDir() + "/.imgpkg"
}

func getDefaultAppConfigFile() string {
	return getDefaultAppConfigDir() + "/config.yml"
}

func getDefaultAppValuesFile() string {
	return getDefaultAppConfigDir() + "/values.yml"
}

func getDefaultAppImgpkgFile() string {
	return getDefaultAppImgpkgDir() + "/images.yml"
}

func createAppDirs(appConfigDir, appImgpkgDir string) error {
	if err := mkdir(appConfigDir); err != nil {
		return err
	}
	return mkdir(appImgpkgDir)
}

func createAppConfigs(appConfigFile, appConfigValuesFile string) error {
	if err := file(appConfigFile, appDeploymentYML, 0644); err != nil {
		return err
	}
	return file(appConfigValuesFile, appValuesYML, 0644)
}

func createAppBundle(appConfigDir, appImgpkgFile string) error {
	return kbld("-f", appConfigDir, "--imgpkg-lock-output", appImgpkgFile)
}

func publishAppBundle(sourceDir string) error {
	return imgpkg("push", "-b", "${REGISTRY_NAME}:${REGISTRY_PORT}/packages/${APP_BUNDLE_NAME}:${APP_BUNDLE_VERSION}", "-f", sourceDir)
}
