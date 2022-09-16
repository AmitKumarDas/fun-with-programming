package shellx_carvel

import (
	shx "carvel.shellx.dev/internal/sh"
	"github.com/magefile/mage/sh"
	"runtime"
)

// Environment variables in a format that can be expanded
var (
	EnvGOOS    = "${GOOS}" // Generic envs
	EnvGOARCH  = "${GOARCH}"
	EnvVersion = "${VERSION}"

	EnvBinPathKind       = "${BIN_PATH_KIND}" // Kind envs
	EnvArtifactsPathKind = "${ARTIFACTS_PATH_KIND}"
	EnvFileKindCluster   = "${FILE_KIND_CLUSTER}"
	EnvSetupKindCluster  = "${SETUP_KIND_CLUSTER}"
	EnvKindVersion       = "${KIND_VERSION}"

	EnvSetupLocalRegistry = "${SETUP_LOCAL_REGISTRY}" // Registry envs
	EnvRegistryName       = "${REGISTRY_NAME}"
	EnvRegistryPort       = "${REGISTRY_PORT}"

	EnvK8sNamespace      = "${K8S_NAMESPACE}" // K8s envs
	EnvK8sServiceAccount = "${K8S_SERVICE_ACCOUNT}"
	EnvK8sRole           = "${K8S_ROLE}"
	EnvK8sRoleBinding    = "${K8S_ROLE_BINDING}"

	EnvAppDeploymentName     = "${APP_DEPLOYMENT_NAME}" // Deploy envs
	EnvAppDeploymentLabelKey = "${APP_DEPLOYMENT_LABEL_KEY}"
	EnvAppDeploymentLabelVal = "${APP_DEPLOYMENT_LABEL_VAL}"
	EnvAppImageName          = "${APP_IMAGE_NAME}"
	EnvAppImageVersion       = "${APP_IMAGE_VERSION}"

	EnvBinPathCarvel      = "${BIN_PATH_CARVEL}" // Carvel envs
	EnvKappCtrlVersion    = "${KAPP_CTRL_VERSION}"
	EnvAppBundleName      = "${APP_BUNDLE_NAME}"
	EnvAppBundleVersion   = "${APP_BUNDLE_VERSION}"
	EnvPackageName        = "${PACKAGE_NAME}"
	EnvPackageVersion     = "${PACKAGE_VERSION}"
	EnvPackageRepoName    = "${PACKAGE_REPO_NAME}"
	EnvPackageRepoVersion = "${PACKAGE_REPO_VERSION}"
)

// Immediate setting of few environment variables
var version = maybeSetEnv(EnvVersion, "v1.0.1")
var binPathCarvel = maybeSetEnv(EnvBinPathCarvel, "tmp")
var binPathKind = maybeSetEnv(EnvBinPathKind, "tmp")

// Carvel binaries / CLIs as functions
var kbld = shx.RunCmdStrict(binPathCarvel + "/kbld")
var whichKbld = shx.RunCmdStrict("ls", binPathCarvel+"/kbld")
var imgpkg = shx.RunCmdStrict(binPathCarvel + "/imgpkg")
var whichImgpkg = shx.RunCmdStrict("ls", binPathCarvel+"/imgpkg")
var ytt = shx.RunCmdStrict(binPathCarvel + "/ytt")
var whichYtt = shx.RunCmdStrict("ls", binPathCarvel+"/ytt")

// KIND CLI as function
var kind = shx.RunCmdStrict(binPathKind + "/kind")
var whichKind = shx.RunCmdStrict("ls", binPathKind+"/kind")

// kubectl cli as function
var kubectl = shx.RunCmdStrict("kubectl")

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
		EnvSetupKindCluster:  "false",
		EnvKindVersion:       "v0.15.0",
		EnvArtifactsPathKind: "tmp/artifacts/kind",
		EnvFileKindCluster:   "kind-cluster.yml",

		// container registry
		EnvSetupLocalRegistry: "false",
		EnvRegistryName:       "kind-registry.local",
		EnvRegistryPort:       "5000",

		// k8s rbac
		EnvK8sNamespace:      appName + "-system",
		EnvK8sServiceAccount: appName,
		EnvK8sRole:           appName + "-role",
		EnvK8sRoleBinding:    appName + "-role-binding",

		// deployment
		EnvAppDeploymentName:     appName,
		EnvAppDeploymentLabelKey: packageDomain + "/app",
		EnvAppDeploymentLabelVal: appName + "-controller",
		EnvAppImageName:          "amitnist/tkg-remediator", // should exist
		EnvAppImageVersion:       "latest",                  // should exist

		// carvel
		EnvKappCtrlVersion:    "v0.40.0",
		EnvAppBundleName:      appName + "-app",
		EnvAppBundleVersion:   version,
		EnvPackageName:        appName + "." + packageDomain,
		EnvPackageVersion:     version,
		EnvPackageRepoName:    appName + "-repo." + packageDomain,
		EnvPackageRepoVersion: version,
	}
	for k, v := range envs {
		maybeSetEnv(k, v)
	}

	// display all the environment variables for debuggability
	sh.RunV("env")
}
