### Certificate Management - Bootstrap
```yaml
- https://github.com/square/certstrap
- your own certificate authority
- your own public key infrastructure
```

### Pod Security
```yaml
- https://kubernetes.io/blog/2021/12/09/pod-security-admission-beta/
```

#### Feature Rich Libraries
```yaml
- https://github.com/open-policy-agent/gatekeeper

- https://github.com/kubewarden
- rust - go - fellow

- https://kyverno.io/policies/pod-security/
```

#### Compliance
```yaml
- kubectl get --raw /metrics | grep pod_security_evaluations_total
```

### Sign Verify & Protect Software
```yaml
- https://github.com/sigstore
- https://dlorenc.medium.com/the-sigstore-trust-model-4b146b2ecf2c
- https://github.com/golang/go/issues/25530

- go - rust - fellow
- sigstore vs pgp
- oidc - domains
```
