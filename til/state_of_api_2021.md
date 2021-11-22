### State of API 2021


```yaml
- https://dev.to/harmonydevelopment/introducing-hrpc-a-simple-rpc-system-for-user-facing-apis-16ge
- hrpc - user facing APIs - protobuf

- https://blog.cloudflare.com/moving-k8s-communication-to-grpc/
- pod to pod communciation - http vs grpc - dns pipeline - protobuf
- large DNS zones - protobuf - quick record change propagation for serving API at the edge
- kind: NetworkPolicy - what goes in - what goes out - ingress - egress
- network policy - protect APIs at network level - no protection at application level
- https://grpc.io/docs/guides/auth/
- per rpc credentials - protection at app level - idea - mayastor
- HTTP/3 - quic

- https://itnext.io/websocket-1-million-connections-using-appwrite-2d2a2c363a37
- performance - design - testing

- https://github.com/processone/tsung
- scalability testing - multi protocol - MySQL - PostgreSQL - HTTP

- https://github.blog/2021-11-10-make-your-monorepo-feel-small-with-gits-sparse-index/
- monorepo - git - sparse index - contributing guide

- https://kubernetes.io/blog/2021/04/13/kube-state-metrics-v-2-0/
- learn from release notes - testing - troubleshooting - api design
- naming - standards
- Flag --namespace was renamed to --namespaces
- Flag --collectors was renamed to --resources
- --metric-blacklist to --metric-denylist
- --metric-whitelist to --metric-allowlist
- kube_hpa_* were renamed to kube_horizontalpodautoscaler_*
- Metric labels that relate to Kubernetes were converted to snake_case

- https://github.com/Peripli/service-broker-proxy-k8s
- service broker - k8s - proxy

- https://github.com/kubernetes/kube-aggregator
- https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/aggregated-api-servers.md
- register - discover - proxy - multiple api servers - custom api servers
```

```yaml
https://github.com/kubernetes-sigs/prometheus-adapter/
https://github.com/kubernetes-sigs/prometheus-adapter/blob/master/deploy/manifests/custom-metrics-apiservice.yaml

// custom api server implementation
// Kubernetes API server arguments // authentication // authorization
```

### Programming Language / Runtime
```yaml
- https://blog.cloudflare.com/automatically-generated-types/
- workers - rust - typescript - intermediate representation
- abstract syntax tree
```
