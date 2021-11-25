### State of CLI | Tools | Commands | Protocols | Ideas

### Preview Environments
```yaml
- https://github.com/k8s-at-home/template-cluster-k3s
- https://github.com/toboshii/hajimari
- https://github.com/stakater/Forecastle
- k3s - dns - cloudflare - ingress discovery - flux the helm - helm as custom resource
- you may not need a EKS or ETCD - just k3s maybe
- helm - flux - kustomization - cert manager - acme - no AWS - no EKS - on-premise possible

- https://github.com/k8s-at-home/awesome-home-kubernetes
- awesome - k3s - home lab - raspberry

- https://www.cncf.io/blog/2021/11/16/prometheus-announces-an-agent-to-address-a-new-range-of-use-cases/
- https://prometheus.io/blog/2021/11/16/agent/
- forward from preview to prod prometheus
- makes preview environment support monitoring & still be light weight

```

### Compliance | E2E
```yaml
- https://prometheus.io/blog/2021/10/14/prometheus-conformance-results/
- https://prometheus.io/blog/2021/05/04/prometheus-conformance-remote-write-compliance/
- https://prometheus.io/blog/2021/05/03/introducing-prometheus-conformance-program/
```

### Network
```yaml
- dig
- https://gist.github.com/mrlesmithjr/c42ebb99a01e8eeeca6a5eb4fa52f852
- dig output in json
- dig k8s.io +noall +answer | awk '{if (NR>3){print}}'| tr '[:blank:]' ';'| jq -R 'split(";") |{Name:.[0],TTL:.[1],Class:.[2],Type:.[3],IpAddress:.[4]}' | jq --slurp '.'
```

### Git
```yaml
- https://www.linode.com/docs/guides/git-rebase-command/
- git - rebase vs merge

- steps you need to put your commits on top of latest commit in upstream
- git checkout main
- git pull upstream main
- git checkout my-feature-branch
- git rebase main

- git log --pretty=oneline
```
