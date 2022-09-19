package shellx_carvel

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
	return file(joinPaths(EnvArtifactsPathK8s, EnvFilePackageRepository), packageRepositoryYML, 0644)
}

func deleteThenCreatePackageRepo() error {
	pkgRepoFilePath := joinPaths(EnvArtifactsPathK8s, EnvFilePackageRepository)
	_ = kubectl("delete", "-f", pkgRepoFilePath) // ignore error if any
	return kubectl("create", "-f", pkgRepoFilePath)
}
