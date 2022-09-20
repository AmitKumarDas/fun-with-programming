package shellx_carvel

import shx "carvel.shellx.dev/internal/sh"

func verifyApplication() error {
	if isNotEq(EnvTestCarvelRelease, "true") {
		return nil
	}
	var fns = []func() error{
		deleteThenCreateAppNamespace,
		createK8sArtifactsDir,
		createFilePackageRepo,
		deleteThenCreatePackageRepo,
	}
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

func deleteThenCreateAppNamespace() error {
	_ = kubectl("delete", "ns", EnvK8sNamespace) // ignore error if any
	return kubectl("create", "ns", EnvK8sNamespace)
}

func createK8sArtifactsDir() error {
	return mkdir(EnvArtifactsPathK8s)
}

func createFilePackageRepo() error {
	fullPath, pathErr := shx.JoinPaths(EnvArtifactsPathK8s, EnvFilePackageRepository)
	if pathErr != nil {
		return pathErr
	}
	return file(fullPath, packageRepositoryYML, 0644)
}

func deleteThenCreatePackageRepo() error {
	fullPath, pathErr := shx.JoinPaths(EnvArtifactsPathK8s, EnvFilePackageRepository)
	if pathErr != nil {
		return pathErr
	}
	_ = kubectl("delete", "-f", fullPath) // ignore error if any
	return kubectl("create", "-f", fullPath)
}
