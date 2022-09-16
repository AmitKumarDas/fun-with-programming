package shellx_carvel

import (
	"fmt"
)

func installKindCLI() error {
	if err := mkdir(binPathKind); err != nil {
		return err
	}
	// Install Kind CLI only if it was not installed previously
	if isErr(whichKind()) {
		installPath := format("%s/kind", EnvBinPathKind)
		if err := curl(format("https://kind.sigs.k8s.io/dl/%s/kind-%s-%s", EnvKindVersion, EnvGOOS, EnvGOARCH), "-Lo", installPath); err != nil {
			return err
		}
		if err := chmod("+x", installPath); err != nil {
			return err
		}
	}
	return nil
}

func createKindClusterConfigForLocalRegistry() error {
	if isNotEq(EnvSetupLocalRegistry, "true") {
		return nil
	}
	if err := mkdir(EnvArtifactsPathKind); err != nil {
		return err
	}
	return file(joinPaths(EnvArtifactsPathKind, EnvFileKindCluster), kindClusterLocalRegistryYML, 0644)
}

func setupKindCluster() error {
	if isNotEq(EnvSetupKindCluster, "true") {
		return nil
	}
	kindClusterFilePath := joinPaths(EnvArtifactsPathKind, EnvFileKindCluster)
	if !exists(kindClusterFilePath) {
		return fmt.Errorf("file %q not found", kindClusterFilePath)
	}
	if err := kubectl("cluster-info", "--context", "kind-kind"); err == nil {
		// No error implies Kind cluster is up & running
		// Hence, nothing to do
		return nil
	}
	return kind("create", "cluster", "--config", kindClusterFilePath)
}
