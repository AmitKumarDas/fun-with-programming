package shellx_carvel

import (
	shx "carvel.shellx.dev/internal/sh"
	"fmt"
	"sync"
)

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
		if err := shx.RunWith(env, "bash", installScriptPath); err != nil {
			return err
		}
	}
	return nil
}

var onlyOnce sync.Once

var (
	dirCarvelSource            string
	dirCarvelSourceConfig      string
	dirCarvelSourceImgpkg      string
	dirCarvelSourceSchema      string
	dirCarvelRelease           string
	dirCarvelReleaseImgpkg     string
	dirCarvelReleasePkgRepo    string
	dirCarvelReleaseTemplates  string
	dirCarvelReleasePkgRepoPkg string
	fileConfig                 string
	fileConfigValues           string
	fileConfigValuesOpenAPI    string
	fileCarvelSourceImgpkg     string
	filePackageTemplate        string
	filePackageMetadata        string
	filePackageVerion          string
	fileCarvelReleaseImgpkg    string
)

func setupCarvelDirAndFilePaths() error {
	var err shx.Error
	onlyOnce.Do(func() {
		// directories
		dirCarvelSource = shx.JoinPathsWithErrHandle(&err, dirCarvelPackaging, "source")
		dirCarvelSourceConfig = shx.JoinPathsWithErrHandle(&err, dirCarvelSource, "config")
		dirCarvelSourceImgpkg = shx.JoinPathsWithErrHandle(&err, dirCarvelSource, ".imgpkg")
		dirCarvelSourceSchema = shx.JoinPathsWithErrHandle(&err, dirCarvelSource, "schema")
		dirCarvelRelease = shx.JoinPathsWithErrHandle(&err, dirCarvelPackaging, "release")
		dirCarvelReleaseImgpkg = shx.JoinPathsWithErrHandle(&err, dirCarvelRelease, ".imgpkg")
		dirCarvelReleasePkgRepo = shx.JoinPathsWithErrHandle(&err, dirCarvelRelease, "packages")
		dirCarvelReleaseTemplates = shx.JoinPathsWithErrHandle(&err, dirCarvelRelease, "templates")
		dirCarvelReleasePkgRepoPkg = shx.JoinPathsWithErrHandle(&err, dirCarvelReleasePkgRepo, EnvPackageName)

		// files
		fileConfig = dirCarvelSourceConfig + "/config.yml"
		fileConfigValues = dirCarvelSourceConfig + "/values.yml"
		fileConfigValuesOpenAPI = dirCarvelSourceSchema + "/schema-openapi.yml"
		fileCarvelSourceImgpkg = dirCarvelSourceImgpkg + "/images.yml"
		filePackageTemplate = shx.JoinPathsWithErrHandle(&err, dirCarvelReleaseTemplates, EnvPackageName+"-template.yml")
		filePackageMetadata = dirCarvelReleasePkgRepoPkg + "/package-metadata.yml"
		filePackageVerion = shx.JoinPathsWithErrHandle(&err, dirCarvelReleasePkgRepoPkg, EnvPackageVersion+".yml")
		fileCarvelReleaseImgpkg = dirCarvelReleaseImgpkg + "/images.yml"
	})
	return (&err).ErrOrNil()
}

func cutAppBundle() error {
	var fns = []func() error{
		setupCarvelDirAndFilePaths,
		createAppDirs,
		createAppConfigs,
		createAppBundle,
		publishAppBundle,
	}
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

func createAppDirs() error {
	return mkdirAll(dirCarvelSourceConfig, dirCarvelSourceImgpkg, dirCarvelSourceSchema)
}

func createAppConfigs() error {
	if err := file(fileConfig, appDeploymentYML, 0644); err != nil {
		return err
	}
	return file(fileConfigValues, appValuesYML, 0644)
}

func createAppBundle() error {
	return kbld("-f", dirCarvelSourceConfig, "--imgpkg-lock-output", fileCarvelSourceImgpkg)
}

func publishAppBundle() error {
	bundle := format("%s:%s/packages/%s:%s", EnvRegistryName, EnvRegistryPort, EnvAppBundleName, EnvAppBundleVersion)
	return imgpkg("push", "-b", bundle, "-f", dirCarvelSource)
}

func cutCarvelRelease() error {
	var fns = []func() error{
		setupCarvelDirAndFilePaths,
		mkdirForCarvelRelease,
		createPackageMetadata,
		generateOpenAPISchema,
		createPackageTemplate,
		createPackageVersion,
		createPackageRepoBundle,
		publishPackageRepoBundle,
	}
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

func mkdirForCarvelRelease() error {
	return mkdirAll(dirCarvelReleasePkgRepoPkg, dirCarvelReleaseTemplates, dirCarvelReleaseImgpkg)
}

func createPackageMetadata() error {
	return file(filePackageMetadata, packageMetadataYML, 0644)
}

func generateOpenAPISchema() error {
	if !exists(fileConfigValues) {
		return fmt.Errorf("file %s not found", fileConfigValues) // TODO: Check if this error message is helpful
	}
	out, outErr := shx.Output(EnvBinPathCarvel+"/ytt", "-f", fileConfigValues, "--data-values-schema-inspect", "-o", "openapi-v3")
	if outErr != nil {
		return outErr
	}
	return file(fileConfigValuesOpenAPI, out+"\n", 0644)
}

func createPackageTemplate() error {
	return file(filePackageTemplate, packageTemplateYML, 0644)
}

func createPackageVersion() error {
	out, outErr := shx.Output(EnvBinPathCarvel+"/ytt", "-f", filePackageTemplate, "--data-value-file", "openapi="+fileConfigValuesOpenAPI, "-v", "version="+EnvPackageVersion)
	if outErr != nil {
		return outErr
	}
	return file(filePackageVerion, out+"\n", 0644)
}

func createPackageRepoBundle() error {
	return kbld("-f", dirCarvelReleasePkgRepo, "--imgpkg-lock-output", fileCarvelReleaseImgpkg)
}

func publishPackageRepoBundle() error {
	return imgpkg("push", "-b", format("%s:%s/packages/%s:%s", EnvRegistryName, EnvRegistryPort, EnvPackageRepoName, EnvPackageRepoVersion), "-f", dirCarvelRelease)
}
