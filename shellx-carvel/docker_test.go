package shellx_carvel

import "testing"

func TestSetupRegistryAsLocalDockerContainer(t *testing.T) {
	requireNoErr(t, setupRegistryAsLocalDockerContainer())
}
