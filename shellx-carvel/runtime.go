package shellx_carvel

import (
	shx "carvel.shellx.dev/internal/sh"
	"runtime"
	"strings"
)

// Environment variables in a format that can be expanded
var (
	EnvGOOS    = "${GOOS}" // Generic
	EnvGOARCH  = "${GOARCH}"
	EnvVersion = "${VERSION}"

	EnvBinPathKind                        = "${BIN_PATH_KIND}" // KIND
	EnvArtifactsPathKind                  = "${ARTIFACTS_PATH_KIND}"
	EnvFileKindCluster                    = "${FILE_KIND_CLUSTER}"
	EnvFileKindConfigLocalRegistryHosting = "${FILE_KIND_CONFIG_LOCAL_REGISTRY_HOSTING}"
	EnvSetupKindCluster                   = "${SETUP_KIND_CLUSTER}"
	EnvKindVersion                        = "${KIND_VERSION}"

	EnvRegistryName       = "${REGISTRY_NAME}" // Registry
	EnvRegistryPort       = "${REGISTRY_PORT}"
	EnvSetupLocalRegistry = "${SETUP_LOCAL_REGISTRY}"

	EnvK8sNamespace            = "${K8S_NAMESPACE}" // K8s RBAC
	EnvK8sServiceAccount       = "${K8S_SERVICE_ACCOUNT}"
	EnvK8sRole                 = "${K8S_ROLE}"
	EnvK8sRoleBinding          = "${K8S_ROLE_BINDING}"
	EnvK8sServiceAccountCarvel = "${K8S_SERVICE_ACCOUNT_CARVEL}"
	EnvK8sRoleCarvel           = "${K8S_ROLE_CARVEL}"
	EnvK8sRoleBindingCarvel    = "${K8S_ROLE_BINDING_CARVEL}"

	EnvAppDeploymentName     = "${APP_DEPLOYMENT_NAME}" // Application
	EnvAppDeploymentLabelKey = "${APP_DEPLOYMENT_LABEL_KEY}"
	EnvAppDeploymentLabelVal = "${APP_DEPLOYMENT_LABEL_VAL}"
	EnvAppImageName          = "${APP_IMAGE_NAME}"
	EnvAppImageVersion       = "${APP_IMAGE_VERSION}"

	EnvBinPathCarvel      = "${BIN_PATH_CARVEL}" // Carvel
	EnvKappCtrlVersion    = "${KAPP_CTRL_VERSION}"
	EnvAppBundleName      = "${APP_BUNDLE_NAME}"
	EnvAppBundleVersion   = "${APP_BUNDLE_VERSION}"
	EnvPackageName        = "${PACKAGE_NAME}"
	EnvPackageVersion     = "${PACKAGE_VERSION}"
	EnvPackageRepoName    = "${PACKAGE_REPO_NAME}"
	EnvPackageRepoVersion = "${PACKAGE_REPO_VERSION}"
	EnvPackageInstallName = "${PACKAGE_INSTALL_NAME}"
	EnvTestCarvelRelease  = "${TEST_CARVEL_RELEASE}"

	EnvDirCarvelPackaging = "${DIR_CARVEL_PACKAGING}" // Folder to store carvel artifacts
	EnvDirK8sArtifacts    = "${DIR_K8S_ARTIFACTS}"    // Folder to store K8s artifacts
)

// Immediate setting of few environment variables
var version = shx.MaybeSetEnv(EnvVersion, "v1.0.1")
var versionSemver = strings.TrimPrefix(version, "v")
var binPathCarvel = shx.MaybeSetEnv(EnvBinPathCarvel, "tmp")
var binPathKind = shx.MaybeSetEnv(EnvBinPathKind, "tmp")
var dirCarvelPackaging = shx.MaybeSetEnv(EnvDirCarvelPackaging, "packaging")
var dirK8sArtifacts = shx.MaybeSetEnv(EnvDirK8sArtifacts, "artifacts/k8s")

// Carvel binaries / CLIs as functions
var kbld = shx.RunCmd(binPathCarvel + "/kbld")
var whichKbld = shx.RunCmd("ls", binPathCarvel+"/kbld")
var imgpkg = shx.RunCmd(binPathCarvel + "/imgpkg")
var whichImgpkg = shx.RunCmd("ls", binPathCarvel+"/imgpkg")
var ytt = shx.RunCmd(binPathCarvel + "/ytt")
var whichYtt = shx.RunCmd("ls", binPathCarvel+"/ytt")

// KIND CLI as function
var kind = shx.RunCmd(binPathKind + "/kind")
var whichKind = shx.RunCmd("ls", binPathKind+"/kind")

func init() {
	const (
		appName       = "k8s-remediator"
		packageDomain = "experiment.dev.com"
	)

	// environment keys & corresponding default values
	//
	// Note: This helps in expanding an env variable as $ENV_KEY_NAME
	// during command execution
	envs := map[string]string{
		// OS & Architecture
		EnvGOOS:   runtime.GOOS,
		EnvGOARCH: runtime.GOARCH,

		// KIND cluster
		EnvSetupKindCluster:                   "false",
		EnvKindVersion:                        "v0.15.0",
		EnvArtifactsPathKind:                  "tmp/artifacts/kind",
		EnvFileKindCluster:                    "kind-cluster.yml",
		EnvFileKindConfigLocalRegistryHosting: "local-registry-hosting.yml",

		// container registry
		EnvSetupLocalRegistry: "false",
		EnvRegistryName:       "kind-registry.local",
		EnvRegistryPort:       "5000",

		// k8s rbac
		EnvK8sNamespace:            appName + "-system",
		EnvK8sServiceAccount:       appName,
		EnvK8sRole:                 appName + "-role",
		EnvK8sRoleBinding:          appName + "-role-binding",
		EnvK8sServiceAccountCarvel: "carvel-pkg-install",
		EnvK8sRoleCarvel:           "carvel-pkg-install-role",
		EnvK8sRoleBindingCarvel:    "carvel-pkg-install-role-binding",

		// deployment
		EnvAppDeploymentName:     appName,
		EnvAppDeploymentLabelKey: packageDomain + "/app",
		EnvAppDeploymentLabelVal: appName + "-controller",
		EnvAppImageName:          "amitnist/tkg-remediator", // should exist
		EnvAppImageVersion:       "latest",                  // should exist

		// carvel
		EnvKappCtrlVersion:    "v0.40.0",
		EnvAppBundleName:      appName + "-app",
		EnvAppBundleVersion:   versionSemver,
		EnvPackageName:        appName + "." + packageDomain,
		EnvPackageVersion:     versionSemver,
		EnvPackageRepoName:    appName + "-repo." + packageDomain,
		EnvPackageRepoVersion: versionSemver,
		EnvPackageInstallName: appName + "-install",
		EnvTestCarvelRelease:  "false",
	}
	for k, v := range envs {
		_ = shx.MaybeSetEnv(k, v)
	}

	// display all the environment variables for debuggability
	if shx.IsDebug() {
		_ = shx.RunV("env")
	}
}
