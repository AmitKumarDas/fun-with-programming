### State of Ideas | Tools | Protocols | Commands

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

### System
```yaml
- https://github.com/mudler/yip
- cloud init
```

### Serverless
```yaml
- https://groups.google.com/g/prometheus-developers/c/FPe0LsTfo2E/m/yS7up2YzAwAJ?pli=1
- prometheus - serverless
```

### Compliance | E2E
```yaml
- https://prometheus.io/blog/2021/10/14/prometheus-conformance-results/
- https://prometheus.io/blog/2021/05/04/prometheus-conformance-remote-write-compliance/
- https://prometheus.io/blog/2021/05/03/introducing-prometheus-conformance-program/
- https://github.com/prometheus/compliance/tree/main/remote_write_sender

- quickcheck
```

### Network
```yaml
- dig
- https://gist.github.com/mrlesmithjr/c42ebb99a01e8eeeca6a5eb4fa52f852
- dig output in json
- dig k8s.io +noall +answer | awk '{if (NR>3){print}}'| tr '[:blank:]' ';'| jq -R 'split(";") |{Name:.[0],TTL:.[1],Class:.[2],Type:.[3],IpAddress:.[4]}' | jq --slurp '.'

- https://github.com/inguardians/peirates/blob/master/enumerate_dns.go
- https://github.com/inguardians/peirates/blob/master/portscan.go
```

### SSH | Troubleshoot
```yaml
- SSH-Key has not been added to the ssh-agent
- Solve this with ssh-add -K
- You may put that command inside your .bashrc or .zshrc

- SSH does not know the host key of the bastion host yet
- Run below
- ssh -v <githubusername>@bastion.jimdo-platform-eks.net
- and then accept the question with yes
```
