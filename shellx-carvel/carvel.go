package shellx_carvel

import "github.com/magefile/mage/sh"

func installCarvelCLIs() error {
	if err := mkdir(binPathCarvel); err != nil {
		return err
	}
	if isErr(whichKbld(), whichImgpkg(), whichYtt()) {
		installScriptPath := format("%s/install-carvel.sh", EnvBinPathCarvel)
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
