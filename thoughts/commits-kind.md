## Learn kind from its commits
This is one of my ideas to learn a project. In other words, read and perhaps try interesting
commits from the project. This should help me in understanding parts of the project by focusing
on some particular fix or feature. Alternative ways e.g. getting involved with the community
or spending weeks &/ months may not be feasible unless it is part of my day job. Needless to say
this works better when the project follows atomic commits.

## Commits
### Issues when a new node re-uses PodCIDR of deleted node
- Pod to Pod communication is broken
- TAG: kube-apiserver, webhook communication, CNI
- Node route via the old node is problematic
- FILE: images/kindnetd/cmd/kindnetd/routes.go
- PR: https://github.com/kubernetes-sigs/kind/pull/2941

### Install open-iscsi & impact on image size
- TIL: Maintain custom KIND images via Dockerfiles
- PR: https://github.com/kubernetes-sigs/kind/pull/2905

### no_proxy env variable to automatically include all the nodes of the cluster
- TIL: TROUBLESHOOT: `systemctl status kubelet`
- TIL: TROUBLESHOOT: `journalctl -xeu kubelet`
- TIL: List Containers by crictl: `crictl --runtime-endpoint unix:///run/containerd/containerd.sock ps -a | grep kube | grep -v pause`
- TIL: Inspect Container Logs: `crictl --runtime-endpoint unix:///run/containerd/containerd.sock logs CONTAINERID`
- TIL: API server health checks: `curl --verbose https://test-control-plane:6443/healthz?timeout=10s`
- TIL: Docker: `docker info`
- ISSUE: https://github.com/kubernetes-sigs/kind/issues/2884
- FILE: pkg/cluster/internal/providers/podman/provision.go
- PR: https://github.com/kubernetes-sigs/kind/pull/2885
