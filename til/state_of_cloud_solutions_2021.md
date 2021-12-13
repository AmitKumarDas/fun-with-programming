### Bring Your Own CA
```yaml
- https://github.com/square/certstrap
```

### Ingress Based Solutions - Examples / Samples
```yaml
- https://github.com/nginxinc/kubernetes-ingress/tree/master/examples
```

### Troubleshooting
```yaml
- https://learnk8s.io/troubleshooting-deployments
```

### CI CD
```yaml
- https://cloud.google.com/blog/topics/developers-practitioners/devops-and-cicd-google-cloud-explained
```

### Solutioning via Helm
```
- https://github.com/tilt-dev/tilt/tree/master/configs/storybook
- Maybe Helm for Inner Dev Loop

- Helm's Unique Selling Point
  - Broad Community
  - Registry
  - Push
  - Artifacts
  - Code Generators
  - Controllers via Flux, K3s, etc.
```

#### Declarative Helming
```yaml
- https://www.openfaas.com/blog/tilt/
- Helmfile
```

### Architecture Diagram
```yaml
- https://threedots.tech/post/auto-generated-c4-architecture-diagrams-in-go/
```

### One Page Terraform to Create EC2 Instance
```yaml
- https://gist.github.com/ankyit/d180cdc2843a21204f27473a6c7eeb2c
```

```yaml
- https://thenewstack.io/weave-gitops-core-integrates-git-with-kubernetes/
- git - kubernetes - gitops
```

```yaml
- https://github.com/cloudposse
- terraform - automation - aws - helm
```

```yaml
- https://engineering.monday.com/kubernetes_migration/
- hpa - aws cloudwatch - argo rollouts
```

```yaml
- https://docs.webhookrelay.com/installation-options/containerized/kubernetes-installation
- ingress tunnel
```

```yaml
- https://cloud.google.com/blog/topics/retail/shopify-and-google-cloud-team-up-for-an-epic-bfcm-weekend
- google virtual private cloud - https://cloud.google.com/vpc
- multi region - disaster recovery
```

```yaml
- https://engineering.monday.com/monday-coms-multi-regional-architecture-a-deep-dive/
- saas platform
- multi region - privacy & compliance - latency - availability

- request flow:
- user - cloudflare - load balancer - ambassador edge stack - routes to services
- ambassador edge stack - Web Application Firewall (WAF)
- ambassador edge stack - auth service

- ambassador edge stack:
- envoy based L7 API gateway

- us-user.monday.com - us load balancer
- eu-user.monday.com - eu load balancer
- api.monday.com - ? load balancer
- anonymous-user.monday.com - ? load balancer
- webhooks - ?.monday.com # how will third party know which region to invoke

- does not scale with growing number of customers:
- mapping each account to unique DNS record
- expensive
- not scalable when customers grow beyond thousands

- how to be vendor neutral:
- cloudflare workers
- aws cloudfront with lambda@edge
- fastly compute@edge

- reproducibility:
- how do you build a new region

- multi region design:
- us:
  vpc:
    - front facing lb - ambassador edge stack - internal lb
    - internal lb - transit gateway
  transit gateway # out of vpc
- eu:
  vpc:
    - front facing lb - ambassador edge stack - internal lb
    - internal lb - transit gateway
  transit gateway # out of vpc
- us-eu:
  - transit gateway peering
```
