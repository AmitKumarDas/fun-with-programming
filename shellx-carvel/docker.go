package shellx_carvel

import (
	"carvel.shellx.dev/internal/sh"
	"errors"
)

func setupRegistryAsLocalDockerContainer() error {
	if err := docker("inspect", "-f", "{{.State.Running}}", "${REGISTRY_NAME}"); err != nil {
		var envErr *sh.InvalidArgError // Must be a pointer
		if errors.As(err, &envErr) {
			return err
		}
		return docker("run", "-d", "--restart", "always", "-p", "127.0.0.1:${REGISTRY_PORT}:5000", "--name", "${REGISTRY_NAME}", "registry:2")
	}
	return nil
}
