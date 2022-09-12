package shellx_carvel

func createConfigYML() error {
	if err := mkdir("config"); err != nil {
		return err
	}
	return file("config/config.yml", configYML, 0644)
}
