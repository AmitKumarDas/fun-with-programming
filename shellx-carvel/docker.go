package shellx_carvel

import (
	shx "carvel.shellx.dev/internal/sh"
	"errors"
)

func setupRegistryAsLocalDockerContainer() error {
	if shx.IsNotEq(EnvSetupLocalRegistry, "true") {
		return nil
	}
	if err := docker("inspect", "-f", "{{.State.Running}}", EnvRegistryName); err != nil {
		var envErr *shx.InvalidEnvError // Note: Must be a pointer
		if errors.As(err, &envErr) {
			return err
		}
		// Note: This error is mostly due to docker in NOT RUNNING state
		// Start a local registry as docker service
		return docker("run", "-d", "--restart", "always", "-p", format("127.0.0.1:%s:5000", EnvRegistryPort), "--name", EnvRegistryName, "registry:2")
	}
	return nil
}
