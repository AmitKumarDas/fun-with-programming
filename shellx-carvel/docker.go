package shellx_carvel

import (
	shx "carvel.shellx.dev/internal/sh"
)

func setupRegistryAsLocalDockerContainer() error {
	if shx.IsNotEq(EnvSetupLocalRegistry, "true") {
		return nil
	}
	_ = docker("stop", EnvRegistryName)
	_ = docker("rm", EnvRegistryName)
	return docker("run", "-d", "--restart", "always", "-p", format("127.0.0.1:%s:5000", EnvRegistryPort), "--name", EnvRegistryName, "registry:2")
}
