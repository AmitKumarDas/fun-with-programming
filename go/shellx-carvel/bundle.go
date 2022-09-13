package shellx_carvel

// This file provides functions that cater to Carvel related packaging

func getDefaultAppSourceDir() string {
	return "package/source"
}

func getDefaultAppConfigDir() string {
	return getDefaultAppSourceDir() + "/config"
}

func getDefaultAppImgpkgDir() string {
	return getDefaultAppSourceDir() + "/.imgpkg"
}

func createAppConfigs(configDir string) error {
	if err := mkdir(configDir); err != nil {
		return err
	}
	if err := file(configDir+"/config.yml", appDeploymentYML, 0644); err != nil {
		return err
	}
	return file(configDir+"/values.yml", appValuesYML, 0644)
}

func createAppBundle(imgpkgDir, configDir string) error {
	if err := mkdir(imgpkgDir); err != nil {
		return err
	}
	return kbld("-f", configDir, "--imgpkg-lock-output", imgpkgDir+"/images.yml")
}

func publishAppBundle(sourceDir string) error {
	return imgpkg("push", "-b", "${REGISTRY_NAME}:${REGISTRY_PORT}/packages/${APP_BUNDLE_NAME}:${APP_BUNDLE_VERSION}", "-f", sourceDir)
}
