package shellx_carvel

func installKindCLI() error {
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
