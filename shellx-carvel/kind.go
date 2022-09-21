package shellx_carvel

import (
	shx "carvel.shellx.dev/internal/sh"
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

func setupKindCluster() error {
	if shx.IsNotEq(EnvSetupKindCluster, "true") {
		return nil
	}
	var fns = []func() error{
		createKindClusterConfigForLocalRegistry,
		createKindCluster,
		disconnectThenConnectNetworkToKindEndpoint,
		createKindConfigFileLocalRegistryHosting,
		applyKindConfigLocalRegistryHosting,
		printEtcHostsUpdateMsg,
	}
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

func createKindClusterConfigForLocalRegistry() error {
	if err := mkdir(EnvArtifactsPathKind); err != nil {
		return err
	}
	fullPath, pathErr := shx.JoinPaths(EnvArtifactsPathKind, EnvFileKindCluster)
	if pathErr != nil {
		return pathErr
	}
	return file(fullPath, kindClusterLocalRegistryYML, 0644)
}

func createKindCluster() error {
	kindClusterFilePath, pathErr := shx.JoinPaths(EnvArtifactsPathKind, EnvFileKindCluster)
	if pathErr != nil {
		return pathErr
	}
	if !exists(kindClusterFilePath) {
		return fmt.Errorf("file %q not found", kindClusterFilePath)
	}
	if err := kubectl("cluster-info", "--context", "kind-kind"); err != nil {
		// Create kind cluster on error
		// Note: error is swallowed
		return kind("create", "cluster", "--config", kindClusterFilePath)
	}
	return nil
}

func disconnectThenConnectNetworkToKindEndpoint() error {
	_ = docker("network", "disconnect", "kind", EnvRegistryName) // ignore error if any
	return docker("network", "connect", "kind", EnvRegistryName)
}

func createKindConfigFileLocalRegistryHosting() error {
	fullPath, pathErr := shx.JoinPaths(EnvArtifactsPathKind, EnvFileKindConfigLocalRegistryHosting)
	if pathErr != nil {
		return pathErr
	}
	return file(fullPath, kindConfigLocalRegistryHostingYML, 0644)
}

func applyKindConfigLocalRegistryHosting() error {
	fullPath, pathErr := shx.JoinPaths(EnvArtifactsPathKind, EnvFileKindConfigLocalRegistryHosting)
	if pathErr != nil {
		return pathErr
	}
	return kubectl("apply", "-f", fullPath)
}

func printEtcHostsUpdateMsg() error {
	return shx.RunV("echo", etcHostsUpdateMsg)
}
