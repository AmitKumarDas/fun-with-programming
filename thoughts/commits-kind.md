## Learn kind from its commits
This is one of my ideas to learn a project. In other words, read and perhaps try interesting
commits from the project. This should help me in understanding parts of the project by focusing
on some particular fix or feature. Alternative ways e.g. getting involved with the community
or spending weeks &/ months may not be feasible unless it is part of my day job. Needless to say
this works better when the project follows atomic commits.

## Commits
### Issues when a new node re-uses PodCIDR of deleted node
```yaml
- Pod to Pod communication is broken
- Node route via the old node is problematic
- TAG: kube-apiserver, webhook communication, CNI
- FILE: images/kindnetd/cmd/kindnetd/routes.go
- PR: https://github.com/kubernetes-sigs/kind/pull/2941
```

### Install open-iscsi & impact on image size
```yaml
- TIL: Maintain custom KIND images via Dockerfiles
- PR: https://github.com/kubernetes-sigs/kind/pull/2905
```

### ENV: no_proxy to automatically include all the nodes of the cluster
```yaml
- TROUBLESHOOT: `systemctl status kubelet`
- TROUBLESHOOT: `journalctl -xeu kubelet`
- List Containers: `crictl --runtime-endpoint unix:///run/containerd/containerd.sock ps -a | grep kube | grep -v pause`
- Container Logs: `crictl --runtime-endpoint unix:///run/containerd/containerd.sock logs CONTAINERID`
- API Server Health: `curl --verbose https://test-control-plane:6443/healthz?timeout=10s`
- Docker: `docker info`
- ISSUE: https://github.com/kubernetes-sigs/kind/issues/2884
- FILE: pkg/cluster/internal/providers/podman/provision.go
- PR: https://github.com/kubernetes-sigs/kind/pull/2885
```
