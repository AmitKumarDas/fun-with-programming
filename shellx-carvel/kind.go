package shellx_carvel

import (
	"fmt"
	"github.com/magefile/mage/sh"
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
	if isNotEq(EnvSetupKindCluster, "true") {
		return nil
	}
	var fns = []func() error{
		createKindClusterConfigForLocalRegistry,
		createKindCluster,
		createKindNetwork,
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
	return file(joinPaths(EnvArtifactsPathKind, EnvFileKindCluster), kindClusterLocalRegistryYML, 0644)
}

func createKindCluster() error {
	kindClusterFilePath := joinPaths(EnvArtifactsPathKind, EnvFileKindCluster)
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

func createKindNetwork() error {
	if err := docker("inspect", "-f", "{{json .NetworkSettings.Networks.kind}}", EnvRegistryName); err != nil {
		// Connect the network on error
		// Note: error is swallowed
		return docker("network", "connect", "kind", EnvRegistryName)
	}
	return nil
}

func createKindConfigFileLocalRegistryHosting() error {
	return file(joinPaths(EnvArtifactsPathKind, EnvFileKindConfigLocalRegistryHosting), kindConfigLocalRegistryHostingYML, 0644)
}

func applyKindConfigLocalRegistryHosting() error {
	return kubectl("apply", "-f", joinPaths(EnvArtifactsPathKind, EnvFileKindConfigLocalRegistryHosting))
}

func printEtcHostsUpdateMsg() error {
	return sh.RunV("echo", etcHostsUpdateMsg)
}
