package shellx_carvel

import (
	"fmt"
	"github.com/magefile/mage/sh"
	"testing"
)

func tryAppBundleCreateAndPublish(t *testing.T) {
	sourceDir := "tmp/source"
	configDir := sourceDir + "/config"
	imgpkgDir := sourceDir + "/.imgpkg"

	requireNoErr(t, createAppDirs(configDir, imgpkgDir))
	requireTrue(t, exists(configDir))
	requireTrue(t, exists(imgpkgDir))

	requireNoErr(t, createAppConfigs(configDir+"/config.yml", configDir+"/values.yml"))
	requireTrue(t, exists(configDir+"/config.yml"))
	requireTrue(t, exists(configDir+"/values.yml"))

	requireNoErr(t, createAppBundle(configDir, imgpkgDir+"/images.yml"))
	requireTrue(t, exists(imgpkgDir+"/images.yml"))

	requireNoErr(t, publishAppBundle(sourceDir))
	out, outErr := sh.Output("curl", "${REGISTRY_NAME}:${REGISTRY_PORT}/v2/_catalog")
	requireNoErr(t, outErr)
	requireContains(t, out, fmt.Sprintf("packages/%s", getEnvTrimKey(EnvAppBundleName)))
}
