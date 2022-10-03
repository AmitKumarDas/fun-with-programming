package shellx_carvel

import (
	shx "carvel.shellx.dev/internal/sh"
	"fmt"
	"log"
	"strings"
)

var (
	fileK8sAppRBAC              string
	fileK8sCarvelPackageRepo    string
	fileK8sCarvelPackageRBAC    string
	fileK8sCarvelPackageInstall string
)

func setupK8sDeploymentDirAndFilePaths() error {
	// files
	fileK8sCarvelPackageRepo = dirK8sArtifacts + "/package-repo.yml"
	fileK8sAppRBAC = dirK8sArtifacts + "/application-rbac.yml"
	fileK8sCarvelPackageRBAC = dirK8sArtifacts + "/package-rbac.yml"
	fileK8sCarvelPackageInstall = dirK8sArtifacts + "/package-install.yml"
	return nil
}

func verifyApplication() error {
	if shx.IsNotEq(EnvTestCarvelRelease, "true") {
		return nil
	}

	if err := setupK8sDeploymentDirAndFilePaths(); err != nil {
		return err
	}

	// clean up resources from previous run if any
	if err := cleanK8sResources(); err != nil && shx.IsDebug() {
		log.Println(err)
	}

	var fns = []func() error{
		deployKappController,
		createK8sArtifactsDir,
		createFilePackageRepository,
		applyAppNamespace,
		applyPackageRepository,
		verifyPresenceOfPackageRepository,
		verifyPresenceOfPackage,
		createFileAppRBACResources,
		applyAppRBACResources,
		createFileCarvelRBACResources,
		applyCarvelRBACResources,
		createFilePackageInstallResources,
		applyPackageInstallResources,
		isAppPodRunning,
	}
	for _, fn := range fns {
		if err := fn(); err != nil {
			debugK8sResources()
			return err
		}
	}
	return nil
}

func createK8sArtifactsDir() error {
	return mkdir(dirK8sArtifacts)
}

func deployKappController() error {
	return kubectl("apply", "-f", format("https://github.com/vmware-tanzu/carvel-kapp-controller/releases/download/%s/release.yml", EnvKappCtrlVersion))
}

func createFilePackageRepository() error {
	return file(fileK8sCarvelPackageRepo, packageRepositoryYML, 0644)
}

func applyAppNamespace() error {
	return kubectl("create", "ns", EnvK8sNamespace)
}

func applyPackageRepository() error {
	return eventually(func() error {
		return kubectl("create", "-f", fileK8sCarvelPackageRepo)
	})
}

func verifyPresenceOfPackageRepository() error {
	return eventually(func() error {
		return kubectl("get", "packagerepository", "-n", EnvK8sNamespace, EnvPackageRepoName)
	})
}

func verifyPresenceOfPackage() error {
	return eventually(func() error {
		return kubectl("get", "package", "-n", EnvK8sNamespace, EnvPackageName+"."+EnvPackageVersion)
	})
}

func createFileAppRBACResources() error {
	return file(fileK8sAppRBAC, applicationRBACYML, 0644)
}

func applyAppRBACResources() error {
	return eventually(func() error {
		return kubectl("create", "-f", fileK8sAppRBAC)
	})
}

func createFileCarvelRBACResources() error {
	return file(fileK8sCarvelPackageRBAC, carvelPackageRBACYML, 0644)
}

func applyCarvelRBACResources() error {
	return eventually(func() error {
		return kubectl("create", "-f", fileK8sCarvelPackageRBAC)
	})
}

func createFilePackageInstallResources() error {
	return file(fileK8sCarvelPackageInstall, carvelPackageInstallYML, 0644)
}

func applyPackageInstallResources() error {
	return eventually(func() error {
		return kubectl("create", "-f", fileK8sCarvelPackageInstall)
	})
}

func debugK8sResources() {
	_ = shx.RunV("kubectl", "describe", "pkgr", "-n", EnvK8sNamespace)
	_ = shx.RunV("kubectl", "describe", "package", "-n", EnvK8sNamespace)
	_ = shx.RunV("kubectl", "describe", "pkgi", "-n", EnvK8sNamespace)
}

func cleanK8sResources() error {
	var mErr shx.MultiError
	(&mErr).Add(kubectl("delete", "app", "-n", EnvK8sNamespace, EnvPackageInstallName))
	(&mErr).Add(kubectl("delete", "-f", fileK8sCarvelPackageInstall))
	(&mErr).Add(kubectl("delete", "-f", fileK8sCarvelPackageRBAC))
	(&mErr).Add(kubectl("delete", "-f", fileK8sAppRBAC))
	(&mErr).Add(kubectl("delete", "-f", fileK8sCarvelPackageRepo))
	(&mErr).Add(kubectl("delete", "ns", EnvK8sNamespace))
	(&mErr).Add(kubectl("delete", "-f", format("https://github.com/vmware-tanzu/carvel-kapp-controller/releases/download/%s/release.yml", EnvKappCtrlVersion)))
	return (&mErr).ErrOrNil()
}

func isAppPodRunning() error {
	cmd := "kubectl"
	args, cmdErr := shx.ExpandStrictAll([]string{"get", "po", "-n", EnvK8sNamespace, "-l", EnvAppDeploymentLabelKey + "=" + EnvAppDeploymentLabelVal, "-o", "custom-columns=:.status.phase"}...)
	if cmdErr != nil {
		return cmdErr
	}
	err := eventually(func() error {
		out, err := shx.Output(cmd, args...)
		if err != nil {
			return err
		}
		if strings.Contains(strings.ToLower(strings.TrimSpace(out)), "running") { // SUCCESS
			return nil
		}
		return fmt.Errorf("found pod with state %q", out)
	})
	if err != nil {
		return fmt.Errorf("%s %s: %w", cmd, format("%s", strings.Join(args, " ")), err)
	}
	return nil
}
