package shellx_carvel

import "testing"

func tryVerifyApplication(t *testing.T) {
	requireNoErr(t, deployKappController())
	requireNoErr(t, verifyApplication())
}
