package shellx_carvel

import (
	"fmt"
	"github.com/magefile/mage/sh"
	"testing"
)

func TestCreateAndPublishAppBundle(t *testing.T) {
	sourceDir := "tmp/source"
	configDir := sourceDir + "/config"
	imgpkgDir := sourceDir + "/.imgpkg"

	requireNoErr(t, createAppConfigs(configDir))
	requireTrue(t, exists(configDir))
	requireTrue(t, exists(configDir+"/config.yml"))
	requireTrue(t, exists(configDir+"/values.yml"))

	requireNoErr(t, createAppBundle(imgpkgDir, configDir))
	requireTrue(t, exists(imgpkgDir))
	requireTrue(t, exists(imgpkgDir+"/images.yml"))

	requireNoErr(t, publishAppBundle(sourceDir))
	out, outErr := sh.Output("curl", "${REGISTRY_NAME}:${REGISTRY_PORT}/v2/_catalog")
	requireNoErr(t, outErr)
	requireEqual(t, fmt.Sprintf("{\"repositories\":[\"packages/%s\"]}", getEnvTrimKey(EnvAppBundleName)), out)
}
