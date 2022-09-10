package shellx_carvel

func setupRegistryAsLocalDockerContainer() error {
	if err := docker("inspect", "-f", "{{.State.Running}}", "${REGISTRY_NAME}"); err != nil {
		return docker("run", "-d", "--restart", "always", "-p", "127.0.0.1:${REGISTRY_PORT}:5000", "--name", "${REGISTRY_NAME}", "registry:2")
	}
	return nil
}
