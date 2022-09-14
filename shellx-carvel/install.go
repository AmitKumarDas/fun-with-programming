package shellx_carvel

import "github.com/magefile/mage/sh"

func installCarvelCLIs() error {
	if err := mkdir(binPathCarvel); err != nil {
		return err
	}
	if isErr(whichKbld(), whichImgpkg(), whichYtt()) {
		installScriptPath := "$BIN_PATH_CARVEL/install-carvel.sh"
		if err := curl("https://carvel.dev/install.sh", "-Lo", installScriptPath); err != nil {
			return err
		}
		env := map[string]string{"K14SIO_INSTALL_BIN_DIR": binPathCarvel}
		if err := sh.RunWith(env, "bash", installScriptPath); err != nil {
			return err
		}
	}
	return nil
}

func installKind() error {
	if err := mkdir(binPathKind); err != nil {
		return err
	}
	if isErr(whichKind()) {
		installPath := "$BIN_PATH_KIND/kind"
		if err := curl("https://kind.sigs.k8s.io/dl/$KIND_VERSION/kind-$GOOS-$GOARCH", "-Lo", installPath); err != nil {
			return err
		}
		if err := chmod("+x", installPath); err != nil {
			return err
		}
	}
	return nil
}
