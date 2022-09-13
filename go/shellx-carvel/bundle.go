package shellx_carvel

// This file provides functions that cater to Carvel related packaging

func createAppConfigs(dir string) error {
	if err := mkdir(dir); err != nil {
		return err
	}
	if err := file(dir+"/config.yml", appDeploymentYML, 0644); err != nil {
		return err
	}
	return file(dir+"/values.yml", appValuesYML, 0644)
}
