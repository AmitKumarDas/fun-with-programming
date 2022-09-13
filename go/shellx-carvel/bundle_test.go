package shellx_carvel

import "testing"

func TestCreateConfigYML(t *testing.T) {
	requireNoErr(t, createAppConfigs("tmp"))
}
