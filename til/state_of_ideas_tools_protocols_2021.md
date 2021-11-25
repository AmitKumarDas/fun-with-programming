### State of Ideas | Tools | Protocols | Commands

### Protocols
```yaml
- https://docs.google.com/document/d/1LPhVRSFkGNSuU1fBd81ulhsCPR4hkSZyyBj1SZ8fWOM/edit#heading=h.3p42p5s8n0ui
- https://prometheus.io/blog/2021/11/16/agent/
- Prometheus Remote Write Specs - Write Ahead Logging - WAL - TSDB WAL
- idea - can / should mayastor be a Remote Write Compliant Project?

- https://prometheus.io/blog/2016/07/23/pull-does-not-scale-or-does-it/
- push vs pull vs event - at scale
```

### Serverless
```yaml
- https://groups.google.com/g/prometheus-developers/c/FPe0LsTfo2E/m/yS7up2YzAwAJ?pli=1
- prometheus - serverless
```

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
- https://github.com/prometheus/compliance/tree/main/remote_write_sender
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