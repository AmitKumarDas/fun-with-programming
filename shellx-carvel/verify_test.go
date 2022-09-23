package shellx_carvel

import "testing"

func tryVerifyApplication(t *testing.T) {
	requireNoErr(t, verifyApplication())
}
