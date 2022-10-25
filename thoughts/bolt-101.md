# Learn BOLT from a BOLT PR(s)

## Release
```yaml
- A release is a collection of components that make up a release graph
- A release graph is also referred to as a build graph or a dependency graph
  - It is a Directed Acyclic Graph
  - It represents the build dependencies between all the components involved
```

## Component
```yaml
- Any project that needs to be built as part of a release is considered a component
- Each component is defined by a version specific config YAML, a publish.yaml and a url.yaml
- BOLT supports the following types of components:
  - Managed
  - Unmanaged
  - Image
  - BOM
  - Aggregator
```

## Learn From MR 2500
```yaml
- MR: https://gitlab.eng.vmware.com/TKG/bolt/bolt-release-yamls/-/merge_requests/2500/
- TITLE: 'add v1.7.0-zshippable release graph'
- RELEASE: 'release/tkg-v1.7.0-zshippable.yaml'
```

### release/tkg-v1.7.0-zshippable.yaml
```yaml
stagingImageRepo:
  imageRepository: projects-stg.registry.vmware.com/tkg
description: tkg release yaml
componentReferences:
- name: tkg_release
  version:
  - v1.7.0-zshippable
```

### component/tkg_release/v1.7.0-zshippable.yaml
```yaml
name: tkg_release
aggregators:
- version: v1.7.0-zshippable
  gobuildBaseBranch: vmware-0.0.0+vmware.0~b72ff37523de3fb6de99fa16c21e61707f8c70d3
  dependencies:
    gobuild.linux_hosttype: linux-centos72-gc32
  componentReferences:
  - name: tkg-compatibility
    version:
    - v1.7.0-zshippable
  - name: tanzu-framework-management-packages
    version:
    - v0.25.0
  - name: standalone-plugins-package
    version:
    - v0.25.0-standalone-plugins
```

### component/tkg-compatibility/v1.7.0-zshippable.yaml
```yaml
name: tkg-compatibility
boms:
- version: v1.7.0-zshippable
  type: tkg-compatibility
  componentReferences:
  - name: tanzu-framework
    version:
    - v0.25.0
  - name: tkg-bom
    version:
    - v1.7.0-zshippable
```

### component/tkg-bom/v1.7.0-zshippable.yaml
```yaml
name: tkg-bom
boms:
- version: v1.7.0-zshippable
  type: tkg
  imageRepository: projects-stg.registry.vmware.com/tkg
  k8sVersion: v1.23.8+vmware.2-tkg.2-zshippable
  dependencies:
    apiVersion: run.tanzu.vmware.com/v1alpha2
    componentTanzuCoreVersion: v0.25.0
  componentReferences:
  - name: alertmanager
    version:
    - v0.24.0+vmware.1
  - name: cloud-provider-azure
    version:
    - v0.7.4+vmware.1
  - name: aad-pod-identity
    version:
    - v1.8.0+vmware.1
  - name: cluster_api
    version:
    - v1.1.5+vmware.1
  - name: cluster_api_aws
    version:
    - v1.2.0+vmware.1
  - name: cluster_api_vsphere
    version:
    - v1.3.1+vmware.1
  - name: cluster-api-provider-azure
    version:
    - v1.4.0+vmware.1
  - name: cluster-api-provider-bringyourownhost
    version:
    - v0.2.0+vmware.4
  - name: cluster-api-provider-oci
    version:
    - v0.3.0+vmware.1
  - name: configmap-reload
    version:
    - v0.7.1+vmware.1
  - name: contour
    version:
    - v1.17.2+vmware.1
    - v1.18.2+vmware.1
    - v1.20.2+vmware.1
  - name: crash-diagnostics
    version:
    - v0.3.7+vmware.5
  - name: envoy
    version:
    - v1.18.4+vmware.1
    - v1.19.1+vmware.1
    - v1.21.3+vmware.1
  - name: external-dns
    version:
    - v0.11.0+vmware.1
  - name: fluent-bit
    version:
    - v1.8.15+vmware.1
  - name: grafana
    version:
    - v7.5.16+vmware.1
  - name: harbor
    version:
    - v2.5.3+vmware.1
  - name: imgpkg
    version:
    - v0.29.0+vmware.1
  - name: jetstack_cert-manager
    version:
    - v1.5.3+vmware.4
    - v1.7.2+vmware.1
  - name: k8s-sidecar
    version:
    - v1.12.1+vmware.2
    - v1.15.6+vmware.1
  - name: k14s_kapp
    version:
    - v0.49.0+vmware.1
  - name: k14s_ytt
    version:
    - v0.41.1+vmware.1
  - name: vendir
    version:
    - v0.27.0+vmware.1
  - name: kbld
    version:
    - v0.34.0+vmware.1
  - name: kube_rbac_proxy
    version:
    - v0.11.0+vmware.2
  - name: kube-state-metrics
    version:
    - v2.5.0+vmware.1
  - name: kube-vip
    version:
    - v0.4.2+vmware.1
  - name: kubernetes_autoscaler
    version:
    - v1.23.0+vmware.1
    - v1.22.0+vmware.1
    - v1.21.0+vmware.1
  - name: kubernetes-sigs_kind
    version:
    - v1.23.8+vmware.2-tkg.1_v0.11.1
  - name: prometheus
    version:
    - v2.36.2+vmware.1
  - name: prometheus_node_exporter
    version:
    - v1.3.1+vmware.1
  - name: pushgateway
    version:
    - v1.4.3+vmware.1
  - name: sonobuoy
    version:
    - v0.56.6+vmware.1
  - name: tkg-standard-packages
    version:
    - v1.7.0-zshippable
  - name: tkg_telemetry
    version:
    - v1.6.0+vmware.1
  - name: tkr-bom
    version:
    - v1.23.8+vmware.2-tkg.2-zshippable
    - v1.21.14+vmware.2-tkg.2-zshippable
    - v1.22.11+vmware.2-tkg.2-zshippable
  - name: tkr-compatibility
    version:
    - v18-v1.7.0-zshippable
  - name: velero
    version:
    - v1.8.1+vmware.1
  - name: velero-plugin-for-aws
    version:
    - v1.4.1+vmware.1
  - name: velero-plugin-for-microsoft-azure
    version:
    - v1.4.1+vmware.1
  - name: velero-plugin-for-vsphere
    version:
    - v1.3.1+vmware.1
  - name: multus-cni
    version:
    - v3.8.0+vmware.1
  - name: image-builder
    version:
    - v0.1.12+vmware.1
  - name: whereabouts
    version:
    - v0.5.1+vmware.2
```

### Other Interesting Bolt PRs
```yaml
- MR: https://gitlab.eng.vmware.com/TKG/bolt/bolt-release-yamls/-/merge_requests/2930
- TAG: Fuji.1
- Why: Fixes CVEs reported in https://jira.eng.vmware.com/browse/TKG-14564
- Title: Update Velero components for TKG 1.6.1
```

```yaml
- MR: https://gitlab.eng.vmware.com/TKG/bolt/bolt-release-yamls/-/merge_requests/2933
- TAG: Fuji.1
- Title: [build] add kapp-controller release v0.38.5+vmware.2
```

## References:
```yaml
- https://gitlab.eng.vmware.com/TKG/bolt/bolt-release-yamls/-/blob/main/release/README.md
- https://gitlab.eng.vmware.com/TKG/bolt/bolt-release-yamls/-/blob/main/component/README.md
```
