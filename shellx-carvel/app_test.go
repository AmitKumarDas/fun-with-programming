package shellx_carvel

import "testing"

func tryAppDeployment(t *testing.T) {
	requireNoErr(t, verifyApplication())
}
