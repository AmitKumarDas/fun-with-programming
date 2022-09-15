package shellx_carvel

import "testing"

func trySetupRegistryAsLocalDockerContainer(t *testing.T) {
	requireNoErr(t, setupRegistryAsLocalDockerContainer())
}
