package shellx_carvel

import (
	"fmt"
	"github.com/magefile/mage/sh"
	"os"
	"runtime"
	"strings"
)

func maybeSetEnv(envKey, defaultVal string) string {
	if value := os.Getenv(envKey); value == "" {
		os.Setenv(envKey, defaultVal)
	}
	return os.Getenv(envKey)
}

func maybeSetEnvTrimKey(envKey, defaultVal string) string {
	k := strings.TrimPrefix(envKey, "$")
	return maybeSetEnv(k, defaultVal)
}

func isErr(err error, more ...error) bool {
	if err != nil {
		return true
	}
	for _, e := range more {
		if e != nil {
			return true
		}
	}
	return false
}

func isNoErr(err error, more ...error) bool {
	return !isErr(err, more...)
}

// isAlphaNum reports whether the byte is an ASCII letter, number, or underscore
//
// Note: Borrowed from os package
func isAlphaNum(c uint8) bool {
	return c == '_' || '0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}

// isShellSpecialVar reports whether the character identifies a special
// shell variable such as $*.
//
// Note: Borrowed from os package
func isShellSpecialVar(c uint8) bool {
	switch c {
	case '*', '#', '$', '@', '!', '?', '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	}
	return false
}

// getShellName returns the name that begins the string and the number of bytes
// consumed to extract it. If the name is enclosed in {}, it's part of a ${}
// expansion and two more bytes are needed than the length of the name.
//
// Note: Borrowed from os package
func getShellName(s string) (string, int) {
	switch {
	case s[0] == '{':
		if len(s) > 2 && isShellSpecialVar(s[1]) && s[2] == '}' {
			return s[1:2], 3
		}
		// Scan to closing brace
		for i := 1; i < len(s); i++ {
			if s[i] == '}' {
				if i == 1 {
					return "", 2 // Bad syntax; eat "${}"
				}
				return s[1:i], i + 1
			}
		}
		return "", 1 // Bad syntax; eat "${"
	case isShellSpecialVar(s[0]):
		return s[0:1], 1
	}
	// Scan alphanumerics.
	var i int
	for i = 0; i < len(s) && isAlphaNum(s[i]); i++ {
	}
	return s[:i], i
}

type invalidArgError struct{ Msg string }

func (e *invalidArgError) Error() string { return e.Msg }

// expandEnvStrict replaces ${var} or $var in the string. It returns
// error for invalid use of $ or env expansion returns empty
//
// Note: Borrowed from os.ExpandEnv
func expandEnvStrict(s string) (string, error) {
	var buf []byte
	// ${} is all ASCII, so bytes are fine for this operation.
	i := 0
	for j := 0; j < len(s); j++ {
		if s[j] == '$' && j+1 < len(s) {
			if buf == nil {
				buf = make([]byte, 0, 2*len(s))
			}
			buf = append(buf, s[i:j]...)
			name, w := getShellName(s[j+1:])
			if name == "" && w > 0 {
				// Encountered invalid syntax
				return "", &invalidArgError{fmt.Sprintf("invalid %s", s)}
			} else if name == "" {
				// Valid syntax, but $ was not followed by a
				// name. Leave the dollar character untouched.
				buf = append(buf, s[j])
			} else {
				val, found := os.LookupEnv(name)
				if !found {
					return "", &invalidArgError{fmt.Sprintf("%s lookup failed", name)}
				}
				buf = append(buf, val...)
			}
			j += w
			i = j + 1
		}
	}
	if buf == nil {
		return s, nil
	}
	return string(buf) + s[i:], nil
}

// verifyArgs returns error if the arguments makes use of unset
// environment variables
func verifyArgs(args ...string) error {
	if len(args) == 0 {
		return nil
	}
	var invalidArgs = make([]string, 0, len(args))
	for _, arg := range args {
		if _, err := expandEnvStrict(arg); err != nil {
			invalidArgs = append(invalidArgs, err.Error())
		}
	}
	if len(invalidArgs) == 0 {
		return nil
	}
	return &invalidArgError{fmt.Sprintf("verify args ['%s']: %s", strings.Join(args, "', '"), strings.Join(invalidArgs, ", "))}
}

// runCmdStrict is a wrapper over sh.Run. The returned function
// returns error if either verification of arguments fails or
// command execution fails
func runCmdStrict(cmd string, args ...string) func(args ...string) error {
	return func(args2 ...string) error {
		allArgs := append(args, args2...)
		if err := verifyArgs(allArgs...); err != nil {
			return err
		}
		return sh.Run(cmd, allArgs...)
	}
}

func file(name, data string, perm os.FileMode) error {
	if err := verifyArgs(data); err != nil {
		return err
	}
	out, outErr := sh.Output("echo", data)
	if outErr != nil {
		return outErr
	}
	return os.WriteFile(name, []byte(out), perm)
}

// Environment variables in a format that can be expanded
var (
	EnvGOOS                  = "$GOOS"
	EnvGOARCH                = "$GOARCH"
	EnvVersion               = "$VERSION"
	EnvBinPathCarvel         = "$BIN_PATH_CARVEL"
	EnvBinPathKind           = "$BIN_PATH_KIND"
	EnvSetupKindCluster      = "$SETUP_KIND_CLUSTER"
	EnvKindVersion           = "$KIND_VERSION"
	EnvRegistryName          = "$REGISTRY_NAME"
	EnvRegistryPort          = "$REGISTRY_PORT"
	EnvK8sNamespace          = "$K8S_NAMESPACE"
	EnvK8sServiceAccount     = "$K8S_SERVICE_ACCOUNT"
	EnvK8sRole               = "$K8S_ROLE"
	EnvK8sRoleBinding        = "$K8S_ROLE_BINDING"
	EnvKappCtrlVersion       = "$KAPP_CTRL_VERSION"
	EnvAppDeploymentName     = "$APP_DEPLOYMENT_NAME"
	EnvAppDeploymentLabelKey = "$APP_DEPLOYMENT_LABEL_KEY"
	EnvAppDeploymentLabelVal = "$APP_DEPLOYMENT_LABEL_VAL"
	EnvAppImageName          = "$APP_IMAGE_NAME"
	EnvAppImageVersion       = "$APP_IMAGE_VERSION"
	EnvAppBundleName         = "$APP_BUNDLE_NAME"
	EnvAppBundleVersion      = "$APP_BUNDLE_VERSION"
	EnvPackageName           = "$PACKAGE_NAME"
	EnvPackageVersion        = "$PACKAGE_VERSION"
	EnvPackageRepoName       = "$PACKAGE_REPO_NAME"
	EnvPackageRepoVersion    = "$PACKAGE_REPO_VERSION"
)

// Environment values that are accessed directly i.e. not as expanded format
var appName = "k8s-remediator"
var packageDomain = "experiment.dev.com"
var version = maybeSetEnvTrimKey(EnvVersion, "v1.0.1")
var binPathCarvel = maybeSetEnvTrimKey(EnvBinPathCarvel, "tmp")
var binPathKind = maybeSetEnvTrimKey(EnvBinPathKind, "tmp")

// Carvel binaries / CLIs as functions
var kbld = runCmdStrict(binPathCarvel + "/kbld")
var whichKbld = runCmdStrict("ls", binPathCarvel+"/kbld")
var imgpkg = runCmdStrict(binPathCarvel + "/imgpkg")
var whichImgpkg = runCmdStrict("ls", binPathCarvel+"/imgpkg")
var ytt = runCmdStrict(binPathCarvel + "/ytt")
var whichYtt = runCmdStrict("ls", binPathCarvel+"/ytt")

// KIND CLI as function
var kind = runCmdStrict(binPathKind + "/kind")
var whichKind = runCmdStrict("ls", binPathKind+"/kind")

// Docker CLI as function
var docker = runCmdStrict("docker")

// Unix & generic commands as functions
var mkdir = runCmdStrict("mkdir", "-p")
var curl = runCmdStrict("curl")
var ls = runCmdStrict("ls")
var chmod = runCmdStrict("chmod")

func init() {
	// environment keys & corresponding default values
	//
	// Note: This helps in expanding an env variable as $ENV_KEY_NAME
	// during command execution
	envs := map[string]string{
		// OS & Architecture
		EnvGOOS:   runtime.GOOS,
		EnvGOARCH: runtime.GOARCH,

		// kind cluster
		EnvSetupKindCluster: "false",
		EnvKindVersion:      "v0.15.0",

		// container registry
		EnvRegistryName: "kind-registry.local",
		EnvRegistryPort: "5000",

		// k8s rbac
		EnvK8sNamespace:      appName + "-system",
		EnvK8sServiceAccount: appName,
		EnvK8sRole:           appName + "-role",
		EnvK8sRoleBinding:    appName + "-role-binding",

		// versions & names
		EnvKappCtrlVersion:       "v0.40.0",
		EnvAppDeploymentName:     appName,
		EnvAppDeploymentLabelKey: packageDomain + "/app",
		EnvAppDeploymentLabelVal: appName + "-controller",
		EnvAppImageName:          appName,
		EnvAppImageVersion:       version,
		EnvAppBundleName:         appName + "-app",
		EnvAppBundleVersion:      version,
		EnvPackageName:           appName + "." + packageDomain,
		EnvPackageVersion:        version,
		EnvPackageRepoName:       appName + "-repo." + packageDomain,
		EnvPackageRepoVersion:    version,
	}
	for k, v := range envs {
		maybeSetEnvTrimKey(k, v)
	}

	// display all the environment variables for debuggability
	sh.RunV("env")
}
