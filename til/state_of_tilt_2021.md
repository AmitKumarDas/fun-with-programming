### Tilt & Friends - Inner Development Loop
```yaml
- https://github.com/tilt-dev/tilt/tree/master/configs/storybook
- Helm chart, Tiltfile, Dockerfile
```

### Testing via Helm
```yaml
- Test via a Pod that is Helm-ed
- Pod with Helm annotation to test if a deployment installed via Helm Chart works
  - "helm.sh/hook": test-success
  - Use wget or any simple thing
  - This test pod itself is installed via Helm & then used a test tool
```

### Tilt Commands
```yaml
- str
- local
- rstrip
- split

- allow_k8s_contexts
- local_resource
- experimental_analytics_report
- analytics_settings
- docker_build
- sync
- k8s_yaml
- k8s_resource
- helm
```

### Learn Tilt / Starlark / Py
```yaml
- https://github.com/kubernetes-sigs/cluster-api/blob/main/Tiltfile
```
```py
settings = {
  "deploy_cert_manager": True,     # bool
  "preload_images_for_kind": True,
  "enable_providers": ["docker"],  # list
  "kind_cluster_name": "kind",
  "debug": {},                     # structure
}

# global settings
settings.update(read_json(           # .update? json.update?
  "tilt-settings.json",
  default = {},
))

allow_k8s_contexts(settings.get("allowed_contexts"))   # .get? json.get?

if settings.get("trigger_mode") == "manual": # json.get?
  trigger_mode(TRIGGER_MODE_MANUAL)
```

```py
envsubst_cmd = "./hack/tools/bin/envsubst"
kustomize_cmd = "./hack/tools/bin/kustomize"
yq_cmd = "./hack/tools/bin/yq"

os_name = local_output("go env GOOS")         # capture from local process
os_arch = local_output("go env GOARCH")
```

```py
default_registry(settings.get("default_registry"))  # set get
```

```py
# just in time config
providers = {
  "kubeadm-bootstrap": {
    "context": "bootstrap/kubeadm",
    "image": "gcr.io/k8s-staging-cluster-api/kubeadm-bootstrap-controller",
    "live_reload_deps": [
      "main.go",
      "api",
      "../../go.mod",
      "../../go.sum",
    ],
    "label": "CABPK",
  },
  "docker": {
    "additional_docker_helper_commands": "RUN wget -qO- https://dl.k8s.io/v1.21.2/kubernetes-client-linux-{arch}.tar.gz | tar xvz".format(
        arch = os_arch,
    ),
    "additional_docker_build_commands": """
COPY --from=tilt-helper /go/kubernetes/client/bin/kubectl /usr/bin/kubectl
""",
  },
}
```

